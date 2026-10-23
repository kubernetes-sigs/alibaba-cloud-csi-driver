package common

const SubstrateModeKey = CsiAlibabaCloudPrefix + "/substrate-mode"

func IsSubstrateVolumeContext(ctx map[string]string) bool {
	return ctx != nil && ctx[SubstrateModeKey] == "true"
}
