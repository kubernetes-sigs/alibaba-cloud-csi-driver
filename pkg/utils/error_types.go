package utils

import "strings"

const (
	// DiskAttachDetach ...
	DiskAttachDetach = "diskAttachDetach"
	// DiskProvision ...
	DiskProvision = "diskProvision"
	// DiskMount ...
	DiskMount = "diskMount"
	// DiskDelete ...
	DiskDelete = "diskDelete"
	// NasFilesystemCreate ...
	NasFilesystemCreate = "nasFilesystemCreate"
	// NasFilesystemDelete ...
	NasFilesystemDelete = "nasFilesystemDelete"
	// NasMountTargetCreate ...
	NasMountTargetCreate = "nasMountTargetCreate"
	// NasMountTargetDelete ...
	NasMountTargetDelete = "nasMountTargetDelete"
)

type errorInfo struct {
	errorName           string
	errorTypicalMessage string
	errorSuggestion     string
}

var errorTypeMaps = map[string]map[string]*errorInfo{
	DiskAttachDetach:     DiskAttachDetachErrors,
	DiskProvision:        DiskProvisionErrors,
	DiskMount:            DiskMountErrors,
	DiskDelete:           DiskDeleteErrors,
	NasFilesystemCreate:  NasFilesystemCreateErrors,
	NasFilesystemDelete:  NasFilesystemDeleteErrors,
	NasMountTargetCreate: NasMountTargetCreateErrors,
	NasMountTargetDelete: NasMountTargetDeleteErrors,
}

// DiskProvisionErrors are errors throwed by ecs create disk api
var DiskProvisionErrors = map[string]*errorInfo{
	"disk size is not supported.": {
		errorName:           "SpecificSizeNotValid",
		errorTypicalMessage: "disk size is not supported.",
		errorSuggestion:     "faq：https://help.aliyun.com/document_detail/286495.htm?spm=a2c4g.11186623.0.0.77bc690cSyWVYM#section-r6k-e7i-l78",
	},
	"The specified AZone inventory is insufficient": {
		errorName:           "InsufficientInventory",
		errorTypicalMessage: "The specified AZone inventory is insufficient",
		errorSuggestion:     "faq：https://help.aliyun.com/document_detail/286495.htm?spm=a2c4g.11186623.0.0.77bc690clAAUic#section-6n3-8ue-fyf",
	},
}

// DiskAttachDetachErrors are errors throwed by disk attach
var DiskAttachDetachErrors = map[string]*errorInfo{
	"had volume node affinity conflict": {
		errorName:           "VolumeNodeAffinityConflict",
		errorTypicalMessage: "had volume node affinity conflict",
		errorSuggestion:     "faq: https://help.aliyun.com/document_detail/286495.htm?spm=a2c4g.11186623.0.0.77bc690cSyWVYM#section-7o2-zf7-eff",
	},
	"can't find disk:": {
		errorName:           "VolumeDiskIDError",
		errorTypicalMessage: "can't find disk:",
		errorSuggestion:     "faq: https://help.aliyun.com/document_detail/286495.htm?spm=a2c4g.11186623.0.0.77bc690cSyWVYM#section-7o2-zf7-eff",
	},
	"instance does not support this disk category.": {
		errorName:           "NotSupportDiskCategory",
		errorTypicalMessage: "instance does not support this disk category.",
		errorSuggestion:     "faq: https://help.aliyun.com/document_detail/286495.htm?spm=a2c4g.11186623.0.0.77bc690cSyWVYM#section-ihn-gds-9mm",
	},
	"The specified disk is not a portable disk.": {
		errorName:           "DiskNotPortable",
		errorTypicalMessage: "The specified disk is not a portable disk.",
		errorSuggestion:     "faq: https://help.aliyun.com/document_detail/286495.htm?spm=a2c4g.11186623.0.0.77bc690cSyWVYM#section-s48-866-q58",
	},
}

// DiskMountErrors are errors throwed by disk mount
var DiskMountErrors = map[string]*errorInfo{}

// DiskDeleteErrors are errors throwed by disk delete
var DiskDeleteErrors = map[string]*errorInfo{}

// NasFilesystemCreateErrors are errors throwed by nas create
var NasFilesystemCreateErrors = map[string]*errorInfo{}

// NasFilesystemDeleteErrors are errors throwed by nas filesystem delete
var NasFilesystemDeleteErrors = map[string]*errorInfo{}

// NasMountTargetCreateErrors are errors throwed by nas mount target create
var NasMountTargetCreateErrors = map[string]*errorInfo{}

// NasMountTargetDeleteErrors are errors throwed by nas mount target delete
var NasMountTargetDeleteErrors = map[string]*errorInfo{}

// FindSuggestionByErrorMessage get new error message by error type & error message
func FindSuggestionByErrorMessage(errMsg, errorType string) string {
	newErrMsg := errMsg
	if errType, ok := errorTypeMaps[errorType]; ok {
		for typicalMsg, info := range errType {
			if strings.Contains(errMsg, typicalMsg) {
				newErrMsg = newErrMsg + "\n" + info.errorSuggestion
				return newErrMsg
			}
		}
	}
	return newErrMsg
}
