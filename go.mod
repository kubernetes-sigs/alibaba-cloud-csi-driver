module github.com/kubernetes-sigs/alibaba-cloud-csi-driver

go 1.12

require (
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.751
	github.com/container-storage-interface/spec v1.2.0
	github.com/docker/docker v0.7.3-0.20190327010347-be7ac8be2ae0
	github.com/emirpasic/gods v1.12.0
	github.com/go-ping/ping v0.0.0-20201022122018-3977ed72668a
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.4.1
	github.com/google/credstore v0.0.0-20181218150457-e184c60ef875 // indirect
	github.com/google/go-microservice-helpers v0.0.0-20190205165657-a91942da5417
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/googleapis/gnostic v0.5.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/kubernetes-csi/csi-lib-utils v0.7.1 // indirect
	github.com/kubernetes-csi/drivers v1.0.2
	github.com/kubernetes-csi/external-snapshotter/client/v4 v4.2.0
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.0.0
	github.com/prometheus/common v0.4.1
	github.com/prometheus/procfs v0.0.2
	github.com/sevlyar/go-daemon v0.1.5
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	go.uber.org/ratelimit v0.1.0
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/net v0.0.0-20200904194848-62affa334b73
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/sys v0.0.0-20200821140526-fda516888d29 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/grpc v1.28.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.3.0 // indirect
	k8s.io/api v0.19.0
	k8s.io/apiextensions-apiserver v0.0.0
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v0.19.0
	k8s.io/kube-openapi v0.0.0-20200811211545-daf3cbb84823 // indirect
	k8s.io/kubernetes v1.18.8
	k8s.io/utils v0.0.0-20200324210504-a9aa75ae1b89
)

