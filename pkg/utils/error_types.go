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
		errorSuggestion:     "recommand: Please adjust the size of the cloud disk；\n reference：https://help.aliyun.com/document_detail/25412.html#BlockStorageQuota",
	},
}

// DiskAttachDetachErrors are errors throwed by disk attach
var DiskAttachDetachErrors = map[string]*errorInfo{
	"had volume node affinity conflict": {
		errorName:           "VolumeNodeAffinityConflict",
		errorTypicalMessage: "had volume node affinity conflict",
		errorSuggestion:     "recommand: Please check the node affinity of the pvc/pv and the pod",
	},
	"can't find disk:": {
		errorName:           "VolumeDiskIDError",
		errorTypicalMessage: "can't find disk:",
		errorSuggestion:     "recommand: Please check the specified disk exists \n",
	},
	"instance does not support this disk category.": {
		errorName:           "NotSupportDiskCategory",
		errorTypicalMessage: "instance does not support this disk category.",
		errorSuggestion:     "recommand: Please change the cloud disk type or current ecs instance type. \n reference: https://help.aliyun.com/document_detail/25378.html",
	},
	"The specified disk is not a portable disk.": {
		errorName:           "DiskNotPortable",
		errorTypicalMessage: "The specified disk is not a portable disk.",
		errorSuggestion:     "recommand: The specified disk can't be detach, please change disk type to portable on ecs console",
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
