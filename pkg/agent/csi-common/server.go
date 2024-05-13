/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package csicommon

import (
	"net"
	"os"
	"sync"

	"github.com/golang/glog"
	"google.golang.org/grpc"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

// Servers holds the list of servers.
type Servers struct {
	IdentityServer        csi.IdentityServer
	ControllerServer      csi.ControllerServer
	NodeServer            csi.NodeServer
	GroupControllerServer csi.GroupControllerServer
}

// Defines Non blocking GRPC server interfaces
type NonBlockingGRPCServer interface {
	// Start services at the endpoint
	Start(endpoint string, servers Servers)
	// Waits for the service to stop
	Wait()
	// Stops the service gracefully
	Stop()
	// Stops the service forcefully
	ForceStop()
}

func NewNonBlockingGRPCServer() NonBlockingGRPCServer {
	return &nonBlockingGRPCServer{}
}

// NonBlocking server
type nonBlockingGRPCServer struct {
	wg     sync.WaitGroup
	server *grpc.Server
}

func (s *nonBlockingGRPCServer) Start(endpoint string, servers Servers) {

	s.wg.Add(1)

	go s.serve(endpoint, servers)

}

func (s *nonBlockingGRPCServer) Wait() {
	s.wg.Wait()
}

func (s *nonBlockingGRPCServer) Stop() {
	s.server.GracefulStop()
}

func (s *nonBlockingGRPCServer) ForceStop() {
	s.server.Stop()
}

func (s *nonBlockingGRPCServer) serve(endpoint string, servers Servers) {

	proto, addr, err := ParseEndpoint(endpoint)
	if err != nil {
		glog.Fatal(err.Error())
	}

	if proto == "unix" {
		addr = "/" + addr
		if err := os.Remove(addr); err != nil && !os.IsNotExist(err) {
			glog.Fatalf("Failed to remove %s, error: %s", addr, err.Error())
		}
	}

	listener, err := net.Listen(proto, addr)
	if err != nil {
		glog.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(logGRPC),
	}
	server := grpc.NewServer(opts...)
	s.server = server

	if servers.IdentityServer != nil {
		csi.RegisterIdentityServer(server, servers.IdentityServer)
	}
	if servers.ControllerServer != nil {
		csi.RegisterControllerServer(server, servers.ControllerServer)
	}
	if servers.NodeServer != nil {
		csi.RegisterNodeServer(server, servers.NodeServer)
	}
	if servers.GroupControllerServer != nil {
		csi.RegisterGroupControllerServer(server, servers.GroupControllerServer)
	}

	glog.Infof("Listening for connections on address: %#v", listener.Addr())

	server.Serve(listener)

}
