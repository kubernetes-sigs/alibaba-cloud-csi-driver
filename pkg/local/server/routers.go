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
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/lib/commands"
	pb "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/lib/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server lvm grpc server
type Server struct{}

// NewServer new server
func NewServer() Server {
	return Server{}
}

// ListLV list lvm volume
func (s Server) ListLV(ctx context.Context, in *pb.ListLVRequest) (*pb.ListLVReply, error) {
	log.Infof("List LVM for vg: %s", in.VolumeGroup)
	lvs, err := commands.ListLV(in.VolumeGroup)
	if err != nil {
		log.Errorf("List LVM with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to list LVs: %v", err)
	}

	pblvs := make([]*pb.LogicalVolume, len(lvs))
	for i, v := range lvs {
		pblvs[i] = v.ToProto()
	}
	log.Infof("List LVM Successful with result: %+v", pblvs)
	return &pb.ListLVReply{Volumes: pblvs}, nil
}

// ListLV list lvm volume
func (s Server) ListNameSpace(ctx context.Context, in *pb.ListNameSpaceRequest) (*pb.ListNameSpaceReply, error) {
	log.Infof("List NameSpace for namespace: " +  in.NameSpace +" == " + in.Region)
	namespaces, err := commands.ListNameSpace()
	if err != nil {
		log.Errorf("List NameSpace with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to list NameSpace: %v", err)
	}
	for _, namespace := range namespaces {
		if namespace.Name == in.NameSpace {
			log.Infof("List NameSpace with sigle response: %v", namespace)
			return &pb.ListNameSpaceReply{NameSpace: []*pb.NameSpace{namespace}}, nil
		}
	}
	log.Infof("List NameSpace with response: %v", namespaces)
	return &pb.ListNameSpaceReply{NameSpace: namespaces}, nil
}

// CreateLV create lvm volume
func (s Server) CreateNameSpace(ctx context.Context, in *pb.CreateNameSpaceRequest) (*pb.CreateNameSpaceReply, error) {
	log.Infof("Create LVM with: %+v", in)
	out, err := commands.CreateNameSpace(ctx, in.Region, in.Name, in.Size)
	if err != nil {
		log.Errorf("Create LVM with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to create lv: %v", err)
	}
	log.Infof("Create LVM Successful with result: %+v", out)
	return &pb.CreateNameSpaceReply{CommandOutput: out}, nil
}

// RemoveLV remove lvm volume
func (s Server) RemoveNameSpace(ctx context.Context, in *pb.RemoveNameSpaceRequest) (*pb.RemoveNameSpaceReply, error) {
	log.Infof("Remove NameSpace with: %+v", in)
	out, err := commands.RemoveNameSpace(ctx, in.NameSpace)
	if err != nil {
		log.Errorf("Remove NameSpace with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to remove lv: %v", err)
	}
	log.Infof("Remove NameSpace Successful with result: %+v", out)
	return &pb.RemoveNameSpaceReply{CommandOutput: out}, nil
}

// CreateLV create lvm volume
func (s Server) CreateLV(ctx context.Context, in *pb.CreateLVRequest) (*pb.CreateLVReply, error) {
	log.Infof("Create LVM with: %+v", in)
	out, err := commands.CreateLV(ctx, in.VolumeGroup, in.Name, in.Size, in.Mirrors, in.Tags, in.Striping)
	if err != nil {
		log.Errorf("Create LVM with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to create lv: %v", err)
	}
	log.Infof("Create LVM Successful with result: %+v", out)
	return &pb.CreateLVReply{CommandOutput: out}, nil
}

// RemoveLV remove lvm volume
func (s Server) RemoveLV(ctx context.Context, in *pb.RemoveLVRequest) (*pb.RemoveLVReply, error) {
	log.Infof("Remove LVM with: %+v", in)
	out, err := commands.RemoveLV(ctx, in.VolumeGroup, in.Name)
	if err != nil {
		log.Errorf("Remove LVM with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to remove lv: %v", err)
	}
	log.Infof("Remove LVM Successful with result: %+v", out)
	return &pb.RemoveLVReply{CommandOutput: out}, nil
}

// CloneLV clone lvm volume
func (s Server) CloneLV(ctx context.Context, in *pb.CloneLVRequest) (*pb.CloneLVReply, error) {
	out, err := commands.CloneLV(ctx, in.SourceName, in.DestName)
	if err != nil {
		log.Errorf("Clone LVM with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to clone lv: %v", err)
	}
	log.Infof("Clone LVM with result: %+v", out)
	return &pb.CloneLVReply{CommandOutput: out}, nil
}

// ListVG list volume group
func (s Server) ListVG(ctx context.Context, in *pb.ListVGRequest) (*pb.ListVGReply, error) {
	vgs, err := commands.ListVG()
	if err != nil {
		log.Errorf("List VG with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to list LVs: %v", err)
	}

	pbvgs := make([]*pb.VolumeGroup, len(vgs))
	for i, v := range vgs {
		pbvgs[i] = v.ToProto()
	}
	log.Infof("List VG with result: %+v", pbvgs)
	return &pb.ListVGReply{VolumeGroups: pbvgs}, nil
}

// CreateVG create volume group
func (s Server) CreateVG(ctx context.Context, in *pb.CreateVGRequest) (*pb.CreateVGReply, error) {
	out, err := commands.CreateVG(ctx, in.Name, in.PhysicalVolume, in.Tags)
	if err != nil {
		log.Errorf("Create VG with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to create vg: %v", err)
	}
	log.Infof("Create VG with result: %+v", out)
	return &pb.CreateVGReply{CommandOutput: out}, nil
}

// RemoveVG remove volume group
func (s Server) RemoveVG(ctx context.Context, in *pb.CreateVGRequest) (*pb.RemoveVGReply, error) {
	out, err := commands.RemoveVG(ctx, in.Name)
	if err != nil {
		log.Errorf("Remove VG with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to remove vg: %v", err)
	}
	log.Infof("Remove VG with result: %+v", out)
	return &pb.RemoveVGReply{CommandOutput: out}, nil
}

// CleanPath remove file under path
func (s Server) CleanPath(ctx context.Context, in *pb.CleanPathRequest) (*pb.CleanPathReply, error) {
	err := commands.CleanPath(ctx, in.Path)
	if err != nil {
		log.Errorf("CleanPath with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to remove vg: %v", err)
	}
	log.Infof("CleanPath with result Successful")
	return &pb.CleanPathReply{CommandOutput: "Successful remove path: " + in.Path}, nil
}

// AddTagLV add tag
func (s Server) AddTagLV(ctx context.Context, in *pb.AddTagLVRequest) (*pb.AddTagLVReply, error) {
	log, err := commands.AddTagLV(ctx, in.VolumeGroup, in.Name, in.Tags)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add tags to lv: %v", err)
	}
	return &pb.AddTagLVReply{CommandOutput: log}, nil
}

// RemoveTagLV remove tag
func (s Server) RemoveTagLV(ctx context.Context, in *pb.RemoveTagLVRequest) (*pb.RemoveTagLVReply, error) {
	log, err := commands.RemoveTagLV(ctx, in.VolumeGroup, in.Name, in.Tags)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to remove tags from lv: %v", err)
	}
	return &pb.RemoveTagLVReply{CommandOutput: log}, nil
}
