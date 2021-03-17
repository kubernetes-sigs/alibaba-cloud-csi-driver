package mem

import (
	"k8s.io/client-go/kubernetes"
)

// GlobalConfig var
type GlobalConfig struct {
	Region              string
	NodeID              string
	Scheduler           string
	BaseDir             string
	KmemEnable          bool
	ControllerProvision bool
	KubeClient          *kubernetes.Clientset
}

var (
	// GlobalConfigVar var
	GlobalConfigVar GlobalConfig
)

const (
	// MemDriverName tag
	MemDriverName = "mem.csi.alibabacloud.com"
	// KmemNodeLabel ...
	KmemNodeLabel = "pmem.csi.alibabacloud.com/type"
	// KmemValue ...
	KmemValue = "KMEM"
)

// PmemRegions ...
type PmemRegions struct {
	Regions []PmemRegion `json:"regions"`
}

// PmemRegion define on pmem region
type PmemRegion struct {
	Dev               string          `json:"dev"`
	Size              int64           `json:"size,omitempty"`
	AvailableSize     int64           `json:"available_size,omitempty"`
	MaxAvailableExent int64           `json:"max_available_extent,omitempty"`
	RegionType        string          `json:"type,omitempty"`
	IsetID            int64           `json:"iset_id,omitempty"`
	PersistenceDomain string          `json:"persistence_domain,omitempty"`
	Namespaces        []PmemNameSpace `json:"namespaces,omitempty"`
}

// PmemNameSpace define one pmem namespaces
type PmemNameSpace struct {
	Dev        string `json:"dev,omitempty"`
	Mode       string `json:"mode,omitempty"`
	MapType    string `json:"map,omitempty"`
	Size       int64  `json:"size,omitempty"`
	UUID       string `json:"uuid,omitempty"`
	SectorSize int64  `json:"sectorsize,omitempty"`
	Align      int64  `json:"align,omitempty"`
	BlockDev   string `json:"blockdev,omitempty"`
	CharDev    string `json:"chardev,omitempty"`
	Name       string `json:"name,omitempty"`
}

// DaxctrlMem list all mems
type DaxctrlMem struct {
	Chardev    string `json:"chardev"`
	Size       int64  `json:"size"`
	TargetNode int    `json:"target_node"`
	Mode       string `json:"mode"`
	Movable    bool   `json:"movable"`
}
