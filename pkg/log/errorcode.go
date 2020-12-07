package log

// const variables for Storage Type
const (
	TypeOSS   = "oss"
	TypeDisk  = "disk"
	TypeNAS   = "nas"
	TypeDBFS  = "dbfs"
	TypeCPFS  = "cpfs"
	TypeLVM   = "lvm"
	TypePMEM  = "pmem"
	TypeMEM   = "memory"
	TypeQuota = "quota"
)

// const variables for Error Code
const (
	StatusOK                  = "Ok"
	StatusFailed              = "Failed"
	StatusCreateVolumeFailed  = "CreateVolumeFailed"
	StatusDeleteVolumeFailed  = "DeleteVolumeFailed"
	StatusAttachVolumeFailed  = "AttachFailed"
	StatusDetachVolumeFailed  = "DetachFailed"
	StatusStageVolumeFailed   = "StageFailed"
	StatusUnstageVolumeFailed = "UnstageFailed"
	StatusMountFailed         = "MountFailed"
	StatusUmountFailed        = "UmountFailed"
	StatusBdfErr              = "BdfError"
	StatusSnapShotError       = "SnapShotError"
	StatusEcsApiErr           = "EcsApiError"
	StatusExecuCmdFailed      = "ExecCmdError"
	StatusArgsInvalid         = "ArgsInvalid"
	StatusVolumeTypeErr       = "VolumeTypeError"
	StatusNotFound            = "NotFound"
	StatusAlreadyExists       = "AlreadyExists"
	StatusInProcessing        = "InProcessing"
	StatusSocketError         = "SocketError"
	StatusTimeout             = "Timeout"
	StatusUnauthenticated     = "Unauthenticated"
	StatusUnimplemented       = "Unimplemented"
	StatusInternalError       = "InternalError"
	StatusUnknown             = "Unknown"
)
