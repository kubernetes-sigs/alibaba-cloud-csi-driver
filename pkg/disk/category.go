package disk

type Category string
type PerformanceLevel string

const (
	DiskCommon     Category = "cloud"
	DiskEfficiency Category = "cloud_efficiency"
	DiskSSD        Category = "cloud_ssd"
	DiskESSD       Category = "cloud_essd"
	//special type as _xc* only used in financial cloud
	DiskESSDXc0   Category = "cloud_essd_xc0"
	DiskESSDXc1   Category = "cloud_essd_xc1"
	DiskESSDAuto  Category = "cloud_auto"
	DiskESSDEntry Category = "cloud_essd_entry"
	DiskRegional  Category = "cloud_regional_disk_auto"

	DiskPPerf Category = "cloud_pperf"
	DiskSPerf Category = "cloud_sperf"

	DiskEEDStandard Category = "elastic_ephemeral_disk_standard"
	DiskEEDPremium  Category = "elastic_ephemeral_disk_premium"

	PERFORMANCE_LEVEL0 PerformanceLevel = "PL0"
	PERFORMANCE_LEVEL1 PerformanceLevel = "PL1"
	PERFORMANCE_LEVEL2 PerformanceLevel = "PL2"
	PERFORMANCE_LEVEL3 PerformanceLevel = "PL3"
)

type SizeRange struct {
	Min int64
	Max int64
}

type PerformanceLevelDesc struct {
	Size SizeRange
}

type CategoryDesc struct {
	Size                    SizeRange
	PerformanceLevel        map[PerformanceLevel]PerformanceLevelDesc
	InstantAccessSnapshot   bool
	ProvisionedIops         bool
	Bursting                bool
	SingleInstance          bool // Cannot be attached to another instance once it is attached.
	Regional                bool
	SnapshotConsistentGroup bool
}

var AllCategories = map[Category]CategoryDesc{
	DiskCommon: {},
	DiskEfficiency: {
		Size: SizeRange{Min: 20, Max: 65536},
	},
	DiskSSD: {
		Size: SizeRange{Min: 20, Max: 65536},
	},
	DiskESSDXc0: {
		Size:                    SizeRange{Min: 20, Max: 65536},
		InstantAccessSnapshot:   true,
		SnapshotConsistentGroup: true,
	},
	DiskESSDXc1: {
		Size:                    SizeRange{Min: 20, Max: 65536},
		InstantAccessSnapshot:   true,
		SnapshotConsistentGroup: true,
	},
	DiskESSD: {
		Size: SizeRange{Min: 20, Max: 65536},
		PerformanceLevel: map[PerformanceLevel]PerformanceLevelDesc{
			PERFORMANCE_LEVEL0: {Size: SizeRange{Min: 1, Max: 65536}},
			PERFORMANCE_LEVEL1: {Size: SizeRange{Min: 20, Max: 65536}},
			PERFORMANCE_LEVEL2: {Size: SizeRange{Min: 461, Max: 65536}},
			PERFORMANCE_LEVEL3: {Size: SizeRange{Min: 1261, Max: 65536}},
		},
		InstantAccessSnapshot:   true,
		SnapshotConsistentGroup: true,
	},
	DiskESSDAuto: {
		Size:                    SizeRange{Min: 1, Max: 65536},
		InstantAccessSnapshot:   true,
		ProvisionedIops:         true,
		Bursting:                true,
		SnapshotConsistentGroup: true,
	},
	DiskESSDEntry: {
		Size:                    SizeRange{Min: 10, Max: 65536},
		SnapshotConsistentGroup: true,
	},
	DiskEEDStandard: {
		Size:           SizeRange{Min: 64, Max: 8192},
		SingleInstance: true,
	},
	DiskEEDPremium: {
		Size:           SizeRange{Min: 64, Max: 8192},
		SingleInstance: true,
	},
	DiskRegional: {
		Size:                  SizeRange{Min: 10, Max: 65536},
		InstantAccessSnapshot: true,
		Regional:              true,
	},

	// Only available in private cloud
	DiskPPerf: {},
	DiskSPerf: {},
}

func GetSizeRange(category Category, pl PerformanceLevel) SizeRange {
	desc := AllCategories[category]
	limit := desc.Size
	if plDesc, ok := desc.PerformanceLevel[pl]; ok {
		limit = plDesc.Size
	}
	return limit
}
