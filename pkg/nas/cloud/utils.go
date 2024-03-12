package cloud

import "strings"

func GetFilesystemTypeByMountTargetDomain(domain string) string {
	switch {
	case strings.HasSuffix(domain, "extreme.nas.aliyuncs.com"):
		return FilesystemTypeExtreme
	case strings.HasSuffix(domain, "nas.aliyuncs.com"):
		return FilesystemTypeStandard
	case strings.HasSuffix(domain, "cpfs.nas.aliyuncs.com"):
		return FilesystemTypeCpfs
	default:
		return ""
	}
}
