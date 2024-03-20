package cloud

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/sets"
)

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

var regionsCanUseVpcEndpoint = sets.NewString(
	"cn-hangzhou",
	"cn-zhangjiakou",
	"cn-shanghai",
	"cn-hongkong",
	"us-west-1",
	"eu-west-1",
	"ap-southeast-5",
	"cn-huhehaote",
	"cn-beijing",
	"cn-shenzhen",
	"eu-central-1",
	"ap-southeast-3",
	"ap-south-1",
	"ap-southeast-2",
	"cn-north-2-gov-1",
	"ap-southeast-1",
	"us-east-1",
	"cn-chengdu",
	"cn-wulanchabu",
	"cn-guangzhou",
	"ap-northeast-1",
	"cn-heyuan-acdr-1",
	"cn-beijing-finance-1",
	"ap-southeast-7",
	"cn-heyuan",
	"me-central-1",
	"ap-northeast-2",
	"cn-hangzhou-finance",
	"cn-qingdao",
	"cn-shanghai-finance-1",
	"cn-shenzhen-finance-1",
	"ap-southeast-6",
)

func GetEndpointForRegion(region string) string {
	if regionsCanUseVpcEndpoint.Has(region) {
		return fmt.Sprintf("nas-vpc.%s.aliyuncs.com", region)
	}
	endpoint := fmt.Sprintf("nas.%s.aliyuncs.com", region)
	logrus.Warnf("region %s do not support vpc endpoint, use public network endpoint: %s", region, endpoint)
	return endpoint
}
