package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
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
	case "test":
		agent = &fakeAgent{}
	case "ossplugin.csi.alibabacloud.com":
		meta := metadata.NewMetadata()
		agent = oss.NewCSIAgent(meta, mountProxySocket)
	default:
		printError(fmt.Errorf("invalid CSI_DRIVER: %q", driver))
		os.Exit(1)
	}

	var resp proto.Message

	switch cmd := os.Getenv("CSI_COMMAND"); cmd {
	case "NodePublishVolume":
		var req csi.NodePublishVolumeRequest
		err := jsonpb.Unmarshal(os.Stdin, &req)
		if err != nil {
			printError(err)
			os.Exit(1)
		}
		resp, err = agent.NodePublishVolume(context.Background(), &req)
		if err != nil {
			printError(err)
			os.Exit(2)
		}
	default:
		printError(fmt.Errorf("invalid CSI_COMMAND: %q", cmd))
		os.Exit(1)
	}

	_ = new(jsonpb.Marshaler).Marshal(os.Stdout, resp)
}

func printError(err error) {
	_ = json.NewEncoder(os.Stdout).Encode(struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}

type fakeAgent struct{}

func (fakeagent *fakeAgent) NodePublishVolume(context.Context, *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	return &csi.NodePublishVolumeResponse{}, nil
}
