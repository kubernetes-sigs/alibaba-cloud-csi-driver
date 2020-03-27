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
	lvs, err := commands.ListLV(ctx, in.VolumeGroup)
	if err != nil {
		log.Errorf("List LVM with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to list LVs: %v", err)
	}

	pblvs := make([]*pb.LogicalVolume, len(lvs))
	for i, v := range lvs {
		pblvs[i] = v.ToProto()
	}
	log.Infof("List LVM with result: %+v", pblvs)
	return &pb.ListLVReply{Volumes: pblvs}, nil
}

// CreateLV create lvm volume
func (s Server) CreateLV(ctx context.Context, in *pb.CreateLVRequest) (*pb.CreateLVReply, error) {
	out, err := commands.CreateLV(ctx, in.VolumeGroup, in.Name, in.Size, in.Mirrors, in.Tags)
	if err != nil {
		log.Errorf("Create LVM with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to create lv: %v", err)
	}
	log.Infof("Create LVM with result: %+v", out)
	return &pb.CreateLVReply{CommandOutput: out}, nil
}

// RemoveLV remove lvm volume
func (s Server) RemoveLV(ctx context.Context, in *pb.RemoveLVRequest) (*pb.RemoveLVReply, error) {
	out, err := commands.RemoveLV(ctx, in.VolumeGroup, in.Name)
	if err != nil {
		log.Errorf("Remove LVM with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to remove lv: %v", err)
	}
	log.Infof("Remove LVM with result: %+v", out)
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
