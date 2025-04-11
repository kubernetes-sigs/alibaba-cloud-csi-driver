package common

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"google.golang.org/grpc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

// Servers holds the list of servers.
type Servers struct {
	IdentityServer        csi.IdentityServer
	ControllerServer      csi.ControllerServer
	NodeServer            csi.NodeServer
	GroupControllerServer csi.GroupControllerServer
}

func ParseEndpoint(ep string) (string, string, error) {
	if strings.HasPrefix(strings.ToLower(ep), "unix://") || strings.HasPrefix(strings.ToLower(ep), "tcp://") {
		s := strings.SplitN(ep, "://", 2)
		if s[1] != "" {
			return s[0], s[1], nil
		}
	}
	return "", "", fmt.Errorf("Invalid endpoint: %v", ep)
}

func RunCSIServer(driverType, endpoint string, servers Servers) {
	config := options.MustGetRestConfig()
	clientset := kubernetes.NewForConfigOrDie(config)

	proto, addr, err := ParseEndpoint(endpoint)
	if err != nil {
		klog.Fatal(err.Error())
	}

	if proto == "unix" {
		addr = "/" + addr
		if err := os.Remove(addr); err != nil && !os.IsNotExist(err) {
			klog.Fatalf("Failed to remove %s, error: %s", addr, err.Error())
		}
	}

	listener, err := net.Listen(proto, addr)
	if err != nil {
		klog.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(logGRPC, instrumentGRPC(driverType)),
	}
	server := grpc.NewServer(opts...)

	if servers.IdentityServer != nil {
		csi.RegisterIdentityServer(server, servers.IdentityServer)
	}
	if servers.ControllerServer != nil {
		csi.RegisterControllerServer(server, WrapControllerServerWithValidator(servers.ControllerServer))
	}
	if servers.NodeServer != nil {
		csi.RegisterNodeServer(server, WrapNodeServerWithValidator(WrapNodeServerWithMetricRecorder(servers.NodeServer, driverType, clientset)))
	}
	if servers.GroupControllerServer != nil {
		csi.RegisterGroupControllerServer(server, WrapGroupControllerServerWithValidator(servers.GroupControllerServer))
	}

	klog.Infof("Listening for connections on address: %#v", listener.Addr())

	err = server.Serve(listener)
	if err != nil {
		klog.Fatalf("Failed to serve: %v", err)
	}
}
