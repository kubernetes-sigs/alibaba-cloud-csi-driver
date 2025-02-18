package ens

import (
	"time"

	http "github.com/alibabacloud-go/darabonba-openapi/client"
	ensCli "github.com/alibabacloud-go/ens-20171110/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"k8s.io/klog/v2"
)

type ENSClient struct {
	c *ensCli.Client
}

func newENSClient() ENSClient {
	accessControl := utils.GetAccessControl()

	config := new(http.Config)

	config.AccessKeyId = &accessControl.AccessKeyID
	config.AccessKeySecret = &accessControl.AccessKeySecret
	config.SecurityToken = &accessControl.StsToken

	config.Endpoint = tea.String("ens.aliyuncs.com")
	//klog.Infof("newENSClient: fortest, ak: %s, sk: %s, sts: %s", *config.AccessKeyId, *config.AccessKeySecret, *config.SecurityToken)
	ec, err := ensCli.NewClient(config)
	if err != nil {
		klog.Fatalf("newENSClient: failed to create ens client: %+v", err)
	}
	return ENSClient{c: ec}
}

func (ec *ENSClient) DescribeInstance(instanceId string) ([]*ensCli.DescribeInstancesResponseBodyInstancesInstance, error) {
	dir := &ensCli.DescribeInstancesRequest{
		InstanceId: tea.String(instanceId),
	}
	resp, err := ec.c.DescribeInstances(dir)
	if err != nil {
		klog.Errorf("DescribeInstance: describe instance failed err: %+v", err)
		if ec.renewToken() != nil {
			klog.Errorf("DescribeInstance: describe instance failed err: %+v", err)
			return []*ensCli.DescribeInstancesResponseBodyInstancesInstance{}, err
		}
		resp, err = ec.c.DescribeInstances(dir)
	}
	if err != nil {
		klog.Errorf("DescribeInstance: describe instance failed err: %+v", err)
		return []*ensCli.DescribeInstancesResponseBodyInstancesInstance{}, err
	}
	return resp.Body.Instances.Instance, nil
}

func (ec *ENSClient) CreateVolume(regionID, diskType, size string) (string, error) {
	cdr := &ensCli.CreateDiskRequest{
		InstanceChargeType: tea.String("PostPaid"),
		EnsRegionId:        tea.String(regionID),
		Category:           tea.String(diskType),
		Size:               tea.String(size),
	}

	resp, err := ec.c.CreateDisk(cdr)
	if err != nil {
		klog.Errorf("CreateVolume: create volume failed err: %+v", err)
		if ec.renewToken() != nil {
			klog.Errorf("CreateVolume: create volume failed err: %+v", err)
			return "", err
		}
		resp, err = ec.c.CreateDisk(cdr)
	}
	if err != nil {
		klog.Errorf("CreateVolume: create volume failed err: %+v, regionID: %s, category: %s, size: %s", err, regionID, diskType, size)
		return "", err
	}
	return *resp.Body.InstanceIds[0], nil
}

func (ec *ENSClient) DeleteVolume(diskID string) {

}

func (ec *ENSClient) AttachVolume(diskID, instanceID string) error {
	adr := &ensCli.AttachDiskRequest{
		DiskId:     tea.String(diskID),
		InstanceId: tea.String(instanceID),
	}

	resp, err := ec.c.AttachDisk(adr)
	klog.Infof("AttachVolume: attach disk resp: %+v", resp)
	if err != nil {
		klog.Errorf("AttachVolume: attach volume failed err: %+v", err)
		if ec.renewToken() != nil {
			klog.Errorf("AttachVolume: attach volume failed err: %+v", err)
			return err
		}
		_, err = ec.c.AttachDisk(adr)
	}
	if err != nil {
		klog.Errorf("AttachVolume: attach volume failed err: %+v", err)
		return err
	}
	return nil

}

func (ec *ENSClient) DescribeVolume(diskID string) (*ensCli.DescribeDisksResponseBodyDisksDisks, error) {
	ddr := &ensCli.DescribeDisksRequest{
		DiskId:      tea.String(diskID),
		EnsRegionId: tea.String(GlobalConfigVar.RegionID),
	}
	resp, err := ec.c.DescribeDisks(ddr)

	if err != nil {
		klog.Errorf("DescribeVolume: describe volume failed err: %+v", err)
		if ec.renewToken() != nil {
			klog.Errorf("DescribeVolume: describe volume failed err: %+v", err)
			return nil, err
		}
		resp, err = ec.c.DescribeDisks(ddr)
	}
	if err != nil {
		klog.Errorf("DescribeVolume: describe volume failed err: %+v", err)
		return nil, err
	}
	return resp.Body.Disks.Disks[0], nil
}

func (ec *ENSClient) DetachVolume(diskID, instanceID string) error {
	ddr := &ensCli.DetachDiskRequest{
		DiskId:     tea.String(diskID),
		InstanceId: tea.String(instanceID),
	}
	resp, err := ec.c.DetachDisk(ddr)
	klog.Infof("DetachVolume: detach disk resp: %+v", resp)

	if err != nil {
		klog.Errorf("DetachVolume: detach volume failed err: %+v", err)
		if ec.renewToken() != nil {
			klog.Errorf("DetachVolume: detach volume failed err: %+v", err)
			return err
		}
		_, err = ec.c.DetachDisk(ddr)
	}
	if err != nil {
		klog.Errorf("DescribeVolume: describe volume failed err: %+v", err)
		return err
	}
	return nil
}

func (ec *ENSClient) renewToken() error {
	klog.Infof("renewToken: currentTime: %s", time.Now().String())

	accessControl := utils.GetAccessControl()

	config := new(http.Config)
	config.AccessKeyId = &accessControl.AccessKeyID
	config.AccessKeySecret = &accessControl.AccessKeySecret
	config.SecurityToken = &accessControl.StsToken
	config.Endpoint = tea.String("ens.aliyuncs.com")
	klog.Infof("newENSClient: fortest, ak: %s, sk: %s, sts: %s", *config.AccessKeyId, *config.AccessKeySecret, *config.SecurityToken)
	ensClient, err := ensCli.NewClient(config)
	if err != nil {
		klog.Errorf("renewToken: failed to renew token: %+v", err)
		return err
	}
	ec.c = ensClient
	return nil
}
