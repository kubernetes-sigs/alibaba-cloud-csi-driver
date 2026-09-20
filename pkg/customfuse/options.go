//go:build !windows

package customfuse

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	fpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/interceptors"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog/v2"
)

// unknownBucketName is the bucket_name label fallback. Source is not used here: it is an opaque client-specific string, not a bucket.
const unknownBucketName = "unknown"

type fuseOptions struct {
	// Source is the mount source passed to the FUSE entrypoint as $source. Its
	// format is opaque to the driver: a metadata engine URL, a "bucket:path"
	// pair, or any other string the entrypoint understands. The driver composes
	// nothing here — whatever the volume said is what arrives, and a client that
	// wants a derived form builds it from $bucket and $path itself.
	//
	// Being opaque cuts both ways: the driver cannot tell whether one carries a
	// credential, and a URL of that kind commonly embeds a password. So it is
	// only ever handed to the entrypoint. It is not logged and not used as a
	// metrics label, because both reach a far wider audience than the RBAC on
	// the PV that supplied it; metricsBucketName is what the label uses instead.
	Source string
	// Bucket is the object storage bucket name, passed as $bucket.
	// Independent of Source and URL — all three are separate env vars.
	Bucket string
	URL    string
	// OtherOpts originates from volumeAttributes.otherOpts, or from an
	// "otherOpts=" entry in spec.mountOptions where volumeAttributes left it
	// empty. Passed to the entrypoint as one whole $otherOpts string — the
	// driver never splits it, so its internal format (comma separated, "-o"
	// prefixed, ...) belongs to the entrypoint that reads it.
	OtherOpts string
	// Path is the sub-path within the volume, passed as $path.
	Path string
	// ReadOnly is derived from VolumeCapability or the CSI readOnly flag, never
	// from a volume parameter, so spec.mountOptions cannot change it.
	// Passed as $readOnly to the entrypoint, which maps it to FUSE "ro" option.
	ReadOnly bool
	// FuseType identifies the FUSE client for metrics labeling (e.g. "juicefs",
	// "jindo"). Defaults to "customfuse" when not set by the user.
	// Can be set via volumeAttributes.fuseType or VolumeCapability.FsType.
	FuseType string
	// MountOptions are from pv.Spec.MountOptions (via VolumeCapability.Mount.MountFlags).
	// After applyMountOptions this holds only the entries naming a field the
	// driver does not know; each of those becomes its own env var $key in the
	// entrypoint. Entries that named a field the driver does know have been
	// applied to that field and removed from here, so they are emitted once.
	MountOptions []string
	// EntrypointConfig is the name of a ConfigMap in the fuse pod's namespace
	// (ack-csi-customfuse) to mount into the fuse container at /etc/fuse-config/.
	// Kubernetes does not support cross-namespace ConfigMap mounts, so the
	// ConfigMap must be pre-created in ack-csi-customfuse by the user.
	EntrypointConfig string
	// EntrypointKey is the key within the ConfigMap that contains the entrypoint
	// script. Defaults to "entrypoint.sh". The file is always mounted as
	// /etc/fuse-config/entrypoint.sh regardless of the key name.
	EntrypointKey string
	// AuthType selects the authentication method for the FUSE client.
	// The default (empty string) passes Kubernetes Secret entries directly as
	// environment variables to the FUSE entrypoint. "agent-identity" exchanges
	// a sandbox token for a scoped STS credential in mount-proxy and delivers
	// it to the entrypoint as files, refreshed for the mount lifetime.
	// Planned future value: "rrsa".
	AuthType string
	// SandboxId identifies the sandbox whose token is exchanged for a scoped
	// STS credential. Required when AuthType is "agent-identity".
	SandboxId string
	// SandboxCredProviderName names the credential provider that issues the
	// scoped credential. Required when AuthType is "agent-identity".
	SandboxCredProviderName string
	// CredentialDir pins the directory the STS credential files are written to
	// inside the fuse container, for a client whose credential path is fixed in
	// its own configuration and so cannot be told where to look at mount time.
	// When empty each mount gets its own directory, whose path the entrypoint
	// reads from $credentialDir. Only meaningful when AuthType is
	// "agent-identity".
	CredentialDir string
	// CredentialRefreshHookKey is the key in EntrypointConfig holding a script to
	// run after each credential rotation, projected as
	// /etc/fuse-config/refresh-hook.sh. Only needed to override a hook, or to
	// supply one without building an image: a hook shipped in the image at
	// /refresh-hook.sh is picked up on its own. A client that re-reads the
	// credential files needs no hook at all; one whose credential lives elsewhere
	// once it has started must be told a rotation happened, and the hook is where
	// that is expressed. Only meaningful when AuthType is "agent-identity".
	CredentialRefreshHookKey string
	// Capacity is the volume quota passed as $capacity to the entrypoint.
	// A plain integer (e.g. "100") or a Kubernetes Quantity (e.g. "100Gi") is
	// validated and passed through unchanged; the entrypoint converts if its
	// client needs a bare number, e.g. capacity=${capacity%Gi}.
	//
	// A dynamically provisioned volume gets this from the PVC's requested size.
	// Otherwise it comes from volumeAttributes.capacity.
	//
	// This is the one field a "capacity=<value>" entry in pv.spec.mountOptions
	// may redefine, and only upward. volumeAttributes cannot be edited once the
	// PV exists, this driver implements no expansion RPC, and a dynamically
	// provisioned volume always arrives with a capacity — so without this a
	// quota could never be changed on an existing volume at all. Kubernetes
	// treats the same quantity the same way: a PersistentVolumeClaim's storage
	// request may grow but never shrink.
	Capacity string
	// Secrets holds the raw Kubernetes Secret data.
	// All entries are passed through to the FUSE entrypoint as environment
	// variables with the key as the variable name (no prefix, no transformation).
	// This means the entrypoint author controls both the Secret keys and the
	// env vars they consume — they are the same string.
	Secrets   map[string]string
	DnsPolicy corev1.DNSPolicy
	// ServiceAccountName names the ServiceAccount the fuse pod runs as, in the fuse
	// pod namespace. Empty leaves it to the namespace default.
	//
	// The pod calls no API server, so this is not about permissions: it is how a
	// private registry is reached. imagePullSecrets are honoured only from the
	// ServiceAccount or the pod spec, and the driver fills in neither on its own.
	// Spelled as in the oss driver's volumeAttributes.serviceAccountName, which the
	// same field will carry RRSA auth through once authType supports it.
	ServiceAccountName string
}

