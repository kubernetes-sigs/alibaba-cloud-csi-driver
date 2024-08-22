package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/client"
	"k8s.io/klog/v2"
)

func main() {
	var (
		socketPath string
	)
	flag.StringVar(&socketPath, "socket", "", "socket path")
	klog.InitFlags(nil)
	flag.Parse()

	var req proxy.MountRequest
	err := json.NewDecoder(os.Stdin).Decode(&req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data, _ := json.MarshalIndent(req, "", "\t")
	fmt.Println(string(data))

	dclient := client.NewClient(socketPath)
	resp, err := dclient.Mount(&req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data, _ = json.MarshalIndent(resp, "", "\t")
	fmt.Println(string(data))
}
