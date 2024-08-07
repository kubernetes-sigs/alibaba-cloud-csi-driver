package disk

type Category string
type PerformanceLevel string

const (
	DiskCommon     Category = "cloud"
	DiskEfficiency Category = "cloud_efficiency"
	DiskSSD        Category = "cloud_ssd"
	DiskESSD       Category = "cloud_essd"
	DiskESSDAuto   Category = "cloud_auto"
	DiskESSDEntry  Category = "cloud_essd_entry"

	DiskPPerf            Category = "cloud_pperf"
	DiskSPerf            Category = "cloud_sperf"
	DiskSharedSSD        Category = "san_ssd"
	DiskSharedEfficiency Category = "san_efficiency"

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
	Size                  SizeRange
	PerformanceLevel      map[PerformanceLevel]PerformanceLevelDesc
	InstantAccessSnapshot bool
	ProvisionedIops       bool
	Bursting              bool
}

var AllCategories = map[Category]CategoryDesc{
	DiskCommon: {},
	DiskEfficiency: {
		Size: SizeRange{Min: 20, Max: 65536},
	},
	DiskSSD: {
		Size: SizeRange{Min: 20, Max: 65536},
	},
	DiskESSD: {
		Size: SizeRange{Min: 20, Max: 65536},
		PerformanceLevel: map[PerformanceLevel]PerformanceLevelDesc{
			PERFORMANCE_LEVEL0: {Size: SizeRange{Min: 1, Max: 65536}},
			PERFORMANCE_LEVEL1: {Size: SizeRange{Min: 20, Max: 65536}},
			PERFORMANCE_LEVEL2: {Size: SizeRange{Min: 461, Max: 65536}},
			PERFORMANCE_LEVEL3: {Size: SizeRange{Min: 1261, Max: 65536}},
		},
		InstantAccessSnapshot: true,
	},
	DiskESSDAuto: {
		Size:                  SizeRange{Min: 1, Max: 65536},
		InstantAccessSnapshot: true,
		ProvisionedIops:       true,
		Bursting:              true,
	},
	DiskESSDEntry: {
		Size: SizeRange{Min: 10, Max: 65536},
	},

	// Deprecated shared disk
	DiskSharedSSD:        {},
	DiskSharedEfficiency: {},

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
