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

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/lib"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/manager"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"net"
	"strings"
	"time"
)

// Connection lvm connection interface
type Connection interface {
	GetLvm(ctx context.Context, volGroup string, volumeID string) (string, error)
	CreateLvm(ctx context.Context, opt *LVMOptions) (string, error)
	DeleteLvm(ctx context.Context, volGroup string, volumeID string) error
	CleanPath(ctx context.Context, path string) error
	Close() error
	GetNameSpace(ctx context.Context, regionName string, volumeID string) (string, error)
	CreateNameSpace(ctx context.Context, opt *NameSpaceOptions) (*manager.PmemNameSpace, error)
	DeleteNameSpace(ctx context.Context, volumeID string) error
	CreateProjQuotaSubpath(ctx context.Context, pvName, size, rootPath string) (string, string, error)
	SetSubpathProjQuota(ctx context.Context, quotaSubpath, blockSoftlimit, blockHardlimit, inodeSoftlimit, inodeHardlimit string) (string, error)
	RemoveProjQuotaSubpath(ctx context.Context, quotaSubpath string) (string, error)
	CreateLoopDevice(ctx context.Context, pvName, quotaSize string) (string, error)
	DeleteLoopDevice(ctx context.Context, pvName string) (string, error)
}

// LVMOptions lvm options
type LVMOptions struct {
	VolumeGroup string   `json:"volumeGroup,omitempty"`
	Name        string   `json:"name,omitempty"`
	Size        uint64   `json:"size,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Striping    bool     `json:"striping,omitempty"`
}

// NameSpaceOptions lvm options
type NameSpaceOptions struct {
	Region string
	Name   string
	Size   uint64
}

type workerConnection struct {
	conn             *grpc.ClientConn
	mountPathWithTLS string
}

var (
	_          Connection = &workerConnection{}
	serverName            = "csi-local-plugin-server"
)

const (
	caCertFileName     = "ca_cert.pem"
	clientCertFileName = "client_cert.pem"
	clientKeyFileName  = "client_key.pem"
)

// NewGrpcConnection lvm connection
func NewGrpcConnection(address string, timeout time.Duration, caCertFile string, clientCertFile string, clientKeyFile string) (Connection, error) {
	creds, err := utils.NewClientTLSFromFile(serverName, caCertFile, clientCertFile, clientKeyFile)
	conn, err := connect(address, timeout, creds)
	if err != nil {
		return nil, err
	}
	return &workerConnection{
		conn: conn,
	}, nil
}

func (c *workerConnection) Close() error {
	return c.conn.Close()
}

func connect(address string, timeout time.Duration, creds credentials.TransportCredentials) (*grpc.ClientConn, error) {
	log.Infof("New Connecting to %s", address)
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		grpc.WithBackoffMaxDelay(time.Second),
		grpc.WithUnaryInterceptor(logGRPC),
	}
	if strings.HasPrefix(address, "/") {
		dialOptions = append(dialOptions, grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("unix", addr, timeout)
		}))
	}
	conn, err := grpc.Dial(address, dialOptions...)

	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	for {
		if !conn.WaitForStateChange(ctx, conn.GetState()) {
			log.Warnf("Connection to %s timed out", address)
			return conn, nil // return nil, subsequent GetPluginInfo will show the real connection error
		}
		if conn.GetState() == connectivity.Ready {
			log.Warnf("Connected to %s", address)
			return conn, nil
		}
		log.Infof("Still trying to connect %s, connection is %s", address, conn.GetState())
	}
}

func (c *workerConnection) CreateLoopDevice(ctx context.Context, pvName, quotaSize string) (string, error) {
	client := lib.NewLoopDeviceClient(c.conn)
	req := lib.CreateLoopDeviceRequest{
		PvName: pvName,
		QuotaSize: quotaSize,
	}
	rsp, err := client.CreateLoopDevice(ctx, &req)
	if err != nil {
		log.Errorf("Create LoopDevice err: %v", err)
		return "", err
	}
	log.Infof("Create LoopDevice result: %v", rsp.CommandOutput)
	return rsp.CommandOutput, nil
}

func (c *workerConnection) DeleteLoopDevice(ctx context.Context, pvName string) (string, error) {
	client := lib.NewLoopDeviceClient(c.conn)
	req := lib.DeleteLoopDeviceRequest{
		PvName: pvName,
	}
	rsp, err := client.DeleteLoopDevice(ctx, &req)
	if err != nil {
		log.Errorf("Delete loopdevice err: %v", err)
		return "", err
	}
	log.Infof("Delete LoopDevice result: %s", rsp.CommandOutput)
	return rsp.CommandOutput, nil
}

func (c *workerConnection) CreateLvm(ctx context.Context, opt *LVMOptions) (string, error) {
	client := lib.NewLVMClient(c.conn)
	req := lib.CreateLVRequest{
		VolumeGroup: opt.VolumeGroup,
		Name:        opt.Name,
		Size:        opt.Size,
		Tags:        opt.Tags,
		Striping:    opt.Striping,
	}

	rsp, err := client.CreateLV(ctx, &req)
	if err != nil {
		log.Errorf("Create Lvm with error: %s", err.Error())
		return "", err
	}
	log.Infof("Create Lvm with result: %+v", rsp.CommandOutput)
	return rsp.GetCommandOutput(), nil
}

func (c *workerConnection) CreateNameSpace(ctx context.Context, opt *NameSpaceOptions) (*manager.PmemNameSpace, error) {
	client := lib.NewLVMClient(c.conn)
	req := lib.CreateNamespaceRequest{
		Name:   opt.Name,
		Size:   opt.Size,
		Region: opt.Region,
	}

	rsp, err := client.CreateNamespace(ctx, &req)
	if err != nil {
		log.Errorf("Create Lvm with error: %s", err.Error())
		return nil, err
	}
	log.Infof("Create Lvm with result: %+v", rsp.CommandOutput)

	pns := &manager.PmemNameSpace{}
	if err := json.Unmarshal([]byte(rsp.CommandOutput), pns); err != nil {
		return nil, err
	}
	return pns, nil
}

func (c *workerConnection) GetNameSpace(ctx context.Context, regionName string, volumeID string) (string, error) {
	client := lib.NewLVMClient(c.conn)
	req := lib.ListNamespaceRequest{
		NameSpace: volumeID,
		Region:    regionName,
	}

	rsp, err := client.ListNamespace(ctx, &req)
	if err != nil {
		return "", err
	}
	for _, namespace := range rsp.NameSpace {
		if namespace.Name == volumeID {
			return namespace.Name, nil
		}
	}
	return "", nil
}

func (c *workerConnection) GetLvm(ctx context.Context, volGroup string, volumeID string) (string, error) {
	client := lib.NewLVMClient(c.conn)
	req := lib.ListLVRequest{
		VolumeGroup: fmt.Sprintf("%s/%s", volGroup, volumeID),
	}

	rsp, err := client.ListLV(ctx, &req)
	if err != nil {
		log.Errorf("Get Lvm with error: %s", err.Error())
		return "", err
	}
	if len(rsp.GetVolumes()) <= 0 {
		log.Warnf("Volume %s/%s is not exist", volGroup, volumeID)
		return "", nil
	}
	log.Infof("Get Lvm with result: %+v", rsp.Volumes)
	return rsp.GetVolumes()[0].String(), nil
}

func (c *workerConnection) DeleteLvm(ctx context.Context, volGroup, volumeID string) error {
	client := lib.NewLVMClient(c.conn)
	req := lib.RemoveLVRequest{
		VolumeGroup: volGroup,
		Name:        volumeID,
	}
	response, err := client.RemoveLV(ctx, &req)
	if err != nil {
		log.Errorf("Remove Lvm with error: %v", err.Error())
		return err
	}
	log.Infof("Remove Lvm with result: %v", response.GetCommandOutput())
	return err
}

func (c *workerConnection) DeleteNameSpace(ctx context.Context, namespace string) error {
	client := lib.NewLVMClient(c.conn)
	req := lib.RemoveNamespaceRequest{
		NameSpace: namespace,
	}
	response, err := client.RemoveNamespace(ctx, &req)
	if err != nil {
		log.Errorf("Remove Lvm with error: %v", err.Error())
		return err
	}
	log.Infof("Remove Lvm with result: %v", response.GetCommandOutput())
	return err
}

func (c *workerConnection) CleanPath(ctx context.Context, path string) error {
	client := lib.NewLVMClient(c.conn)
	req := lib.CleanPathRequest{
		Path: path,
	}
	response, err := client.CleanPath(ctx, &req)
	if err != nil {
		log.Errorf("CleanPath with error: %v", err.Error())
		return err
	}
	log.Infof("CleanPath with result: %v", response.GetCommandOutput())
	return err
}

func (c *workerConnection) CreateProjQuotaSubpath(ctx context.Context, pvName, size, rootPath string) (string, string, error) {
	client := lib.NewProjQuotaClient(c.conn)
	req := lib.CreateProjQuotaSubpathRequest{
		PvName:    pvName,
		QuotaSize: size,
		RootPath:  rootPath,
	}
	response, err := client.CreateProjQuotaSubpath(ctx, &req)
	if err != nil {
		log.Errorf("CreateProjQuotaSubpath with error: %v", err.Error())
		return "", "", err
	}
	log.Infof("CreateProjQuotaSubpath with result: %v", response)

	return response.CommandOutput, response.ProjQuotaSubpath, nil
}

func (c *workerConnection) SetSubpathProjQuota(ctx context.Context, quotaSubpath, blockSoftlimit, blockHardlimit, inodeSoftlimit, inodeHardlimit string) (string, error) {
	client := lib.NewProjQuotaClient(c.conn)
	req := lib.SetSubpathProjQuotaRequest{
		ProjQuotaSubpath: quotaSubpath,
		BlockSoftlimit:   blockSoftlimit,
		BlockHardlimit:   blockHardlimit,
		InodeSoftlimit:   inodeSoftlimit,
		InodeHardlimit:   inodeHardlimit,
	}
	response, err := client.SetSubpathProjQuota(ctx, &req)
	if err != nil {
		log.Errorf("SetSubpathProjQuota with error: %v", err.Error())
		return "", err
	}
	log.Infof("SetSubpathProjQuota with result: %v", response)
	return response.GetCommandOutput(), nil
}

func (c *workerConnection) RemoveProjQuotaSubpath(ctx context.Context, quotaSubpath string) (string, error) {

	client := lib.NewProjQuotaClient(c.conn)
	req := lib.RemoveProjQuotaSubpathRequest{
		QuotaSubpath: quotaSubpath,
		ProjectId:    "",
	}
	response, err := client.RemoveProjQuotaSubpath(ctx, &req)
	if err != nil {
		log.Errorf("RemoveProjQuotaSubpath with error: %v", err.Error())
		return "", err
	}
	log.Infof("RemoveProjQuotaSubpath with result: %v", response)
	return response.GetCommandOutput(), nil
}

func logGRPC(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Debugf("GRPC request: %s, %+v", method, req)
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Debugf("GRPC response: %+v, %v", reply, err)
	return err
}
