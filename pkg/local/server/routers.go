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

// ProjQuotaServer proj quota grpc server 
type ProjQuotaServer struct {}


// NewProjQuotaServer new proj quota grpc server
func NewProjQuotaServer() ProjQuotaServer {
    return ProjQuotaServer{}
}  

// ListLV list lvm volume
func (s Server) ListLV(ctx context.Context, in *lib.ListLVRequest) (*lib.ListLVReply, error) {
	log.Infof("List LVM for vg: %s", in.VolumeGroup)
	lvs, err := ListLV(in.VolumeGroup)
	if err != nil {
		log.Errorf("List LVM with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to list LVs: %v", err)
	}

	pblvs := make([]*lib.LogicalVolume, len(lvs))
	for i, v := range lvs {
		pblvs[i] = v.ToProto()
	}
	log.Infof("List LVM Successful with result: %+v", pblvs)
	return &lib.ListLVReply{Volumes: pblvs}, nil
}

// CreateLV create lvm volume
func (s Server) CreateLV(ctx context.Context, in *lib.CreateLVRequest) (*lib.CreateLVReply, error) {
	log.Infof("Create LVM with: %+v", in)
	out, err := CreateLV(ctx, in.VolumeGroup, in.Name, in.Size, in.Mirrors, in.Tags, in.Striping)
	if err != nil {
		log.Errorf("Create LVM with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to create lv: %v", err)
	}
	log.Infof("Create LVM Successful with result: %+v", out)
	return &lib.CreateLVReply{CommandOutput: out}, nil
}

// RemoveLV remove lvm volume
func (s Server) RemoveLV(ctx context.Context, in *lib.RemoveLVRequest) (*lib.RemoveLVReply, error) {
	log.Infof("Remove LVM with: %+v", in)
	out, err := RemoveLV(ctx, in.VolumeGroup, in.Name)
	if err != nil {
		log.Errorf("Remove LVM with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to remove lv: %v", err)
	}
	log.Infof("Remove LVM Successful with result: %+v", out)
	return &lib.RemoveLVReply{CommandOutput: out}, nil
}

// CloneLV clone lvm volume
func (s Server) CloneLV(ctx context.Context, in *lib.CloneLVRequest) (*lib.CloneLVReply, error) {
	out, err := CloneLV(ctx, in.SourceName, in.DestName)
	if err != nil {
		log.Errorf("Clone LVM with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to clone lv: %v", err)
	}
	log.Infof("Clone LVM with result: %+v", out)
	return &lib.CloneLVReply{CommandOutput: out}, nil
}

// ListVG list volume group
func (s Server) ListVG(ctx context.Context, in *lib.ListVGRequest) (*lib.ListVGReply, error) {
	vgs, err := ListVG()
	if err != nil {
		log.Errorf("List VG with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to list LVs: %v", err)
	}

	pbvgs := make([]*lib.VolumeGroup, len(vgs))
	for i, v := range vgs {
		pbvgs[i] = v.ToProto()
	}
	log.Infof("List VG with result: %+v", pbvgs)
	return &lib.ListVGReply{VolumeGroups: pbvgs}, nil
}

// CreateVG create volume group
func (s Server) CreateVG(ctx context.Context, in *lib.CreateVGRequest) (*lib.CreateVGReply, error) {
	out, err := CreateVG(ctx, in.Name, in.PhysicalVolume, in.Tags)
	if err != nil {
		log.Errorf("Create VG with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to create vg: %v", err)
	}
	log.Infof("Create VG with result: %+v", out)
	return &lib.CreateVGReply{CommandOutput: out}, nil
}

// RemoveVG remove volume group
func (s Server) RemoveVG(ctx context.Context, in *lib.CreateVGRequest) (*lib.RemoveVGReply, error) {
	out, err := RemoveVG(ctx, in.Name)
	if err != nil {
		log.Errorf("Remove VG with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to remove vg: %v", err)
	}
	log.Infof("Remove VG with result: %+v", out)
	return &lib.RemoveVGReply{CommandOutput: out}, nil
}

// CleanPath remove file under path
func (s Server) CleanPath(ctx context.Context, in *lib.CleanPathRequest) (*lib.CleanPathReply, error) {
	err := CleanPath(ctx, in.Path)
	if err != nil {
		log.Errorf("CleanPath with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to remove vg: %v", err)
	}
	log.Infof("CleanPath with result Successful")
	return &lib.CleanPathReply{CommandOutput: "Successful remove path: " + in.Path}, nil
}

// AddTagLV add tag
func (s Server) AddTagLV(ctx context.Context, in *lib.AddTagLVRequest) (*lib.AddTagLVReply, error) {
	log, err := AddTagLV(ctx, in.VolumeGroup, in.Name, in.Tags)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add tags to lv: %v", err)
	}
	return &lib.AddTagLVReply{CommandOutput: log}, nil
}