replace (
	cloud.google.com/go => cloud.google.com/go v0.34.0
	contrib.go.opencensus.io/exporter/ocagent => contrib.go.opencensus.io/exporter/ocagent v0.2.0
	git.apache.org/thrift.git => git.apache.org/thrift.git v0.0.0-20180902110319-2566ecd5d999
	github.com/Azure/azure-pipeline-go => github.com/Azure/azure-pipeline-go v0.2.1
	github.com/Azure/azure-sdk-for-go => github.com/Azure/azure-sdk-for-go v21.4.0+incompatible
	github.com/Azure/azure-storage-file-go => github.com/Azure/azure-storage-file-go v0.5.0
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v11.5.1+incompatible
	github.com/BurntSushi/toml => github.com/BurntSushi/toml v0.3.1
	github.com/GoogleCloudPlatform/k8s-cloud-provider => github.com/GoogleCloudPlatform/k8s-cloud-provider v0.0.0-20181220005116-f8e995905100
	github.com/PuerkitoBio/purell => github.com/PuerkitoBio/purell v1.1.1
	github.com/PuerkitoBio/urlesc => github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578
	github.com/alecthomas/template => github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc
	github.com/alecthomas/units => github.com/alecthomas/units v0.0.0-20151022065526-2efee857e7cf
	github.com/aws/aws-sdk-go => github.com/aws/aws-sdk-go v1.16.26
	github.com/beorn7/perks => github.com/beorn7/perks v0.0.0-20180321164747-3a771d992973
	github.com/census-instrumentation/opencensus-proto => github.com/census-instrumentation/opencensus-proto v0.1.0
	github.com/client9/misspell => github.com/client9/misspell v0.3.4
	github.com/container-storage-interface/spec => github.com/container-storage-interface/spec v1.1.0
	github.com/coreos/bbolt => github.com/coreos/bbolt v1.3.2
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.12+incompatible
	github.com/coreos/go-semver => github.com/coreos/go-semver v0.3.0
	github.com/coreos/go-systemd => github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e
	github.com/coreos/pkg => github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f
	github.com/davecgh/go-spew => github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go => github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dnaeon/go-vcr => github.com/dnaeon/go-vcr v1.0.1
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20170905204447-5db89f0ca686
	github.com/docker/spdystream => github.com/docker/spdystream v0.0.0-20181023171402-6480d4af844c
	github.com/elazarl/goproxy => github.com/elazarl/goproxy v0.0.0-20190410145444-c548f45dcf1d
	github.com/elazarl/goproxy/ext => github.com/elazarl/goproxy/ext v0.0.0-20190410145444-c548f45dcf1d
	github.com/emicklei/go-restful => github.com/emicklei/go-restful v2.9.3+incompatible
	github.com/evanphx/json-patch => github.com/evanphx/json-patch v3.0.0+incompatible
	github.com/fatih/camelcase => github.com/fatih/camelcase v1.0.0
	github.com/fsnotify/fsnotify => github.com/fsnotify/fsnotify v1.4.7
	github.com/ghodss/yaml => github.com/ghodss/yaml v1.0.0
	github.com/go-kit/kit => github.com/go-kit/kit v0.8.0
	github.com/go-logfmt/logfmt => github.com/go-logfmt/logfmt v0.3.0
	github.com/go-openapi/jsonpointer => github.com/go-openapi/jsonpointer v0.19.0
	github.com/go-openapi/jsonreference => github.com/go-openapi/jsonreference v0.19.0
	github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.0
	github.com/go-openapi/swag => github.com/go-openapi/swag v0.19.0
	github.com/go-stack/stack => github.com/go-stack/stack v1.8.0
	github.com/golang/glog => github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/groupcache => github.com/golang/groupcache v0.0.0-20190129154638-5b532d6fd5ef
	github.com/golang/mock => github.com/golang/mock v1.1.1
	github.com/google/btree => github.com/google/btree v1.0.0
	github.com/google/go-cmp => github.com/google/go-cmp v0.5.0
	github.com/google/gofuzz => github.com/google/gofuzz v0.0.0-20170612174753-24818f796faf
	github.com/google/uuid => github.com/google/uuid v1.1.1
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.2.0
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.0
	github.com/gregjones/httpcache => github.com/gregjones/httpcache v0.0.0-20190212212710-3befbb6ad0cc
	github.com/grpc-ecosystem/go-grpc-middleware => github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/grpc-ecosystem/go-grpc-prometheus => github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway => github.com/grpc-ecosystem/grpc-gateway v1.5.0
	github.com/hashicorp/golang-lru => github.com/hashicorp/golang-lru v0.5.1
	github.com/hpcloud/tail => github.com/hpcloud/tail v1.0.0
	github.com/imdario/mergo => github.com/imdario/mergo v0.3.7
	github.com/inconshreveable/mousetrap => github.com/inconshreveable/mousetrap v1.0.0
	github.com/jmespath/go-jmespath => github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af
	github.com/jonboulle/clockwork => github.com/jonboulle/clockwork v0.1.0
	github.com/json-iterator/go => github.com/json-iterator/go v1.1.5
	github.com/julienschmidt/httprouter => github.com/julienschmidt/httprouter v1.2.0
	github.com/kisielk/errcheck => github.com/kisielk/errcheck v1.1.0
	github.com/kisielk/gotool => github.com/kisielk/gotool v1.0.0
	github.com/konsorten/go-windows-terminal-sequences => github.com/konsorten/go-windows-terminal-sequences v1.0.1
	github.com/kr/logfmt => github.com/kr/logfmt v0.0.0-20140226030751-b84e30acd515
	github.com/kr/pretty => github.com/kr/pretty v0.1.0
	github.com/kr/pty => github.com/kr/pty v1.1.5
	github.com/kr/text => github.com/kr/text v0.1.0
	github.com/kubernetes-csi/csi-test => github.com/kubernetes-csi/csi-test v1.1.0
	github.com/kubernetes-csi/external-snapshotter => github.com/kubernetes-csi/external-snapshotter v1.0.1
	github.com/mailru/easyjson => github.com/mailru/easyjson v0.0.0-20190312143242-1de009706dbe
	github.com/marstr/guid => github.com/marstr/guid v1.1.0
	github.com/mattn/go-ieproxy => github.com/mattn/go-ieproxy v0.0.0-20190610004146-91bb50d98149
	github.com/matttproud/golang_protobuf_extensions => github.com/matttproud/golang_protobuf_extensions v1.0.1
	github.com/modern-go/concurrent => github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 => github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742
	github.com/mwitkow/go-conntrack => github.com/mwitkow/go-conntrack v0.0.0-20161129095857-cc309e4a2223
	github.com/onsi/ginkgo => github.com/onsi/ginkgo v1.7.0
	github.com/onsi/gomega => github.com/onsi/gomega v1.4.3
	github.com/opencontainers/go-digest => github.com/opencontainers/go-digest v1.0.0-rc1
	github.com/openzipkin/zipkin-go => github.com/openzipkin/zipkin-go v0.1.1
	github.com/pborman/uuid => github.com/pborman/uuid v0.0.0-20180906182336-adf5a7427709
	github.com/peterbourgon/diskv => github.com/peterbourgon/diskv v2.0.1+incompatible
	github.com/pkg/errors => github.com/pkg/errors v0.8.1
	github.com/pmezard/go-difflib => github.com/pmezard/go-difflib v1.0.0
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
	github.com/prometheus/client_model => github.com/prometheus/client_model v0.0.0-20190129233127-fd36f4220a90
	github.com/prometheus/common => github.com/prometheus/common v0.2.0
	github.com/prometheus/procfs => github.com/prometheus/procfs v0.0.0-20190306233201-d0f344d83b0c
	github.com/rogpeppe/go-charset => github.com/rogpeppe/go-charset v0.0.0-20180617210344-2471d30d28b4
	github.com/rubiojr/go-vhd => github.com/rubiojr/go-vhd v0.0.0-20160810183302-0bfd3b39853c
	github.com/satori/go.uuid => github.com/satori/go.uuid v1.2.0
	github.com/soheilhy/cmux => github.com/soheilhy/cmux v0.1.4
	github.com/spf13/afero => github.com/spf13/afero v1.2.2
	github.com/spf13/cobra => github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag => github.com/spf13/pflag v1.0.3
	github.com/stretchr/objx => github.com/stretchr/objx v0.1.1
	github.com/stretchr/testify => github.com/stretchr/testify v1.3.0
	github.com/tmc/grpc-websocket-proxy => github.com/tmc/grpc-websocket-proxy v0.0.0-20190109142713-0ad062ec5ee5
	github.com/ugorji/go => github.com/ugorji/go v1.1.4
	github.com/vmware/govmomi => github.com/vmware/govmomi v0.20.1
	github.com/xiang90/probing => github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2
	go.etcd.io/bbolt => go.etcd.io/bbolt v1.3.2
	go.opencensus.io => go.opencensus.io v0.18.0
	go.uber.org/atomic => go.uber.org/atomic v1.3.2
	go.uber.org/multierr => go.uber.org/multierr v1.1.0
	go.uber.org/zap => go.uber.org/zap v1.9.1
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20190228161510-8dd112bcdc25
	golang.org/x/exp => golang.org/x/exp v0.0.0-20190121172915-509febef88a4
	golang.org/x/lint => golang.org/x/lint v0.0.0-20190227174305-5b3e6a55c961
	golang.org/x/net => golang.org/x/net v0.0.0-20190301231341-16b79f2e4e95
	golang.org/x/oauth2 => golang.org/x/oauth2 v0.0.0-20190226205417-e64efc72b421
	golang.org/x/sync => golang.org/x/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190308023053-584f3b12f43e
	golang.org/x/text => golang.org/x/text v0.3.0
	golang.org/x/time => golang.org/x/time v0.0.0-20181108054448-85acf8d2951c
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190226205152-f727befe758c
	google.golang.org/api => google.golang.org/api v0.1.0
	google.golang.org/appengine => google.golang.org/appengine v1.4.0
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20190307195333-5fe7a883aa19
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	gopkg.in/alecthomas/kingpin.v2 => gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/check.v1 => gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
	gopkg.in/fsnotify.v1 => gopkg.in/fsnotify.v1 v1.4.7
	gopkg.in/gcfg.v1 => gopkg.in/gcfg.v1 v1.2.0
	gopkg.in/inf.v0 => gopkg.in/inf.v0 v0.9.1
	gopkg.in/square/go-jose.v2 => gopkg.in/square/go-jose.v2 v2.3.0
	gopkg.in/tomb.v1 => gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7
	gopkg.in/warnings.v0 => gopkg.in/warnings.v0 v0.1.1
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.2
	honnef.co/go/tools => honnef.co/go/tools v0.0.0-20190102054323-c2f93a96b099
	k8s.io/api => k8s.io/api v0.18.8
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.8
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.6
	k8s.io/apiserver => k8s.io/apiserver v0.18.6
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.18.6
	k8s.io/client-go => k8s.io/client-go v0.18.6
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.18.6
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.18.6
	k8s.io/code-generator => k8s.io/code-generator v0.18.6
	k8s.io/component-base => k8s.io/component-base v0.18.6
	k8s.io/cri-api => k8s.io/cri-api v0.18.6
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.18.6
	k8s.io/klog => k8s.io/klog v0.2.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.18.6

	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.18.6
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20190306001800-15615b16d372
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.18.6
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.18.6
	k8s.io/kubectl => k8s.io/kubectl v0.18.6
	k8s.io/kubelet => k8s.io/kubelet v0.18.6
	k8s.io/kubernetes => k8s.io/kubernetes v1.18.6
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20190624091455-d8621ceb9c64
	k8s.io/metrics => k8s.io/metrics v0.18.6
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.18.6
	k8s.io/utils => k8s.io/utils v0.0.0-20200324210504-a9aa75ae1b89
	sigs.k8s.io/kustomize => sigs.k8s.io/kustomize v2.0.3+incompatible
	sigs.k8s.io/yaml => sigs.k8s.io/yaml v1.1.0
	vbom.ml/util => github.com/fvbommel/util v0.0.0-20180919145318-efcd4e0f9787
)
