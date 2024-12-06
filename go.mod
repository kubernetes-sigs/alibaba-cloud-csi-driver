module github.com/kubernetes-sigs/alibaba-cloud-csi-driver

go 1.22.9

require (
	github.com/alibabacloud-go/darabonba-openapi v0.1.16
	github.com/alibabacloud-go/darabonba-openapi/v2 v2.0.6
	github.com/alibabacloud-go/ens-20171110/v3 v3.0.2
	github.com/alibabacloud-go/nas-20170626/v3 v3.2.0
	github.com/alibabacloud-go/tea v1.2.1
	github.com/aliyun/alibaba-cloud-sdk-go v1.62.719
	github.com/aliyun/credentials-go v1.3.1
	github.com/container-storage-interface/spec v1.10.0
	github.com/containerd/ttrpc v1.2.3
	github.com/go-logr/logr v1.4.2
	github.com/go-ping/ping v0.0.0-20201022122018-3977ed72668a
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.4
	github.com/google/go-cmp v0.6.0
	github.com/jarcoal/httpmock v1.3.1
	github.com/kubernetes-csi/csi-lib-utils v0.7.1
	github.com/kubernetes-csi/external-snapshotter/client/v8 v8.0.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.19.1
	github.com/prometheus/common v0.55.0
	github.com/prometheus/procfs v0.15.1
	github.com/sirupsen/logrus v1.9.0
	github.com/stretchr/testify v1.9.0
	go.uber.org/ratelimit v0.1.0
	golang.org/x/sys v0.21.0
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
	gopkg.in/h2non/gock.v1 v1.1.2
	k8s.io/api v0.31.3
	k8s.io/apimachinery v0.31.3
	k8s.io/client-go v0.31.3
	k8s.io/component-base v0.31.3
	k8s.io/klog/v2 v2.130.1
	k8s.io/kubelet v0.26.12
	k8s.io/mount-utils v0.21.1
	k8s.io/utils v0.0.0-20240711033017-18e509b52bc8
)

require (
	github.com/alibabacloud-go/alibabacloud-gateway-spi v0.0.4 // indirect
	github.com/alibabacloud-go/debug v0.0.0-20190504072949-9472017b5c68 // indirect
	github.com/alibabacloud-go/endpoint-util v1.1.0 // indirect
	github.com/alibabacloud-go/openapi-util v0.1.0 // indirect
	github.com/alibabacloud-go/tea-utils v1.4.3 // indirect
	github.com/alibabacloud-go/tea-utils/v2 v2.0.5 // indirect
	github.com/alibabacloud-go/tea-xml v1.1.3 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/clbanning/mxj/v2 v2.5.5 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/fxamacker/cbor/v2 v2.7.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/h2non/parth v0.0.0-20190131123155-b4df798d6542 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/opentracing/opentracing-go v1.2.1-0.20220228012449-10b1cf09e00b // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tjfoc/gmsm v1.3.2 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/oauth2 v0.21.0 // indirect
	golang.org/x/term v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240701130421-f6361c86f094 // indirect
	gopkg.in/evanphx/json-patch.v4 v4.12.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/kube-openapi v0.0.0-20240228011516-70dd3763d340 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
