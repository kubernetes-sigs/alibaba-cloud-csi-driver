package client

import (
	"context"
	"errors"
	"fmt"
	"net"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/lib"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/server"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type mockProjQuotaServer struct {
	server.ProjQuotaServer
}

func (*mockProjQuotaServer) CreateProjQuotaSubpath(ctx context.Context, in *lib.CreateProjQuotaSubpathRequest) (*lib.CreateProjQuotaSubpathReply, error) {
	pvName := in.PvName
	rootPath := in.RootPath
	projectID := server.ConvertString2int(pvName)

	if pvName == "err" {
		return &lib.CreateProjQuotaSubpathReply{}, errors.New("fake error")
	}

	// select quotapath.namespace1.0 as default namespace for mock
	const namespace = "namespace1.0"
	var fullPath string
	if rootPath == "" {
		fullPath = fmt.Sprintf(server.ProjQuotaPrefix, namespace, pvName)
	} else {
		fullPath = rootPath + "/quotapath." + namespace + "/" + pvName
	}

	reply := &lib.CreateProjQuotaSubpathReply{
		ProjQuotaSubpath: fullPath,
		CommandOutput:    "",
		ProjectId:        projectID,
	}
	return reply, nil
}

func (*mockProjQuotaServer) SetSubpathProjQuota(ctx context.Context, in *lib.SetSubpathProjQuotaRequest) (*lib.SetSubpathProjQuotaReply, error) {

	quotaSubpath := in.ProjQuotaSubpath

	if quotaSubpath == "err" {
		return &lib.SetSubpathProjQuotaReply{}, errors.New("fake error")
	}
	reply := &lib.SetSubpathProjQuotaReply{
		CommandOutput: "",
	}
	return reply, nil
}

func (*mockProjQuotaServer) RemoveProjQuotaSubpath(ctx context.Context, in *lib.RemoveProjQuotaSubpathRequest) (*lib.RemoveProjQuotaSubpathReply, error) {
	quotaPath := in.QuotaSubpath
	if quotaPath == "err" {
		return &lib.RemoveProjQuotaSubpathReply{}, errors.New("fake error")
	}
	reply := &lib.RemoveProjQuotaSubpathReply{
		CommandOutput: "",
	}
	return reply, nil
}

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	lib.RegisterProjQuotaServer(server, &mockProjQuotaServer{})
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestCreateProjQuotaSubpath(t *testing.T) {
	tests := []struct {
		pvName   string
		rootPath string
		fullPath string
	}{
		{
			"pvName",
			"/data",
			"/data/quotapath.namespace1.0/pvName",
		},
		{
			"err",
			"",
			"",
		},
		{
			"testpv",
			"",
			"/mnt/quotapath.namespace1.0/testpv",
		},
	}
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for _, tt := range tests {
		t.Run(tt.pvName, func(t *testing.T) {
			workConnection := &workerConnection{conn: conn}
			_, fullPath, err := workConnection.CreateProjQuotaSubpath(ctx, tt.pvName, "", tt.rootPath)
			if tt.pvName == "err" {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, tt.fullPath, fullPath)
			}
		})
	}
}