// publishRequest is the common interface for CSI Publish requests.
// Both *csi.ControllerPublishVolumeRequest and *csi.NodePublishVolumeRequest
// satisfy this interface.
type publishRequest interface {
	GetVolumeContext() map[string]string
	GetSecrets() map[string]string
	GetVolumeCapability() *csi.VolumeCapability
	GetReadonly() bool
}

// shellControlChars are rejected in the free-form option strings. The driver runs
// no shell — these go straight into the entrypoint's environment — but the
// entrypoint is a script the driver does not control, it commonly splices them onto
// a command line, and it runs in a privileged container.
//
// Structured fields and secrets are deliberately not checked: they can legitimately
// contain these characters, so quoting those stays the entrypoint's job. The
// boundary is documented in docs/customfuse.md.
var shellControlChars = []string{"\n", "\r", ";", "`", "$("}

func validateShellSafe(name, value string) error {
	for _, c := range shellControlChars {
		if strings.Contains(value, c) {
			return fmt.Errorf("%s must not contain %q", name, c)
		}
	}
	return nil
}

// splitMountOption separates an entry the way mount-proxy's buildEnvVars will,
// on the first "=". The key is matched against the driver's field names
// case-insensitively, as volumeAttributes keys are, but reported and passed
// through exactly as written.
func splitMountOption(entry string) (key, name, value string, hasValue bool) {
	key, value, hasValue = strings.Cut(entry, "=")
	key, value = strings.TrimSpace(key), strings.TrimSpace(value)
	return key, strings.ToLower(key), value, hasValue
}

