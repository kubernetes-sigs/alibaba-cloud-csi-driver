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
	"net"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

// LVMConnection lvm connection interface
type LVMConnection interface {
	GetLvm(ctx context.Context, volGroup string, volumeID string) (string, error)
	CreateLvm(ctx context.Context, opt *LVMOptions) (string, error)
	DeleteLvm(ctx context.Context, volGroup string, volumeID string) error
	CleanPath(ctx context.Context, path string) error
	Close() error
	GetNameSpace(ctx context.Context, regionName string, volumeID string) (string, error)
	CreateNameSpace(ctx context.Context, opt *NameSpaceOptions) (*lib.PmemNameSpace, error)
	DeleteNameSpace(ctx context.Context, volumeID string) error
}

// LVMOptions lvm options
type LVMOptions struct {
	VolumeGroup string
	Name        string
	Size        uint64
	Tags        []string
	Striping    bool
}

// NameSpaceOptions lvm options
type NameSpaceOptions struct {
	Region string
	Name   string
	Size   uint64
}

//
type lvmdConnection struct {
	conn *grpc.ClientConn
}

var (
	_ LVMConnection = &lvmdConnection{}
)

// NewGrpcConnection lvm connection
func NewGrpcConnection(address string, timeout time.Duration) (LVMConnection, error) {
	conn, err := connect(address, timeout)
	if err != nil {
		return nil, err
	}
	return &lvmdConnection{
		conn: conn,
	}, nil
}

func (c *lvmdConnection) Close() error {
	return c.conn.Close()
}

func connect(address string, timeout time.Duration) (*grpc.ClientConn, error) {
	log.Infof("New LVM Connecting to %s", address)
	dialOptions := []grpc.DialOption{
		grpc.WithInsecure(),
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

func (c *lvmdConnection) CreateLvm(ctx context.Context, opt *LVMOptions) (string, error) {
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

func (c *lvmdConnection) CreateNameSpace(ctx context.Context, opt *NameSpaceOptions) (*lib.PmemNameSpace, error) {
	client := lib.NewLVMClient(c.conn)
	req := lib.CreateNameSpaceRequest{
		Name:   opt.Name,
		Size:   opt.Size,
		Region: opt.Region,
	}

	rsp, err := client.CreateNameSpace(ctx, &req)
	if err != nil {
		log.Errorf("Create Lvm with error: %s", err.Error())
		return nil, err
	}
	log.Infof("Create Lvm with result: %+v", rsp.CommandOutput)

	pns := &lib.PmemNameSpace{}
	if err := json.Unmarshal([]byte(rsp.CommandOutput), pns); err != nil {
		return nil, err
	}
	return pns, nil
}

func (c *lvmdConnection) GetNameSpace(ctx context.Context, regionName string, volumeID string) (string, error) {
	client := lib.NewLVMClient(c.conn)
	req := lib.ListNameSpaceRequest{
		NameSpace: volumeID,
		Region:    regionName,
	}

	log.Infof("xxxxxxxxxx: %v", req)
	rsp, err := client.ListNameSpace(ctx, &req)
	if err != nil {
		return "", err
	}
	log.Infof("zzzzzzzzzzzzzz %v", rsp)
	for _, namespace := range rsp.NameSpace {
		if namespace.Name == volumeID {
			return namespace.Name, nil
		}
	}
	return "", nil
}

func (c *lvmdConnection) GetLvm(ctx context.Context, volGroup string, volumeID string) (string, error) {
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

func (c *lvmdConnection) DeleteLvm(ctx context.Context, volGroup, volumeID string) error {
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

func (c *lvmdConnection) DeleteNameSpace(ctx context.Context, namespace string) error {
	client := lib.NewLVMClient(c.conn)
	req := lib.RemoveNameSpaceRequest{
		NameSpace: namespace,
	}
	response, err := client.RemoveNameSpace(ctx, &req)
	if err != nil {
		log.Errorf("Remove Lvm with error: %v", err.Error())
		return err
	}
	log.Infof("Remove Lvm with result: %v", response.GetCommandOutput())
	return err
}

func (c *lvmdConnection) CleanPath(ctx context.Context, path string) error {
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

func logGRPC(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Debugf("GRPC request: %s, %+v", method, req)
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Debugf("GRPC response: %+v, %v", reply, err)
	return err
}
