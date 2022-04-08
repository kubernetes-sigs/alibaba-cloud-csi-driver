package options

import (
	"flag"
)

var (
	// Kubeconfig is the path to kubeconfig file
	Kubeconfig string
	// MasterURL is the url of kube-apiserver
	MasterURL string
)

func init() {
	flag.StringVar(&Kubeconfig, "kubeconfig", "", "the path to kubeconfig file")
	flag.StringVar(&MasterURL, "master-url", "", "the url of kube-apiserver")
}
