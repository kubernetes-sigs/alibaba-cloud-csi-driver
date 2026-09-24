package nas

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
)

// Which credential a mount is running on is tracked here rather than by whoever
// produced the credential, because a mount can get one from three places: an rrsa
// token exchange, the publish secret kubelet re-reads on every republish, or
// mountOptions. All three land in Options.AkID, and all three raise the same
// question once per republish: is this the credential the mount already has?
//
// The access key id answers it: a rotated credential always has a new one, whether
// it comes from STS (a new session each time) or from RAM (an access key secret
// cannot be changed in place). It is also the only part of a credential that is
// safe to log, which is what makes "which credential is this mount on" answerable
// in production.
type installedCredentials struct {
	akIDs sync.Map // mount target -> access key id installed on it
}

func (c *installedCredentials) get(target string) (akID string, known bool) {
	v, ok := c.akIDs.Load(target)
	if !ok {
		return "", false
	}
	akID, _ = v.(string)
	return akID, true
}

// record remembers that target is running on akID. An empty akID records nothing:
// there is no credential to compare against later.
func (c *installedCredentials) record(target, akID string) {
	if akID == "" {
		return
	}
	c.akIDs.Store(target, akID)
}

// forget drops what is known about target, so a later mount of the same path does
// not inherit the credential of an unmounted one.
func (c *installedCredentials) forget(target string) {
	c.akIDs.Delete(target)
}

// checkCredentialInstallSupport reports why this node cannot install a credential
// on a live mount, or nil when it can.
//
// The mount broker is the only mounter that can do it, and it ships in its own
// image, so its support is asked over the socket rather than assumed from this
// process being built with the client half.
func (ns *nodeServer) checkCredentialInstallSupport(ctx context.Context) error {
	// NasMounter satisfies ProxyRefresher whichever mounter it wraps, so which mode
	// NAS is in has to be read from the configuration.
	if ns.config.MountProxySocket == "" {
		return errors.New("NAS is not using the mount broker: enable the AlinasMountProxy feature gate on csi-plugin")
	}
	refresher, ok := ns.mounter.(mounter.ProxyRefresher)
	if !ok {
		return errors.New("this mounter cannot install a credential on a live mount")
	}
	can, err := refresher.CanRefresh(ctx)
	if err != nil {
		return fmt.Errorf("ask the mount broker whether it can refresh credentials: %w", err)
	}
	if !can {
		return errors.New("the mount broker does not support credential refresh: upgrade mount-proxy-server")
	}
	return nil
}

// checkExpiringCredential refuses a volume whose credential expires on a node that
// could never install the next one. Such a mount works, right up to the moment the
// credential expires, and then turns into EACCES on lookups and hanging writes with
// nothing left to connect it to the cause; saying so before anything is mounted is
// the only useful moment.
//
// Whether the credential expires is known before it exists: authType rrsa always
// exchanges an STS credential, and a publish secret carries one when it has a
// security token. A long-term access key is exempt, which is what keeps volumes
// that never needed rotation working on nodes that cannot rotate.
func (ns *nodeServer) checkExpiringCredential(ctx context.Context, opt *Options) error {
	if opt.AuthType != AuthTypeRRSA && opt.SecurityToken == "" {
		return nil
	}
	if err := ns.checkCredentialInstallSupport(ctx); err != nil {
		return status.Errorf(codes.FailedPrecondition,
			"this volume uses a credential that expires, and this node cannot install its replacement on a live mount: %v", err)
	}
	return nil
}

// syncMountCredentials installs the credential opt now carries when the mount at
// target is running on a different one, and remembers it only once the install
// succeeded, so a failure is retried on the next republish instead of being
// forgotten.
//
// How hard it tries depends on whether the credential expires, because that decides
// what silence costs:
//
//   - an STS credential must reach the mount. Not installing it leaves the mount to
//     die when the token expires, as EACCES on lookups and hanging writes far from
//     here, so a node that cannot install one says so on every republish.
//   - a long-term access key is best effort. A node that cannot install credentials
//     leaves the mount on the key it was mounted with, which is the behaviour it
//     always had, and records nothing, so what the user sees does not change with a
//     csi-plugin restart.
func (ns *nodeServer) syncMountCredentials(ctx context.Context, volumeID, target string, opt *Options) error {
	if opt.AkID == "" {
		// An unauthenticated mount, or jwtauth: there the mount broker holds the
		// token and rotates the credential on its own schedule.
		return nil
	}
	// A losetup volume mounts the NAS share elsewhere and exposes a loop device at
	// target, so no credential can be installed on it.
	if opt.MountType == LosetupType {
		return nil
	}
	op, err := prepareRefresh(opt, target, volumeID, ns.config.AgentMode)
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	// Only a mount the broker performed holds a credential it can replace. A plain
	// NFS mount whose Secret happens to carry an access key is not one, and asking
	// anyway would fail every republish with "fstype not supported".
	if !isAlinasFstype(op.FsType) {
		return nil
	}

	// No record means the mount was made by a process that is gone, or by a version
	// that kept no record: what it is running on is unknown, so it is installed.
	if installed, known := ns.installed.get(target); known && installed == opt.AkID {
		return nil
	}
	if err := ns.checkCredentialInstallSupport(ctx); err != nil {
		if opt.SecurityToken != "" {
			return status.Errorf(codes.FailedPrecondition,
				"credential %s of %s expires and has to be installed on the live mount, but this node cannot: %v",
				opt.AkID, target, err)
		}
		// Connector or agent mode, or a mount broker too old to install a credential
		// on a live mount. The key does not expire, so the mount keeps working on
		// whatever it was mounted with.
		klog.V(4).InfoS("node cannot install credentials on a live mount, leaving it as mounted",
			"target", target, "accessKeyId", opt.AkID, "reason", err)
		return nil
	}

	refresher, ok := ns.mounter.(mounter.ProxyRefresher)
	if !ok {
		// Unreachable: checkCredentialInstallSupport asserts the same interface.
		return status.Errorf(codes.Internal, "mounter cannot install credentials on %s", target)
	}
	if err := refresher.Refresh(ctx, op); err != nil {
		return status.Errorf(codes.Internal, "install credential %s on %s: %v", opt.AkID, target, err)
	}
	ns.installed.record(target, opt.AkID)
	klog.InfoS("installed rotated credential on a live mount", "target", target, "accessKeyId", opt.AkID)
	return nil
}