// applyMountOptions reconciles pv.spec.mountOptions with the fields already
// parsed from volumeAttributes, and leaves only the entries naming a field the
// driver does not know. Those still reach the entrypoint as their own env var,
// which is what the channel is for.
//
// The two channels are not interchangeable, because volumeAttributes is what the
// driver's own consumers read: the bucket_name metrics label and the source
// fallback in parseOptions. A field that already has a value therefore keeps it,
// and the mountOptions entry is dropped with a warning rather than silently
// redefining something those consumers have already acted on. A field
// volumeAttributes left empty is filled, which matters because volumeAttributes
// cannot be edited once the PV exists — this is the only way to add a missing
// parameter to an existing volume without recreating it.
//
// capacity is the single exception, and may only be raised; see growCapacity.
//
// Names the driver fills in itself are not reachable from this channel either.
// The control fields decide how the fuse pod is built rather than what the client
// mounts, and fuseType is cross-checked against pv.spec.csi.fsType in
// parseOptions before this runs. readOnly and mountpoint are the driver's own
// outputs. Left unnamed, every one of them would pass through as an env var of
// its own name: for the control fields nothing would read it, and for mountpoint
// the entrypoint would, and mount somewhere the driver is not watching for a
// mount point.
//
// mount-proxy refuses the same names again on its own side, in buildEnvVars. The
// plugin image and the fuse client image are built and rolled out independently,
// so each has to hold the line on its own.
func applyMountOptions(opts *fuseOptions) error {
	fields := map[string]*string{
		"source":    &opts.Source,
		"bucket":    &opts.Bucket,
		"url":       &opts.URL,
		"path":      &opts.Path,
		"otheropts": &opts.OtherOpts,
	}

	var unrecognized []string
	for _, entry := range opts.MountOptions {
		key, name, value, hasValue := splitMountOption(entry)

		if dst, ok := fields[name]; ok {
			if err := fillField(key, name, dst, value, hasValue); err != nil {
				return err
			}
			continue
		}

		switch name {
		case "capacity":
			if err := growCapacity(key, &opts.Capacity, value, hasValue); err != nil {
				return err
			}
		case "readonly":
			klog.Warningf("mountOptions %s is ignored: readOnly comes from the PV's accessModes and from the publish request, not from a volume parameter (currently %v)", key, opts.ReadOnly)
		case "mountpoint":
			klog.Warningf("mountOptions %s is ignored: mountpoint is the path the driver tells the client to mount on, not a volume parameter", key)
		case "fusetype", "entrypointconfig", "entrypointkey", "dnspolicy", "serviceaccountname", "authtype":
			klog.Warningf("mountOptions %s is ignored: set %s in volumeAttributes instead", key, key)
		default:
			unrecognized = append(unrecognized, entry)
		}
	}
	opts.MountOptions = unrecognized
	return nil
}

// fillField records value in dst unless the driver already has one. The field is
// reported under the driver's spelling rather than the entry's, because the match is
// case-insensitive: an entry written PATH= fills the driver's path, and saying
// "PATH is already ..." would name a variable nobody set.
func fillField(key, name string, dst *string, value string, hasValue bool) error {
	if !hasValue || value == "" {
		return fmt.Errorf("mountOptions entry %q carries no value; write it as %s=<value>", key, key)
	}
	if *dst == "" {
		*dst = value
	} else if *dst != value {
		klog.Warningf("mountOptions %s=%q is ignored: %s is already %q, and that is the value the driver's own consumers read", key, value, name, *dst)
	}
	return nil
}

// growCapacity is the one field mountOptions may redefine, and only upward.
// Kubernetes is grow-only on the same quantity: it forbids a
// PersistentVolumeClaim from requesting less storage than it already has. A
// smaller value is dropped rather than honoured, because quietly lowering a
// quota would leave the client enforcing a limit nobody asked for.
func growCapacity(key string, dst *string, value string, hasValue bool) error {
	if !hasValue || value == "" {
		return fmt.Errorf("mountOptions entry %q carries no value; write it as %s=<value>", key, key)
	}
	newQty, err := resource.ParseQuantity(value)
	if err != nil {
		return fmt.Errorf("invalid capacity %q in mountOptions: must be a plain integer or Kubernetes Quantity (e.g. 100, 100Gi): %v", value, err)
	}
	if *dst == "" {
		*dst = value
		return nil
	}
	// volumeAttributes capacity is validated on the way in, so this only fires
	// for a fuseOptions built by hand. Comparing against a value that does not
	// parse would treat it as zero and accept anything.
	oldQty, err := resource.ParseQuantity(*dst)
	if err != nil {
		return fmt.Errorf("cannot compare mountOptions capacity %q with the existing capacity %q: %v", value, *dst, err)
	}
	switch newQty.Cmp(oldQty) {
	case 1:
		klog.Infof("capacity raised from %s to %s by mountOptions", *dst, value)
		*dst = value
	case -1:
		klog.Warningf("mountOptions %s=%s is ignored: capacity can only be raised, and it is already %s", key, value, *dst)
	}
	return nil
}

