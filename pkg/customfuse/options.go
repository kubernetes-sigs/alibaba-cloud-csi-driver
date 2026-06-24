//go:build !windows

package customfuse

import (
	"fmt"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	fpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	corev1 "k8s.io/api/core/v1"
)

type fuseOptions struct {
	// Source is the mount source passed to the FUSE entrypoint as $source.
	// Its format is opaque to the CSI driver — it can be a JuiceFS META-URL
	// (e.g. "redis://host:6379/1"), an OSS-style bucket:path, or any string
	// the entrypoint understands.
	Source string
	// Bucket is the object storage bucket name, passed as $bucket.
	// Independent of Source and URL — all three are separate env vars.
	Bucket string
	URL    string
	// OtherOpts originates from volumeAttributes.otherOpts.
	// An alternative to spec.mountOptions — both are valid styles.
	// Passed to the entrypoint as $otherOpts.
	OtherOpts string
	// Path is the sub-path within the volume, passed as $path.
	Path string
	// ReadOnly is derived from VolumeCapability or the CSI readOnly flag.
	// Passed as $readOnly to the entrypoint, which maps it to FUSE "ro" option.
	ReadOnly bool
	// FuseType identifies the FUSE client for metrics labeling (e.g. "juicefs",
	// "jindo"). Defaults to "customfuse" when not set by the user.
	// Can be set via volumeAttributes.fuseType or VolumeCapability.FsType.
	FuseType string
	// MountOptions are from pv.Spec.MountOptions (via VolumeCapability.Mount.MountFlags).
	// Each entry is a "key=value" or bare "key" string, passed as env var $key in entrypoint.
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
	// Currently only the default (empty string) is supported, which passes
	// Kubernetes Secret entries directly as environment variables to the
	// FUSE entrypoint. Planned future values: "rrsa", "agent-identity".
	AuthType string
	// Secrets holds the raw Kubernetes Secret data.
	// All entries are passed through to the FUSE entrypoint as environment
	// variables with the key as the variable name (no prefix, no transformation).
	// This means the entrypoint author controls both the Secret keys and the
	// env vars they consume — they are the same string.
	Secrets   map[string]string
	DnsPolicy corev1.DNSPolicy
}

func parseOptions(volContext, secrets map[string]string,
	volCaps []*csi.VolumeCapability, readOnly bool) (*fuseOptions, error) {
	if volContext == nil {
		volContext = map[string]string{}
	}

	opts := &fuseOptions{
		FuseType: mounterutils.CustomFuseType,
		Secrets:  secrets,
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
			opts.OtherOpts = value
		case "fusetype":
			opts.FuseType = value
		case "authtype":
			opts.AuthType = value
		case "entrypointconfig":
			opts.EntrypointConfig = value
		case "entrypointkey":
			opts.EntrypointKey = value
		case "dnspolicy":
			switch strings.ToLower(value) {
			case strings.ToLower(string(corev1.DNSClusterFirst)),
				strings.ToLower(string(corev1.DNSClusterFirstWithHostNet)),
				strings.ToLower(string(corev1.DNSDefault)):
				opts.DnsPolicy = corev1.DNSPolicy(value)
			}
		}
	}

	// Extract FsType and MountFlags from VolumeCapability.
	// FsType from PV spec.csi.fsType is an alternative to volumeAttributes.fuseType.
	// MountFlags from pv.Spec.MountOptions — each entry becomes a separate env var.
	for _, cap := range volCaps {
		if cap == nil {
			continue
		}
		if mount := cap.GetMount(); mount != nil {
			// FsType: alternative to volumeAttributes.fuseType
			if mount.FsType != "" {
				if opts.FuseType != "" && opts.FuseType != mounterutils.CustomFuseType && opts.FuseType != mount.FsType {
					// fuseType from volumeAttributes conflicts with fsType from PV spec
					return nil, fmt.Errorf("fuseType %q from volumeAttributes conflicts with fsType %q from PV spec", opts.FuseType, mount.FsType)
				} else {
					opts.FuseType = mount.FsType
				}
			}
			// MountFlags → MountOptions
			if len(mount.MountFlags) > 0 {
				opts.MountOptions = mount.MountFlags
			}
		}
		// ReadOnly from VolumeCapability access mode
		switch cap.AccessMode.GetMode() {
		case csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY, csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY:
			opts.ReadOnly = true
		}
	}
	if readOnly {
		opts.ReadOnly = true
	}

	// If "source" is not set, fall back to bucket/path for backward compatibility.
	if opts.Source == "" && opts.Bucket != "" {
		if opts.Path != "" {
			opts.Source = opts.Bucket + ":" + opts.Path
		} else {
			opts.Source = opts.Bucket
		}
	}

	return opts, nil
}

// precheckAuthConfig validates the auth configuration before creating the fuse pod.
// Currently only the default auth type (empty string) is supported, which passes
// Kubernetes Secret entries as environment variables to the FUSE entrypoint.
// Future auth types (rrsa, agent-identity) will be validated here.
func precheckAuthConfig(opts *fuseOptions) error {
	if opts.AuthType != "" {
		return fmt.Errorf("unsupported authType %q; only default (secret passthrough) is currently supported", opts.AuthType)
	}
	return nil
}

// makeAuthConfig constructs the AuthConfig for fuse pod creation.
// The auth type determines how credentials are provisioned to the FUSE entrypoint:
//   - "" (default): secrets are passed as env vars directly (key=value, no transformation)
//   - "rrsa": (planned) RRSA-based auth via mount-proxy
//   - "agent-identity": (planned) agent identity auth via mount-proxy
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
	opts = append(opts, o.MountOptions...)
	return opts
}
