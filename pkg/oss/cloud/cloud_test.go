package cloud

import (
	"context"
	"testing"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	gomock "github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/stretchr/testify/assert"
	"k8s.io/utils/ptr"
)

func TestGetOSSRegion(t *testing.T) {
	tests := []struct {
		name       string
		location   *string
		wantRegion string
	}{
		{
			name:       "normal region",
			location:   ptr.To("oss-cn-beijing"),
			wantRegion: "cn-beijing",
		},
		{
			name:       "special region",
			location:   ptr.To("oss-cn-hzjbp"),
			wantRegion: "cn-hangzhou-finance",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			region := getOSSRegion(tt.location)
			assert.Equal(t, tt.wantRegion, region)
		})
	}
}

func TestGetBucketRegion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := cloud.NewMockOSSInterface(ctrl)

	result := oss.GetBucketInfoResult{
		BucketInfo: oss.BucketInfo{
			Location: ptr.To("oss-cn-beijing"),
			Name:     ptr.To("test"),
		},
	}
	c.EXPECT().GetBucketInfo(gomock.Any(), gomock.Any()).Return(&result, nil)

	fakeClient := &OSSClient{
		Ossc:          c,
		ClusterRegion: "cn-beijing",
	}
	region, err := fakeClient.GetBucketRegion(context.Background(), "test")
	assert.NoError(t, err)
	assert.Equal(t, "cn-beijing", region)
}
