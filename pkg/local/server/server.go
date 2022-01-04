/*
Copyright 2020 The Kubernetes Authors.

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

package server

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/lib"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"k8s.io/client-go/kubernetes"
	"net"
	"os"
	"path/filepath"
	"time"
)

const (
	// LvmdPort is lvm daemon tcp port
	LvmdPort = "1736"
)

// KubeClient for resource access
var KubeClient = &kubernetes.Clientset{}

const (
	caCertFileName     = "ca_cert.pem"
	serverCertFileName = "server_cert.pem"
	serverKeyFileName  = "server_key.pem"
)

func newServerTransportCredentials(ch chan bool) credentials.TransportCredentials {
	pathWithTLS := filepath.Join(utils.MountPathWithTLS, "/local/grpc")
	utils.CreateDest(pathWithTLS)
	caFile := pathWithTLS + "/" + caCertFileName
	certFile := pathWithTLS + "/" + serverCertFileName
	keyFile := pathWithTLS + "/" + serverKeyFileName
	for {
		creds, err := utils.NewServerTLSFromFile(caFile, certFile, keyFile)
		time.Sleep(5 * time.Second)
		if err == nil {
			ch <- true
			return creds
		}
		log.Errorf("NewServerTLSFromFile is failed, err:%s", err)
	}
}

// Start start lvmd
func Start(kubeClient *kubernetes.Clientset, ch chan bool) {
	kubeNode := os.Getenv("KUBE_NODE_NAME")
	if len(kubeNode) == 0 {
		log.Fatalf("KUBE_NODE_NAME env is empty.")
	}
	address, err := utils.GetNodeAddr(kubeClient, kubeNode, GetLvmdPort())
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Infof("Lvmd Starting with socket: %s ...", address)
	KubeClient = kubeClient
	creds := newServerTransportCredentials(ch)
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	lvmServer := NewServer()
	projQuotaServer := NewProjQuotaServer()
	lib.RegisterLVMServer(grpcServer, &lvmServer)
	lib.RegisterProjQuotaServer(grpcServer, &projQuotaServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Infof("Lvmd End ...")
}

// GetLvmdPort get lvmd port
func GetLvmdPort() string {
	port := LvmdPort
	if value := os.Getenv("LVMD_PORT"); value != "" {
		port = value
	}
	return port
}
