package common

import "os"

// SubstrateActorRootDir is the host directory under which atelet creates
// per-actor volume mount points. It mirrors KubeletRootDir for path
// validation in NodeStageVolume and NodePublishVolume.
var SubstrateActorRootDir = "/var/lib/ate/actors"

func init() {
	if v := os.Getenv("SUBSTRATE_ACTOR_ROOT_DIR"); v != "" {
		SubstrateActorRootDir = v
	}
}
