package interceptors

import (
	"fmt"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
)

// resolveAgentIdentityOpts is the shared preamble for jwtauth interceptors:
// index options, check authType, resolve and validate Opts.
//
// The caller is responsible for any driver-specific option preprocessing
// (e.g. alinas flattens comma-joined compound options before calling this).
func resolveAgentIdentityOpts(options []string) (jwtauth.Opts, map[string]string, bool, error) {
	idx := mounterutils.IndexMountOptions(options)
	if !jwtauth.IsAgentIdentity(idx[jwtauth.OptAuthType]) {
		return jwtauth.Opts{}, nil, false, nil
	}
	opts := jwtauth.ResolveOpts(idx)
	if err := opts.Validate(); err != nil {
		return jwtauth.Opts{}, nil, false, fmt.Errorf("jwtauth config error: %w", err)
	}
	return opts, idx, true, nil
}
