package cloud

import (
	"errors"
	"fmt"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	aliyunep "github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	sdkerrors "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	nassdk "github.com/aliyun/alibaba-cloud-sdk-go/services/nas"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/credentials"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/interfaces"
	utilshttp "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/http"
)

func init() {
	for region := range regionsCanUseVpcEndpoint {
		_ = aliyunep.AddEndpointMapping(region, "Nas", fmt.Sprintf("nas-vpc.%s.aliyuncs.com", region))
	}
}

func newNasClientV1(region string) (interfaces.NasV1Interface, error) {
	if ep := os.Getenv("NAS_ENDPOINT"); ep != "" {
		_ = aliyunep.AddEndpointMapping(region, "Nas", ep)
	}

	provider, err := credentials.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch credential: %w", err)
	}
	credential := credentials.V1ProviderAdaptor(provider)

	config := sdk.NewConfig()
	scheme := "HTTPS"
	if e := os.Getenv("ALICLOUD_CLIENT_SCHEME"); e != "" {
		scheme = e
	}
	config = config.WithScheme(scheme).WithUserAgent(KubernetesAlicloudIdentity)
	headers := utilshttp.MustParseHeaderEnv("NAS_HEADERS")
	if len(headers) > 0 {
		config.Transport = utilshttp.RoundTripperWithHeader(config.Transport, headers)
	}
	client, err := nassdk.NewClientWithOptions(region, config, credential)
	return client, err
}

func IsMountTargetNotFoundError(err error) bool {
	var serverErr *sdkerrors.ServerError
	if !errors.As(err, &serverErr) || serverErr == nil {
		return false
	}
	code := serverErr.ErrorCode()
	return code == "InvalidParam.MountTargetDomain" || code == "InvalidMountTarget.NotFound"
}
