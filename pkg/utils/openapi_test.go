package utils

import (
	"testing"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/stretchr/testify/assert"
	"k8s.io/utils/ptr"
)

func TestEcsEndpoint(t *testing.T) {
	// All regions that have VPC endpoint from https://api.aliyun.com/api/Ecs/2014-05-26 (Regions at top-left)
	vpcRegions := []string{
		"cn-qingdao",
		"cn-beijing-finance-1",
		"cn-beijing",
		"cn-zhangjiakou",
		"cn-huhehaote",
		"cn-wulanchabu",
		"cn-hangzhou",
		"cn-shanghai-finance-1",
		"cn-shanghai",
		"cn-nanjing",
		"cn-fuzhou",
		"cn-shenzhen-finance-1",
		"cn-shenzhen",
		"cn-heyuan",
		"cn-guangzhou",
		"cn-wuhan-lr",
		"me-east-1",
		"ap-southeast-2",
		"eu-central-1",
		"ap-southeast-6",
		"ap-northeast-2",
		"cn-heyuan-acdr-1",
		"ap-southeast-3",
		"us-east-1",
		"us-west-1",
		"us-southeast-1",
		"na-south-1",
		"ap-northeast-1",
		"me-central-1",
		"ap-southeast-7",
		"cn-zhongwei",
		"cn-chengdu",
		"ap-southeast-1",
		"ap-south-1",
		"ap-southeast-5",
		"eu-west-1",
		"cn-zhengzhou-jva",
		"cn-hongkong",
	}
	testEp := func(t *testing.T, region, ep string) {
		t.Run(region, func(t *testing.T) {
			cfg := GetEcsConfig(region)
			cfg.AccessKeyId = ptr.To("foo")
			cfg.AccessKeySecret = ptr.To("bar")

			client, err := ecs20140526.NewClient(cfg)
			assert.NoError(t, err)
			assert.Equal(t, ep, *client.Endpoint)
		})
	}
	testEp(t, "cn-hangzhou", "ecs-cn-hangzhou.aliyuncs.com") // use SDK defaults by default

	t.Run("vpc", func(t *testing.T) {
		t.Setenv("ALIBABA_CLOUD_NETWORK_TYPE", "vpc")
		for _, region := range vpcRegions {
			testEp(t, region, "ecs-vpc."+region+".aliyuncs.com")
		}
	})

	t.Run("public", func(t *testing.T) {
		t.Setenv("ALIBABA_CLOUD_NETWORK_TYPE", "public")
		for _, region := range vpcRegions {
			testEp(t, region, "ecs."+region+".aliyuncs.com")
		}
	})
}
