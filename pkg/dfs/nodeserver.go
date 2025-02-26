/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http:// www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dfs

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
)

const (
	ConditionVscAttached = "pov.alibabacloud.com/vcs-attached"
)

type nodeServer struct {
	common.GenericNodeServer
	mounter    mount.Interface
	kubeClient kubernetes.Interface
}

func newNodeServer(nodeId string, kubeClient kubernetes.Interface) *nodeServer {
	return &nodeServer{
		mounter:    mount.New(""),
		kubeClient: kubeClient,
		GenericNodeServer: common.GenericNodeServer{
			NodeID: nodeId,
		},
	}
}

func (n *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	klog.Info("DFS NodePublishVolume")
	if err := n.setPodVscAttachedCondition(ctx, req); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	notMnt, err := n.mounter.IsLikelyNotMountPoint(req.TargetPath)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(req.TargetPath, os.ModePerm); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			notMnt = true
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	if !notMnt {
		klog.Infof("NodePublishVolume: %s already mounted", req.TargetPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	err = n.mounter.Mount("tmpfs", req.TargetPath, "tmpfs", []string{"ro"})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Mount tmpfs: %v", err)
	}
	return &csi.NodePublishVolumeResponse{}, nil
}

func (n *nodeServer) NodeUnpublishVolume(_ context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	klog.Info("DFS NodeUnpublishVolume")
	err := mount.CleanupMountPoint(req.TargetPath, n.mounter, false)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cleanup mount point: %v", err)
	}
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (n *nodeServer) setPodVscAttachedCondition(ctx context.Context, req *csi.NodePublishVolumeRequest) error {
	pod, err := utils.GetPodFromContextOrK8s(ctx, n.kubeClient, req)
	if err != nil {
		return fmt.Errorf("Get pod: %w", err)
	}

	for _, c := range pod.Status.Conditions {
		if c.Type == ConditionVscAttached && c.Status == corev1.ConditionTrue {
			return nil
		}
	}
	patch := corev1.Pod{
		Status: corev1.PodStatus{
			Conditions: []corev1.PodCondition{
				{
					Type:               ConditionVscAttached,
					Status:             corev1.ConditionTrue,
					LastTransitionTime: metav1.Now(),
				},
			},
		},
	}
	data, err := json.Marshal(patch)
	if err != nil {
		return err
	}
	_, err = n.kubeClient.CoreV1().Pods(pod.Namespace).Patch(ctx, pod.Name, types.StrategicMergePatchType, data, metav1.PatchOptions{}, "status")
	if err != nil {
		return fmt.Errorf("Set vsc-attached condition: %w", err)
	}
	return nil
}
