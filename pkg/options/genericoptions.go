package options

import (
	"os"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/flowcontrol"
)

var (
	// Kubeconfig is the path to kubeconfig file
	Kubeconfig string
	// MasterURL is the url of kube-apiserver
	MasterURL string
	// QPS to use while communicating with the kubernetes apiserver. Defaults to 5.0.
	KubeAPIQPS float32
	// Burst to use while communicating with the kubernetes apiserver. Defaults to 10.
	KubeAPIBurst int
)

func init() {
	flag.StringVar(&Kubeconfig, "kubeconfig", "", "the path to kubeconfig file")
	flag.StringVar(&MasterURL, "master-url", "", "the url of kube-apiserver")
	flag.Float32Var(&KubeAPIQPS, "kube-api-qps", 5.0, "QPS to use while communicating with the kubernetes apiserver.")
	flag.IntVar(&KubeAPIBurst, "kube-api-burst", 10, "Burst to use while communicating with the kubernetes apiserver.")
}

func MustGetRestConfig() *rest.Config {
	config, err := GetRestConfig()
	if err != nil {
		panic(err)
	}
	return config
}

func GetRestConfig() (*rest.Config, error) {
	cfg, err := clientcmd.BuildConfigFromFlags(MasterURL, Kubeconfig)
	if err != nil {
		return cfg, err
	}

	cfg.QPS = KubeAPIQPS
	cfg.Burst = KubeAPIBurst
	// To ensure that the previous configuration takes effect,
	// the ENV parameter is the first priority.
	if qps := os.Getenv("KUBE_CLI_API_QPS"); qps != "" {
		if qpsi, err := strconv.Atoi(qps); err == nil {
			cfg.QPS = float32(qpsi)
		}
	}
	if burst := os.Getenv("KUBE_CLI_API_BURST"); burst != "" {
		if qpsi, err := strconv.Atoi(burst); err == nil {
			cfg.Burst = qpsi
		}
	}
	if cfg.QPS > 0 {
		// Init a RateLimiter to ensure that all clients using this Config share the same rate limiter instance
		cfg.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(cfg.QPS, cfg.Burst)
	}

	cfg.AcceptContentTypes = strings.Join([]string{runtime.ContentTypeProtobuf, runtime.ContentTypeJSON}, ",")
	cfg.ContentType = runtime.ContentTypeProtobuf
	return cfg, nil
}
