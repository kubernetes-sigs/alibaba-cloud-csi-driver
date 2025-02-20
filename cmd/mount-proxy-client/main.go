package main

import (
	"encoding/json"
	"fmt"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/client"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

func main() {
	var (
		socketPath string
	)
	flag.StringVar(&socketPath, "socket", "", "socket path")
	utils.AddKlogFlags(flag.CommandLine)
	utils.AddGoFlags(flag.CommandLine)
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
