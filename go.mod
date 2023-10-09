module github.com/kubernetes-sigs/alibaba-cloud-csi-driver

go 1.18

require (
	github.com/alibabacloud-go/darabonba-openapi v0.1.16
	github.com/alibabacloud-go/ens-20171110/v3 v3.0.2
	github.com/alibabacloud-go/tea v1.1.17
	github.com/aliyun/alibaba-cloud-sdk-go v1.62.401
	github.com/container-storage-interface/spec v1.1.0
	github.com/containerd/ttrpc v1.1.0
	github.com/docker/docker v0.7.3-0.20190327010347-be7ac8be2ae0
	github.com/emirpasic/gods v1.12.0
	github.com/go-ping/ping v0.0.0-20201022122018-3977ed72668a
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/jarcoal/httpmock v1.3.1
	github.com/kubernetes-csi/drivers v1.0.2
	github.com/kubernetes-csi/external-snapshotter/client/v4 v4.2.0
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.1
	github.com/prometheus/common v0.33.0
	github.com/prometheus/procfs v0.7.3
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/sevlyar/go-daemon v0.1.5
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	go.uber.org/ratelimit v0.1.0
	golang.org/x/sys v0.7.0
	google.golang.org/grpc v1.31.0
	google.golang.org/protobuf v1.28.1
	gopkg.in/h2non/gock.v1 v1.1.2
	k8s.io/api v0.19.0
	k8s.io/apiextensions-apiserver v0.18.8
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v0.19.0
	k8s.io/mount-utils v0.21.1
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920
)

require (
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/alibabacloud-go/alibabacloud-gateway-spi v0.0.4 // indirect
	github.com/alibabacloud-go/debug v0.0.0-20190504072949-9472017b5c68 // indirect
	github.com/alibabacloud-go/endpoint-util v1.1.0 // indirect
	github.com/alibabacloud-go/openapi-util v0.0.11 // indirect
	github.com/alibabacloud-go/tea-utils v1.4.3 // indirect
	github.com/aliyun/credentials-go v1.1.2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/distribution v0.0.0-20170905204447-5db89f0ca686 // indirect
	github.com/docker/go-connections v0.3.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/evanphx/json-patch v4.9.0+incompatible // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/google/go-cmp v0.5.5 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/googleapis/gnostic v0.4.1 // indirect
	github.com/gorilla/mux v1.7.0 // indirect
	github.com/h2non/parth v0.0.0-20190131123155-b4df798d6542 // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/kubernetes-csi/csi-lib-utils v0.7.1 // indirect
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opentracing/opentracing-go v1.2.1-0.20220228012449-10b1cf09e00b // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tjfoc/gmsm v1.3.2 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/oauth2 v0.0.0-20220223155221-ee480838109b // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/time v0.0.0-20220210224613-90d013bbcef8 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20200825200019-8632dd797987 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/klog v1.0.0 // indirect
	k8s.io/klog/v2 v2.80.1 // indirect
	k8s.io/kube-openapi v0.0.0-20200805222855-6aeccd4b50c6 // indirect
	sigs.k8s.io/structured-merge-diff/v3 v3.0.0 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	github.com/beorn7/perks => github.com/beorn7/perks v0.0.0-20180321164747-3a771d992973
	github.com/evanphx/json-patch => github.com/evanphx/json-patch v3.0.0+incompatible
	github.com/google/go-cmp => github.com/google/go-cmp v0.5.0
	github.com/google/gofuzz => github.com/google/gofuzz v0.0.0-20170612174753-24818f796faf
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.2.0
	github.com/json-iterator/go => github.com/json-iterator/go v1.1.5
	github.com/modern-go/reflect2 => github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742
	github.com/pkg/errors => github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.11.1
	github.com/spf13/pflag => github.com/spf13/pflag v1.0.3
	github.com/stretchr/testify => github.com/stretchr/testify v1.3.0
	google.golang.org/appengine => google.golang.org/appengine v1.4.0
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20190307195333-5fe7a883aa19
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.2
	k8s.io/api => k8s.io/api v0.18.8
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.6
	k8s.io/client-go => k8s.io/client-go v0.18.6
	k8s.io/klog => k8s.io/klog v0.2.0
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20190306001800-15615b16d372
	k8s.io/utils => k8s.io/utils v0.0.0-20200324210504-a9aa75ae1b89
	sigs.k8s.io/yaml => sigs.k8s.io/yaml v1.1.0
)