func parseOptions(req publishRequest) (*fuseOptions, error) {
	volContext := req.GetVolumeContext()
	if volContext == nil {
		volContext = map[string]string{}
	}

	opts := &fuseOptions{
		FuseType: mounterutils.CustomFuseType,
		Secrets:  req.GetSecrets(),
	}

	for k, v := range volContext {
		key := strings.TrimSpace(strings.ToLower(k))
		value := strings.TrimSpace(v)
		if value == "" {
			continue
		}
		switch key {
		case "source":
			opts.Source = value
		case "bucket":
			opts.Bucket = value
		case "path":
			opts.Path = value
		case "url":
			opts.URL = value
		case "otheropts":
			if err := validateShellSafe("otherOpts", value); err != nil {
				return nil, err
			}
			opts.OtherOpts = value
		case "serviceaccountname":
			opts.ServiceAccountName = value
		case "fusetype":
			opts.FuseType = value
		case "authtype":
			opts.AuthType = value
		case "sandboxid":
			opts.SandboxId = value
		case "sandboxcredprovidername", "credentialprovidername":
			opts.SandboxCredProviderName = value
		case "credentialdir":
			opts.CredentialDir = value
		case "credentialrefreshhookkey":
			opts.CredentialRefreshHookKey = value
		case "capacity":
			// An empty value means unset, same as omitting the key; anything else
			// has to parse, including a bare integer, which the previous
			// last-character heuristic let through unchecked.
			if value != "" {
				if _, err := resource.ParseQuantity(value); err != nil {
					return nil, fmt.Errorf("invalid capacity %q: must be a plain integer or Kubernetes Quantity (e.g. 100, 100Gi): %v", value, err)
				}
			}
			opts.Capacity = value
		case "entrypointconfig":
			opts.EntrypointConfig = value
		case "entrypointkey":
			opts.EntrypointKey = value
		case "dnspolicy":
			// Matched case-insensitively but stored as the canonical constant: the
			// value ends up on the fuse pod spec, which the API server validates
			// against the exact spelling.
			switch {
			case strings.EqualFold(value, string(corev1.DNSClusterFirst)):
				opts.DnsPolicy = corev1.DNSClusterFirst
			case strings.EqualFold(value, string(corev1.DNSClusterFirstWithHostNet)):
				opts.DnsPolicy = corev1.DNSClusterFirstWithHostNet
			case strings.EqualFold(value, string(corev1.DNSDefault)):
				opts.DnsPolicy = corev1.DNSDefault
			default:
				klog.Warningf("dnsPolicy %q is not one of ClusterFirst, ClusterFirstWithHostNet, Default; leaving it unset, which the API server defaults to ClusterFirst — and on a host-network pod that resolves through the node, not the cluster, so an in-cluster endpoint needs ClusterFirstWithHostNet", value)
			}
		}
	}

	// Extract FsType and MountFlags from VolumeCapability.
	if volCap := req.GetVolumeCapability(); volCap != nil {
		if mount := volCap.GetMount(); mount != nil {
			// fsType is the other spelling of fuseType, so where both name a client
			// they have to agree. CustomFuseType is exempt on either side: it is the
			// generic marker, carrying no client name, and the external provisioner
			// stamps it onto every dynamically provisioned volume. Read as a
			// declaration it would contradict whichever client the volume asked for.
			if mount.FsType != "" && mount.FsType != mounterutils.CustomFuseType {
				if opts.FuseType != "" && opts.FuseType != mounterutils.CustomFuseType && opts.FuseType != mount.FsType {
					return nil, fmt.Errorf("fuseType %q from volumeAttributes conflicts with fsType %q from PV spec", opts.FuseType, mount.FsType)
				}
				opts.FuseType = mount.FsType
			}
			if len(mount.MountFlags) > 0 {
				// Checked as well, because an "otherOpts=..." entry here travels the
				// same path into the entrypoint's environment as the attribute above,
				// so rejecting one and not the other would just move the payload.
				for _, flag := range mount.MountFlags {
					if err := validateShellSafe("mountOptions entry", flag); err != nil {
						return nil, err
					}
				}
				opts.MountOptions = mount.MountFlags
			}
		}
		switch volCap.AccessMode.GetMode() {
		case csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY, csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY:
			opts.ReadOnly = true
		}
	}
	if req.GetReadonly() {
		opts.ReadOnly = true
	}

	// After readOnly is final, so a mountOptions entry cannot contradict the
	// access mode.
	if err := applyMountOptions(opts); err != nil {
		return nil, err
	}

	return opts, nil
}

