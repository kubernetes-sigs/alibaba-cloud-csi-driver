package common

import "os"

const SubstrateModeKey = CsiAlibabaCloudPrefix + "/substrate-mode"

// SubstrateActorRootDir is the host directory under which atelet creates
// per-actor volume mount points. It mirrors KubeletRootDir for path
// validation in NodeStageVolume and NodePublishVolume.
var SubstrateActorRootDir = "/var/lib/ate/actors"

func init() {
	if v := os.Getenv("SUBSTRATE_ACTOR_ROOT_DIR"); v != "" {
		SubstrateActorRootDir = v
	}
}

func IsSubstrateVolumeContext(ctx map[string]string) bool {
	return ctx != nil && ctx[SubstrateModeKey] == "true"
}