// RemoveTagLV remove tag
func (s Server) RemoveTagLV(ctx context.Context, in *lib.RemoveTagLVRequest) (*lib.RemoveTagLVReply, error) {
	log, err := RemoveTagLV(ctx, in.VolumeGroup, in.Name, in.Tags)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to remove tags from lv: %v", err)
	}
	return &lib.RemoveTagLVReply{CommandOutput: log}, nil
}

// ListNameSpace list pmem namespace
func (s Server) ListNameSpace(ctx context.Context, in *lib.ListNameSpaceRequest) (*lib.ListNameSpaceReply, error) {
	log.Infof("List NameSpace for pmem: " + in.NameSpace + " == " + in.Region)
	namespaces, err := ListNameSpace()
	if err != nil {
		log.Errorf("List NameSpace with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to list NameSpace: %v", err)
	}
	for _, namespace := range namespaces {
		if namespace.Name == in.NameSpace {
			log.Infof("List NameSpace with single result: %v", namespace)
			return &lib.ListNameSpaceReply{NameSpace: []*lib.NameSpace{namespace}}, nil
		}
	}
	log.Infof("List NameSpace with response: %v", namespaces)
	return &lib.ListNameSpaceReply{NameSpace: namespaces}, nil
}

// CreateNameSpace create pmem namespace
func (s Server) CreateNameSpace(ctx context.Context, in *lib.CreateNameSpaceRequest) (*lib.CreateNameSpaceReply, error) {
	log.Infof("Create NameSpace with: %+v", in)
	out, err := CreateNameSpace(ctx, in.Region, in.Name, in.Size)
	if err != nil {
		log.Errorf("Create NameSpace with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to create NameSpace: %v", err)
	}
	log.Infof("Create NameSpace Successful with result: %+v", out)
	return &lib.CreateNameSpaceReply{CommandOutput: out}, nil
}

// RemoveNameSpace remove pmem namespace
func (s Server) RemoveNameSpace(ctx context.Context, in *lib.RemoveNameSpaceRequest) (*lib.RemoveNameSpaceReply, error) {
	log.Infof("Remove NameSpace with: %+v", in)
	out, err := RemoveNameSpace(ctx, in.NameSpace)
	if err != nil {
		log.Errorf("Remove NameSpace with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to remove NameSpace: %v", err)
	}
	log.Infof("Remove NameSpace Successful with result: %+v", out)
	return &lib.RemoveNameSpaceReply{CommandOutput: out}, nil
}

// CreateProjQuotaSubpath ...
func (s ProjQuotaServer) CreateProjQuotaSubpath(ctx context.Context, in *lib.CreateProjQuotaSubpathRequest) (*lib.CreateProjQuotaSubpathReply, error) {
	log.Infof("CreateProjQuotaSubpath with %+v", in)
	projQuotaSubpath, out, projectID, err := CreateProjQuotaSubpath(ctx, in.PvName, in.QuotaSize)
	if err != nil {
		log.Errorf("CreateProjQuotaSubpath with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to create projQuota subpath: %v", err)
	}
	log.Infof("CreateProjQuotaSubpath successful with result %+v", out)
	return &lib.CreateProjQuotaSubpathReply{ProjQuotaSubpath: projQuotaSubpath, CommandOutput: out, ProjectId: projectID}, nil
}

// SetSubpathProjQuota ...
func (s ProjQuotaServer) SetSubpathProjQuota(ctx context.Context, in *lib.SetSubpathProjQuotaRequest) (*lib.SetSubpathProjQuotaReply, error) {
	log.Infof("SetSubpathProjQuota with %+v", in)
	out, err := SetSubpathProjQuota(ctx, in.ProjQuotaSubpath, in.ProjectId, in.BlockHardlimit, in.BlockSoftlimit)
	if err != nil {
		log.Errorf("SetSubpathProjQuota with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to set projQuota to subpath: %v", err)
	}
	log.Infof("SetSubpathProjQuota successful with result %+v", out)
	return &lib.SetSubpathProjQuotaReply{CommandOutput: out}, nil
}

// RemoveProjQuotaSubpath ...
func (s ProjQuotaServer) RemoveProjQuotaSubpath(ctx context.Context, in *lib.RemoveProjQuotaSubpathRequest) (*lib.RemoveProjQuotaSubpathReply, error) {
	log.Infof("RemoveProjQuotaSubpath with %+v", in)
	out, err := RemoveProjQuotaSubpath(ctx, in.QuotaSubpath)
	if err != nil {
		log.Errorf("RemoveProjQuotaSubpath with error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to remove projQuota subpath: %v", err)
	}
	log.Infof("RemoveProjQuotaSubpath successful with result %+v", out)
	return &lib.RemoveProjQuotaSubpathReply{CommandOutput: out}, nil
}