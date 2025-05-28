package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/oss"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var (
	runInECI         bool
	mountProxySocket string
)

type FileGrpcServer struct {
	input  io.Reader
	output io.Writer
	desc   *grpc.ServiceDesc
	impl   any
}

func NewStdIOGrpcServer() *FileGrpcServer {
	return &FileGrpcServer{
		input:  os.Stdin,
		output: os.Stdout,
	}
}

func (s *FileGrpcServer) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	s.desc = desc
	s.impl = impl
}

func (s *FileGrpcServer) decode(msg any) error {
	pbmsg := msg.(proto.Message)
	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	return protojson.Unmarshal(in, pbmsg)
}

func (s *FileGrpcServer) ServeOneRequest(method string) error {
	for _, desc := range s.desc.Methods {
		if desc.MethodName == method {
			resp, err := desc.Handler(s.impl, context.Background(), s.decode, nil)
			if err != nil {
				return err
			}
			out, err := protojson.Marshal(resp.(proto.Message))
			if err != nil {
				return fmt.Errorf("failed to marshal response: %v", err)
			}
			_, err = s.output.Write(out)
			if err != nil {
				return fmt.Errorf("failed to write response: %v", err)
			}
			return nil
		}
	}
	return fmt.Errorf("unknown method: %q", method)
}

func main() {
	_, ok := os.LookupEnv("KUBELET_ROOT_DIR")
	if !ok {
		// This is necessary to disable the check of sub-directory in NodeServerWithValidator.
		// CSI-agent is not invoked by kubelet, so there is no usual kubelet directory structure.
		utils.KubeletRootDir = "/"
	}

	logrus.SetOutput(os.Stderr)
	utils.AddKlogFlags(flag.CommandLine)
	flag.BoolVar(&runInECI, "eci", false, "run in ECI or not")
	flag.StringVar(&mountProxySocket, "mount-proxy-sock", "/run/kube-agent/csi-agent.sock", "socket path of mount proxy server")
	flag.Parse()

	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("%s %s\n", version.VERSION, version.GetTime())
		return
	}

	var agent csi.NodeServer
	switch driver := os.Getenv("CSI_DRIVER"); driver {
	case "test":
		agent = &fakeAgent{}
	case "ossplugin.csi.alibabacloud.com":
		meta := metadata.NewMetadata()
		agent = oss.NewCSIAgent(meta, runInECI, mountProxySocket)
	case "diskplugin.csi.alibabacloud.com":
		agent = disk.NewCSIAgent()
	case "nasplugin.csi.alibabacloud.com":
		agent = nas.NewCSIAgent(mountProxySocket)
	default:
		printError(fmt.Errorf("invalid CSI_DRIVER: %q", driver))
		os.Exit(1)
	}

	server := NewStdIOGrpcServer()
	csi.RegisterNodeServer(server, common.WrapNodeServerWithValidator(agent))
	err := server.ServeOneRequest(os.Getenv("CSI_COMMAND"))
	if err != nil {
		printError(err)
		os.Exit(1)
	}
}

func printError(err error) {
	s, _ := status.FromError(err)
	encodeErr := json.NewEncoder(os.Stdout).Encode(struct {
		Code  codes.Code `json:"code"`
		Error string     `json:"error"`
	}{
		Code:  s.Code(),
		Error: s.Message(),
	})
	if encodeErr != nil {
		fmt.Fprintf(os.Stderr, "Error: %v, failed to write error to stdout: %v", err, encodeErr)
	}
}

type fakeAgent struct {
	csi.UnimplementedNodeServer
}