// precheckAuthConfig validates the auth configuration before creating the fuse pod.
//
// Only what this side can actually decide is checked. Whether the node can reach
// the credential provider, and whether the sandbox token is present, is known
// only in mount-proxy, which resolves those from its own environment and reports
// them per mount; re-checking here would couple the controller to the node's
// configuration and reject volumes that would in fact mount.
func precheckAuthConfig(opts *fuseOptions) error {
	switch opts.AuthType {
	case "":
		return nil
	case jwtauth.AuthTypeAgentIdentity:
		if opts.SandboxId == "" {
			return fmt.Errorf("authType %s requires sandboxId in volume attributes", opts.AuthType)
		}
		if opts.SandboxCredProviderName == "" {
			return fmt.Errorf("authType %s requires sandboxCredProviderName in volume attributes", opts.AuthType)
		}
		if opts.CredentialRefreshHookKey != "" && opts.EntrypointConfig == "" {
			return fmt.Errorf("credentialRefreshHookKey names a key in entrypointConfig, which is not set; a hook shipped in the image needs neither")
		}
		// A relative path would resolve against a working directory neither the
		// volume author nor the entrypoint can see.
		if opts.CredentialDir != "" && !filepath.IsAbs(opts.CredentialDir) {
			return fmt.Errorf("credentialDir %q must be an absolute path", opts.CredentialDir)
		}
		return nil
	default:
		return fmt.Errorf("unsupported authType %q; supported: default (secret passthrough), %s", opts.AuthType, jwtauth.AuthTypeAgentIdentity)
	}
}

// makeAuthConfig constructs the AuthConfig for fuse pod creation.
// The auth type determines how credentials are provisioned to the FUSE entrypoint:
//   - "" (default): secrets are passed as env vars directly (key=value, no transformation)
//   - "agent-identity": mount-proxy exchanges the sandbox token for a scoped STS
//     credential, delivers it as files under credentialDir, and refreshes it
//     before it expires
//   - "rrsa": (planned) RRSA-based auth via mount-proxy
func makeAuthConfig(opts *fuseOptions) *fpm.AuthConfig {
	authCfg := &fpm.AuthConfig{
		AuthType: opts.AuthType,
	}
	if len(opts.Secrets) > 0 {
		authCfg.Secrets = opts.Secrets
	}
	return authCfg
}

// makeMountOptions serializes volume attributes as key=value pairs to be
// carried through the CSI mount protocol to mount-proxy. Unlike OSS where
// options are actual FUSE mount flags, here they are opaque transport for
// volume attributes — mount-proxy's buildEnvVars maps them to env vars
// for the entrypoint.
//
// MountOptions is appended last, but the position no longer decides anything:
// applyMountOptions has already removed every entry that named a field above,
// so what is left cannot collide with them. Each remaining entry becomes its
// own env var, spelled exactly as the PV wrote it.
//
// Note: Source is not included here — it is passed via MountOperation.Source
// and mapped to $source by mount-proxy directly.
func (o *fuseOptions) makeMountOptions() []string {
	var opts []string
	if o.Bucket != "" {
		opts = append(opts, "bucket="+o.Bucket)
	}
	if o.URL != "" {
		opts = append(opts, "url="+o.URL)
	}
	if o.Path != "" {
		opts = append(opts, "path="+o.Path)
	}
	if o.ReadOnly {
		opts = append(opts, "readOnly=true")
	}
	if o.OtherOpts != "" {
		opts = append(opts, "otherOpts="+o.OtherOpts)
	}
	if o.Capacity != "" {
		opts = append(opts, "capacity="+o.Capacity)
	}
	if o.AuthType != "" {
		opts = append(opts, jwtauth.OptAuthType+"="+o.AuthType)
	}
	// The agent-identity settings travel as options because mount-proxy, not
	// this side, performs the credential exchange. It strips them again before
	// the entrypoint runs, so they never reach the FUSE client.
	if o.SandboxId != "" {
		opts = append(opts, jwtauth.OptSandboxId+"="+o.SandboxId)
	}
	if o.SandboxCredProviderName != "" {
		opts = append(opts, jwtauth.OptSandboxCredProviderName+"="+o.SandboxCredProviderName)
	}
	if o.CredentialDir != "" {
		opts = append(opts, interceptors.OptCredentialDir+"="+o.CredentialDir)
	}
	opts = append(opts, o.MountOptions...)
	return opts
}

// metricsBucketName is the bucket_name label on the volume's fuse_stat metrics.
func (o *fuseOptions) metricsBucketName() string {
	if o.Bucket == "" {
		return unknownBucketName
	}
	return o.Bucket
}
