package interceptors

import (
	"fmt"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
)

// resolveAgentIdentityOpts is the shared preamble for jwtauth interceptors:
// flatten options, index them, check authType, resolve and validate Opts.
// Returns (opts, index, true) when agent-identity is active, or zero values
// and false when the interceptor should pass through.
func resolveAgentIdentityOpts(options []string) (jwtauth.Opts, map[string]string, bool, error) {
	flat := flattenMountOptions(options)
	idx := mounterutils.IndexMountOptions(flat)
	if !jwtauth.IsAgentIdentity(idx[jwtauth.OptAuthType]) {
		return jwtauth.Opts{}, nil, false, nil
	}
	opts := jwtauth.ResolveOpts(idx)
	if err := opts.Validate(); err != nil {
		return jwtauth.Opts{}, nil, false, fmt.Errorf("jwtauth config error: %w", err)
	}
	return opts, idx, true, nil
}

// flattenMountOptions splits any comma-joined compound options into individual
// options so key lookups and rewriting operate on single "key=value" entries.
// For drivers whose options are already single-element (e.g. customfuse), this
// is a no-op.
func flattenMountOptions(options []string) []string {
	flat := make([]string, 0, len(options))
	for _, o := range options {
		for _, part := range mounterutils.SplitMountOptions(o) {
			if part = strings.TrimSpace(part); part != "" {
				flat = append(flat, part)
			}
		}
	}
	return flat
}
