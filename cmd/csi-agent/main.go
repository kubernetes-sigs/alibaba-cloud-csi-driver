package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/oss"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
	"github.com/sirupsen/logrus"
)

var mountProxySocket string

type CSIAgent interface {
	NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error)
}

func main() {
	logrus.SetOutput(os.Stderr)
	flag.StringVar(&mountProxySocket, "mount-proxy-sock", "/run/kube-agent/csi-agent.sock", "socket path of mount proxy server")
	flag.Parse()

	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("%s %s\n", version.VERSION, version.BUILDTIME)
		return
	}

	var agent CSIAgent
	switch driver := os.Getenv("CSI_DRIVER"); driver {
	case "ossplugin.csi.alibabacloud.com":
		meta := metadata.NewMetadata()
		agent = oss.NewCSIAgent(meta, mountProxySocket)
	default:
		printError(fmt.Errorf("invalid CSI_DRIVER: %q", driver))
		os.Exit(1)
	}

	switch cmd := os.Getenv("CSI_COMMAND"); cmd {
	case "NodePublishVolume":
		var req csi.NodePublishVolumeRequest
		err := json.NewDecoder(os.Stdin).Decode(&req)
		if err != nil {
			printError(err)
			os.Exit(1)
		}
		resp, err := agent.NodePublishVolume(context.Background(), &req)
		if err != nil {
			printError(err)
			os.Exit(2)
		}
		printJson(resp)
	default:
		printError(fmt.Errorf("invalid CSI_COMMAND: %q", cmd))
		os.Exit(1)
	}
}

func printJson(data any) error {
	return json.NewEncoder(os.Stdout).Encode(data)
}

func printError(err error) error {
	return printJson(struct {
		Error string `json:"error"`
	}{err.Error()})
}
