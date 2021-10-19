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
	"github.com/google/go-microservice-helpers/server"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/lib"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"os"
)

const (
	// LvmdPort is lvm daemon tcp port
	LvmdPort = "1736"
)

var KubeClient = &kubernetes.Clientset{}

// Start start lvmd
func Start(kubeClient *kubernetes.Clientset) {
	address := "0.0.0.0:" + GetLvmdPort()
	log.Infof("Lvmd Starting with socket: %s ...", address)

	KubeClient = kubeClient
	svr := NewServer()
	serverhelpers.ListenAddress = &address
	grpcServer, _, err := serverhelpers.NewServer()
	if err != nil {
		log.Errorf("failed to init GRPC server: %v", err)
		return
	}

	projQuotaServer := NewProjQuotaServer()
	lib.RegisterLVMServer(grpcServer, &svr)
	lib.RegisterProjQuotaServer(grpcServer, &projQuotaServer)

	err = serverhelpers.ListenAndServe(grpcServer, nil)
	if err != nil {
		log.Errorf("failed to serve: %v", err)
		return
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
