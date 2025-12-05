package oss

import (
	"context"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	ossfpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager/oss"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/kata/directvolume"
	"github.com/stretchr/testify/assert"
)

func Test_nodeServer_publishDirectVolume(t *testing.T) {
	ns := &nodeServer{}
	req := &csi.NodePublishVolumeRequest{
		VolumeId:          "test-id",
		PublishContext:    nil,
		StagingTargetPath: "",
		TargetPath:        "/foo/bar/path/for/publish",
		VolumeCapability:  nil,
		Readonly:          false,
		Secrets:           nil,
		VolumeContext: map[string]string{
			"encrypted": "local_encrypt",
			"kmsKeyId":  "foo-key-id",
		},
	}
	opts := &ossfpm.Options{
		URL:       "https://oss-cn-hangzhou.aliyuncs.com",
		Bucket:    "test-bucket",
		Path:      "bucket/path",
		OtherOpts: "-o test -w abc",
	}
	resp, err := ns.publishDirectVolume(context.TODO(), req, opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	defer directvolume.Remove(req.TargetPath)
	ret := isDirectVolumePath(req.TargetPath)
	assert.True(t, ret)

	info, err := directvolume.VolumeMountInfo(req.TargetPath)
	assert.NoError(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, "alibaba-cloud-oss", info.VolumeType)
	assert.Equal(t, "secure_mount", info.FsType)
	assert.Equal(t, []string{"-o", "test", "-w", "abc"}, info.Options)
	assert.Equal(t, opts.URL, info.Metadata["url"])
	assert.Equal(t, opts.Bucket, info.Metadata["bucket"])
	assert.Equal(t, opts.Path, info.Metadata["path"])
	assert.Equal(t, req.TargetPath, info.Metadata["targetPath"])
	assert.Equal(t, req.VolumeContext["kmsKeyId"], info.Metadata["kmsKeyId"])
	assert.Equal(t, req.VolumeContext["encrypted"], info.Metadata["encrypted"])
	assert.Equal(t, opts.OtherOpts, info.Metadata["otherOpts"])

	// publish twice
	resp, err = ns.publishDirectVolume(context.TODO(), req, opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func Test_nodeServer_publishDirectVolume_overwrite_annotations(t *testing.T) {
	ns := &nodeServer{}
	req := &csi.NodePublishVolumeRequest{
		VolumeId:          "test-id",
		PublishContext:    nil,
		StagingTargetPath: "",
		TargetPath:        "/foo/bar/path/for/publish",
		VolumeCapability:  nil,
		Readonly:          false,
		Secrets:           nil,
		VolumeContext: map[string]string{
			"annotations": `{"kata_fs_type": "type_v2", "kata_volume_type": "volume_type_v", "kata_device": "kata_device_vv" }`,
		},
	}
	opts := &ossfpm.Options{}
	resp, err := ns.publishDirectVolume(context.TODO(), req, opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	defer directvolume.Remove(req.TargetPath)
	ret := isDirectVolumePath(req.TargetPath)
	assert.True(t, ret)

	info, err := directvolume.VolumeMountInfo(req.TargetPath)
	assert.NoError(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, "volume_type_v", info.VolumeType)
	assert.Equal(t, "type_v2", info.FsType)
	assert.Equal(t, "kata_device_vv", info.Device)

	// publish twice
	resp, err = ns.publishDirectVolume(context.TODO(), req, opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func Test_nodeServer_unPublishDirectVolume(t *testing.T) {
	ns := &nodeServer{}
	req := &csi.NodeUnpublishVolumeRequest{
		VolumeId:   "test-id",
		TargetPath: "/foo/bar/path/for/unpublish",
	}

	err := directvolume.AddMountInfo(req.TargetPath, directvolume.MountInfo{})
	assert.NoError(t, err)
	info, err := directvolume.VolumeMountInfo(req.TargetPath)
	assert.NoError(t, err)
	assert.NotNil(t, info)

	resp, err := ns.unPublishDirectVolume(context.TODO(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	info, err = directvolume.VolumeMountInfo(req.TargetPath)
	assert.Error(t, err)
	assert.Nil(t, info)

	// unpublish twice
	resp, err = ns.unPublishDirectVolume(context.TODO(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
