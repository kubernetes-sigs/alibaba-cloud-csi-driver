package cloud

import (
	"context"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
)

type OSSInterface interface {
	GetBucketInfo(context.Context, *oss.GetBucketInfoRequest, ...func(*oss.Options)) (*oss.GetBucketInfoResult, error)
}
