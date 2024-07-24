package nas

import (
	"context"
	"errors"
	aliErrors "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/nas"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/mock/gomock"
	"github.com/jarcoal/httpmock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/interfaces"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"sync"
	"testing"
)

const (
	fileSystemControllerTestRegion    = "cn-hangzhou"
	fileSystemControllerTestMasterUrl = "https://api-server-host:6443"
)

func TestNewFilesystemController(t *testing.T) {
	prepareFilesystemControllerTestFakeK8sContext()
	controller, err := newFilesystemController(&internal.ControllerConfig{})
	assert.NoError(t, err)
	assert.NotNil(t, controller)
}

func prepareFilesystemControllerTestFakeK8sContext() {
	options.MasterURL = fileSystemControllerTestMasterUrl
}

func TestVolumeAs(t *testing.T) {
	controller := newTestFileSystemController(t)
	actual := controller.VolumeAs()
	assert.Equal(t, "filesystem", actual)
}

func newTestFileSystemController(t *testing.T) *filesystemController {
	ctrl := gomock.NewController(t)
	prepareFilesystemControllerTestFakeK8sContext()
	factory := interfaces.NewMockNasClientFactory(ctrl)
	return &filesystemController{
		config: &internal.ControllerConfig{
			Region:           fileSystemControllerTestRegion,
			NasClientFactory: factory,
		},
	}
}

func TestCreateVolume_InitNasClientError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerZoneIdMetadataResponder()
	prepareFilesystemControllerTestFakeK8sContext()

	controller := &filesystemController{
		config: &internal.ControllerConfig{
			NasClientFactory: cloud.NewNasClientFactory(),
		},
	}
	req := &csi.CreateVolumeRequest{
		Parameters: map[string]string{
			VpcID:     "vpc-id",
			VSwitchID: "vswitch-id",
		},
	}
	resp, err := controller.CreateVolume(context.Background(), req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nas client")
	assert.Nil(t, resp)
}

func registerZoneIdMetadataResponder() {
	url := "=~^" + utils.MetadataURL + `/zone-id.*\z`
	responder := httpmock.NewStringResponder(200, "cn-hangzhou-a")
	httpmock.RegisterResponder("GET", url, responder)
}

func TestCreateVolume(t *testing.T) {
	type args struct {
		resetStatus bool
		req         *csi.CreateVolumeRequest
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerZoneIdMetadataResponder()
	prepareFilesystemControllerTestFakeK8sContext()
	prepareNasClientCredentials(t)

	controller := newTestFileSystemControllerWithExpects(t, func(v1Interface *interfaces.MockNasV1Interface) {
		v1NasInterfaceExpectsCreateFileSystem(v1Interface, "request-id", "file-system-id", nil)
		v1NasInterfaceExpectsTagResources(v1Interface, nil)
		v1NasInterfaceExpectsCreateMountTarget(v1Interface, nil)
		v1NasInterfaceExpectsDescribeFileSystems(v1Interface, []string{"test.mount.target.domain"}, nil)
		v1NasInterfaceExpectsDescribeMountTargets(v1Interface, []string{"Active"}, nil)
	})

	tests := []struct {
		name     string
		args     args
		expected *csi.Volume
		wantErr  bool
	}{
		{
			name: "Nas 4.0",
			args: args{
				resetStatus: true,
				req: &csi.CreateVolumeRequest{
					VolumeCapabilities: []*csi.VolumeCapability{
						{
							AccessType: &csi.VolumeCapability_Mount{
								Mount: &csi.VolumeCapability_MountVolume{
									MountFlags: []string{"vers=4.0"},
								},
							},
						},
					},
					Parameters: map[string]string{
						VpcID:     "vpc-id",
						VSwitchID: "vswitch-id",
					},
				},
			},
			expected: &csi.Volume{
				CapacityBytes: 0,
				VolumeId:      "",
				VolumeContext: map[string]string{
					"deleteVolume": "false",
					"fileSystemId": "file-system-id",
					"path":         "/",
					"server":       "",
				},
				ContentSource:      nil,
				AccessibleTopology: nil,
			},
			wantErr: false,
		},
		{
			name: "Nas 4.1",
			args: args{
				resetStatus: false,
				req: &csi.CreateVolumeRequest{
					VolumeCapabilities: []*csi.VolumeCapability{
						{
							AccessType: &csi.VolumeCapability_Mount{
								Mount: &csi.VolumeCapability_MountVolume{
									MountFlags: []string{"vers=4.1"},
								},
							},
						},
					},
					Parameters: map[string]string{
						VpcID:     "vpc-id",
						VSwitchID: "vswitch-id",
					},
				},
			},
			expected: &csi.Volume{
				CapacityBytes: 0,
				VolumeId:      "",
				VolumeContext: map[string]string{
					"deleteVolume": "false",
					"fileSystemId": "file-system-id",
					"path":         "/",
					"server":       "",
				},
				ContentSource:      nil,
				AccessibleTopology: nil,
			},
			wantErr: false,
		},
		{
			name: "Repeat creating volume",
			args: args{
				resetStatus: true,
				req:         &csi.CreateVolumeRequest{},
			},
			expected: &csi.Volume{
				CapacityBytes: 0,
				VolumeId:      "",
				VolumeContext: map[string]string{
					"deleteVolume": "false",
					"fileSystemId": "file-system-id",
					"path":         "/",
					"server":       "",
				},
				ContentSource:      nil,
				AccessibleTopology: nil,
			},
			wantErr: false,
		},
		{
			name: "Extreme file system type",
			args: args{
				resetStatus: true,
				req:         createExtremeFileSystemRequest(),
			},
			expected: &csi.Volume{
				CapacityBytes: 100 * 1024 * 1024 * 1024,
				VolumeId:      "",
				VolumeContext: map[string]string{
					"deleteVolume": "false",
					"fileSystemId": "file-system-id",
					"path":         "/share",
					"server":       "test.mount.target.domain",
					"vers":         "3",
				},
				ContentSource:      nil,
				AccessibleTopology: nil,
			},
			wantErr: false,
		},
		{
			name: "Invalid access type",
			args: args{
				resetStatus: true,
				req: &csi.CreateVolumeRequest{
					VolumeCapabilities: []*csi.VolumeCapability{
						{
							AccessType: &csi.VolumeCapability_Block{},
						},
					},
				},
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Invalid parameters",
			args: args{
				resetStatus: true,
				req: &csi.CreateVolumeRequest{
					Parameters: map[string]string{},
				},
			},
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := controller.CreateVolume(context.Background(), tt.args.req)
			if tt.expected == nil {
				assert.Nil(t, actual)
			} else {
				assert.NotNil(t, actual)
				assert.NotNil(t, actual.Volume)
				assert.Equal(t, *tt.expected, *actual.Volume)
			}
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			if tt.args.resetStatus {
				resetFileSystemControllerStatus(controller)
			}
		})
	}
}

func createExtremeFileSystemRequest() *csi.CreateVolumeRequest {
	return &csi.CreateVolumeRequest{
		CapacityRange: &csi.CapacityRange{
			RequiredBytes: 100 * 1024 * 1024 * 1024,
		},
		Parameters: map[string]string{
			FileSystemType: "extreme",
			VpcID:          "vpc-id",
			VSwitchID:      "vswitch-id",
		},
	}
}

func prepareNasClientCredentials(t *testing.T) {
	t.Setenv("ACCESS_KEY_ID", "ID")
	t.Setenv("ACCESS_KEY_SECRET", "SECRET")
}

func v1NasInterfaceExpectsCreateFileSystem(v1Interface *interfaces.MockNasV1Interface, requestId, fileSystemId string, err error) {
	v1Interface.EXPECT().CreateFileSystem(gomock.Any()).Return(&nas.CreateFileSystemResponse{
		RequestId:    requestId,
		FileSystemId: fileSystemId,
	}, err).AnyTimes()
}

func v1NasInterfaceExpectsTagResources(v1Interface *interfaces.MockNasV1Interface, err error) {
	v1Interface.EXPECT().TagResources(gomock.Any()).Return(&nas.TagResourcesResponse{}, err).AnyTimes()
}

func v1NasInterfaceExpectsCreateMountTarget(v1Interface *interfaces.MockNasV1Interface, err error) {
	v1Interface.EXPECT().CreateMountTarget(gomock.Any()).Return(&nas.CreateMountTargetResponse{}, err).AnyTimes()
}

func v1NasInterfaceExpectsDescribeFileSystems(v1Interface *interfaces.MockNasV1Interface, domains []string, err error) {
	mountTargets := nas.MountTargetsInDescribeFileSystems{}
	for _, domain := range domains {
		mountTargets.MountTarget = append(mountTargets.MountTarget, nas.MountTarget{
			MountTargetDomain: domain,
		})
	}
	v1Interface.EXPECT().DescribeFileSystems(gomock.Any()).Return(&nas.DescribeFileSystemsResponse{
		TotalCount: 1,
		FileSystems: nas.FileSystemsInDescribeFileSystems{
			FileSystem: []nas.FileSystem{{
				MountTargets: mountTargets,
			}},
		},
	}, err).AnyTimes()
}

func v1NasInterfaceExpectsDescribeMountTargets(v1Interface *interfaces.MockNasV1Interface, statuses []string, err error) {
	mountTargets := nas.MountTargetsInDescribeMountTargets{}
	for _, status := range statuses {
		mountTargets.MountTarget = append(mountTargets.MountTarget, nas.MountTarget{
			Status: status,
		})
	}
	v1Interface.EXPECT().DescribeMountTargets(gomock.Any()).Return(&nas.DescribeMountTargetsResponse{
		MountTargets: mountTargets,
	}, err).AnyTimes()
}

func resetFileSystemControllerStatus(controller *filesystemController) {
	controller.pvcFileSystemIDMap = sync.Map{}
	controller.pvcMountTargetMap = sync.Map{}
	controller.pvcProcessSuccess = sync.Map{}
}

func TestCreateVolume_CreateFileSystemError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerZoneIdMetadataResponder()
	prepareFilesystemControllerTestFakeK8sContext()
	prepareNasClientCredentials(t)

	controller := newTestFileSystemControllerWithExpects(t, func(v1Interface *interfaces.MockNasV1Interface) {
		v1NasInterfaceExpectsCreateFileSystem(v1Interface, "request-id", "file-system-id", errors.New("create file system error"))
	})

	actual, err := controller.CreateVolume(context.Background(), &csi.CreateVolumeRequest{
		Parameters: map[string]string{
			VpcID:     "vpc-id",
			VSwitchID: "vswitch-id",
		},
	})
	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestCreateVolume_TagResourcesError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerZoneIdMetadataResponder()
	prepareFilesystemControllerTestFakeK8sContext()
	prepareNasClientCredentials(t)

	controller := newTestFileSystemControllerWithExpects(t, func(v1Interface *interfaces.MockNasV1Interface) {
		v1NasInterfaceExpectsCreateFileSystem(v1Interface, "request-id", "file-system-id", nil)
		v1NasInterfaceExpectsTagResources(v1Interface, errors.New("tag resources error"))
		v1NasInterfaceExpectsCreateMountTarget(v1Interface, nil)
		v1NasInterfaceExpectsDescribeFileSystems(v1Interface, []string{"test.mount.target.domain"}, nil)
		v1NasInterfaceExpectsDescribeMountTargets(v1Interface, []string{"Active"}, nil)
	})
	controller.eventRecorder = utils.NewEventRecorder()
	controller.config.ClusterID = "cluster-id"

	actual, err := controller.CreateVolume(context.Background(), &csi.CreateVolumeRequest{
		Parameters: map[string]string{
			VpcID:     "vpc-id",
			VSwitchID: "vswitch-id",
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, actual)
}

func TestCreateVolume_CreateMountTargetError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerZoneIdMetadataResponder()
	prepareFilesystemControllerTestFakeK8sContext()
	prepareNasClientCredentials(t)

	controller := newTestFileSystemControllerWithExpects(t, func(v1Interface *interfaces.MockNasV1Interface) {
		v1NasInterfaceExpectsCreateFileSystem(v1Interface, "request-id", "file-system-id", nil)
		v1NasInterfaceExpectsTagResources(v1Interface, nil)
		v1NasInterfaceExpectsCreateMountTarget(v1Interface, errors.New("create mount target error"))
	})

	actual, err := controller.CreateVolume(context.Background(), &csi.CreateVolumeRequest{
		Parameters: map[string]string{
			VpcID:     "vpc-id",
			VSwitchID: "vswitch-id",
		},
	})
	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestCreateVolume_DescribeFileSystemsError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerZoneIdMetadataResponder()
	prepareFilesystemControllerTestFakeK8sContext()
	prepareNasClientCredentials(t)

	controller := newTestFileSystemControllerWithExpects(t, func(v1Interface *interfaces.MockNasV1Interface) {
		v1NasInterfaceExpectsCreateFileSystem(v1Interface, "request-id", "file-system-id", nil)
		v1NasInterfaceExpectsTagResources(v1Interface, nil)
		v1NasInterfaceExpectsCreateMountTarget(v1Interface, nil)
		v1NasInterfaceExpectsDescribeFileSystems(v1Interface, []string{"test.mount.target.domain"}, errors.New("describe file systems error"))
	})

	actual, err := controller.CreateVolume(context.Background(), createExtremeFileSystemRequest())
	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestCreateVolume_DescribeFileSystemsCountError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerZoneIdMetadataResponder()
	prepareFilesystemControllerTestFakeK8sContext()
	prepareNasClientCredentials(t)

	controller := newTestFileSystemControllerWithExpects(t, func(v1Interface *interfaces.MockNasV1Interface) {
		v1NasInterfaceExpectsCreateFileSystem(v1Interface, "request-id", "file-system-id", nil)
		v1NasInterfaceExpectsTagResources(v1Interface, nil)
		v1NasInterfaceExpectsCreateMountTarget(v1Interface, nil)
		v1Interface.EXPECT().DescribeFileSystems(gomock.Any()).Return(&nas.DescribeFileSystemsResponse{
			TotalCount: 2,
		}, nil).AnyTimes()
	})

	actual, err := controller.CreateVolume(context.Background(), createExtremeFileSystemRequest())
	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestCreateVolume_DescribeFileSystemsMountTargetCountTwoError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerZoneIdMetadataResponder()
	prepareFilesystemControllerTestFakeK8sContext()
	prepareNasClientCredentials(t)

	controller := newTestFileSystemControllerWithExpects(t, func(v1Interface *interfaces.MockNasV1Interface) {
		v1NasInterfaceExpectsCreateFileSystem(v1Interface, "request-id", "file-system-id", nil)
		v1NasInterfaceExpectsTagResources(v1Interface, nil)
		v1NasInterfaceExpectsCreateMountTarget(v1Interface, nil)
		v1Interface.EXPECT().DescribeFileSystems(gomock.Any()).Return(&nas.DescribeFileSystemsResponse{
			TotalCount: 1,
			FileSystems: nas.FileSystemsInDescribeFileSystems{
				FileSystem: []nas.FileSystem{{
					MountTargets: nas.MountTargetsInDescribeFileSystems{
						MountTarget: []nas.MountTarget{
							{}, {},
						},
					},
				}},
			},
		}, nil).AnyTimes()
	})

	actual, err := controller.CreateVolume(context.Background(), createExtremeFileSystemRequest())
	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestCreateVolume_DescribeMountTargetsError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerZoneIdMetadataResponder()
	prepareFilesystemControllerTestFakeK8sContext()
	prepareNasClientCredentials(t)

	controller := newTestFileSystemControllerWithExpects(t, func(v1Interface *interfaces.MockNasV1Interface) {
		v1NasInterfaceExpectsCreateFileSystem(v1Interface, "request-id", "file-system-id", nil)
		v1NasInterfaceExpectsTagResources(v1Interface, nil)
		v1NasInterfaceExpectsCreateMountTarget(v1Interface, nil)
		v1NasInterfaceExpectsDescribeFileSystems(v1Interface, []string{"test.mount.target.domain"}, nil)
		v1NasInterfaceExpectsDescribeMountTargets(v1Interface, []string{"Active"}, errors.New("describe mount targets error"))
	})

	actual, err := controller.CreateVolume(context.Background(), &csi.CreateVolumeRequest{
		Parameters: map[string]string{
			VpcID:     "vpc-id",
			VSwitchID: "vswitch-id",
		},
	})
	assert.Error(t, err)
	assert.Nil(t, actual)
}

func newTestFileSystemControllerWithExpects(t *testing.T, v1ClientMockExpects func(*interfaces.MockNasV1Interface)) *filesystemController {
	ctrl := gomock.NewController(t)
	prepareFilesystemControllerTestFakeK8sContext()
	factory := interfaces.NewMockNasClientFactory(ctrl)
	v1ClientMockExpects(factory.V1Client)

	return &filesystemController{
		config: &internal.ControllerConfig{
			Region:           fileSystemControllerTestRegion,
			NasClientFactory: factory,
			KubeClient: fake.NewSimpleClientset(&v1.StorageClass{
				ObjectMeta: metav1.ObjectMeta{
					Name: "mock-pv",
				},
				Parameters: map[string]string{
					RegionID: "cn-hangzhou-a",
				},
			}),
		},
	}
}

func TestDeleteVolume(t *testing.T) {
	type args struct {
		req *csi.DeleteVolumeRequest
		pv  *corev1.PersistentVolume
	}
	tests := []struct {
		name     string
		args     args
		expected *csi.DeleteVolumeResponse
		wantErr  bool
	}{
		{
			name: "Error-Empty nfs server",
			args: args{
				req: &csi.DeleteVolumeRequest{VolumeId: "volume-id"},
				pv: &corev1.PersistentVolume{
					Spec: corev1.PersistentVolumeSpec{
						PersistentVolumeSource: corev1.PersistentVolumeSource{
							CSI: &corev1.CSIPersistentVolumeSource{
								VolumeAttributes: map[string]string{
									"fileSystemId": "file-system-id",
									"deleteVolume": "true",
									"server":       "",
								},
							},
						},
					},
					Status: corev1.PersistentVolumeStatus{},
				},
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Error-Empty storage class name",
			args: args{
				req: &csi.DeleteVolumeRequest{VolumeId: "volume-id"},
				pv: &corev1.PersistentVolume{
					Spec: corev1.PersistentVolumeSpec{
						StorageClassName: "",
						PersistentVolumeSource: corev1.PersistentVolumeSource{
							CSI: &corev1.CSIPersistentVolumeSource{
								VolumeAttributes: map[string]string{
									"server": "nas.server.com",
								},
							},
						},
					},
				},
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Error-No specified persistent volume",
			args: args{
				req: &csi.DeleteVolumeRequest{VolumeId: "volume-id"},
				pv: &corev1.PersistentVolume{
					Spec: corev1.PersistentVolumeSpec{
						StorageClassName: "storage-class-name",
						PersistentVolumeSource: corev1.PersistentVolumeSource{
							CSI: &corev1.CSIPersistentVolumeSource{
								VolumeAttributes: map[string]string{
									"server": "nas.server.com",
								},
							},
						},
					},
				},
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Error-Empty file system id",
			args: args{
				req: &csi.DeleteVolumeRequest{VolumeId: "volume-id"},
				pv: &corev1.PersistentVolume{
					Spec: corev1.PersistentVolumeSpec{
						StorageClassName: "mock-pv",
						PersistentVolumeSource: corev1.PersistentVolumeSource{
							CSI: &corev1.CSIPersistentVolumeSource{
								VolumeAttributes: map[string]string{
									"fileSystemId": "",
									"server":       "nas.server.com",
									"deleteVolume": "true",
								},
							},
						},
					},
				},
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Success-Delete volume true",
			args: args{
				req: &csi.DeleteVolumeRequest{VolumeId: "volume-id"},
				pv: &corev1.PersistentVolume{
					Spec: corev1.PersistentVolumeSpec{
						StorageClassName: "mock-pv",
						PersistentVolumeSource: corev1.PersistentVolumeSource{
							CSI: &corev1.CSIPersistentVolumeSource{
								VolumeAttributes: map[string]string{
									"fileSystemId": "file-system-id",
									"server":       "nas.server.com",
									"deleteVolume": "true",
								},
							},
						},
					},
				},
			},
			expected: &csi.DeleteVolumeResponse{},
			wantErr:  false,
		},
		{
			name: "Success-Delete volume false",
			args: args{
				req: &csi.DeleteVolumeRequest{VolumeId: "volume-id"},
				pv: &corev1.PersistentVolume{
					Spec: corev1.PersistentVolumeSpec{
						StorageClassName: "mock-pv",
						PersistentVolumeSource: corev1.PersistentVolumeSource{
							CSI: &corev1.CSIPersistentVolumeSource{
								VolumeAttributes: map[string]string{
									"fileSystemId": "file-system-id",
									"server":       "nas.server.com",
								},
							},
						},
					},
				},
			},
			expected: &csi.DeleteVolumeResponse{},
			wantErr:  false,
		},
	}

	controller := newTestFileSystemControllerWithExpects(t, func(v1Interface *interfaces.MockNasV1Interface) {
		v1NasInterfaceExpectsDescribeMountTargets(v1Interface, []string{"Active"}, nil)
		v1NasInterfaceExpectsDeleteMountTarget(v1Interface, nil)
		v1NasInterfaceExpectsDeleteFileSystem(v1Interface, nil)
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := controller.DeleteVolume(context.Background(), tt.args.req, tt.args.pv)
			assert.Equal(t, tt.expected, actual)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func v1NasInterfaceExpectsDeleteMountTarget(v1Interface *interfaces.MockNasV1Interface, err error) {
	v1Interface.EXPECT().DeleteMountTarget(gomock.Any()).Return(&nas.DeleteMountTargetResponse{}, err).AnyTimes()
}

func v1NasInterfaceExpectsDeleteFileSystem(v1Interface *interfaces.MockNasV1Interface, err error) {
	v1Interface.EXPECT().DeleteFileSystem(gomock.Any()).Return(&nas.DeleteFileSystemResponse{}, err).AnyTimes()
}

func TestDeleteVolume_MountTargetNotFoundError(t *testing.T) {
	controller := newTestFileSystemControllerWithExpects(t, func(v1Interface *interfaces.MockNasV1Interface) {
		v1Interface.EXPECT().DescribeMountTargets(gomock.Any()).Return(
			nil, aliErrors.NewServerError(200, `{"Code": "InvalidParam.MountTargetDomain"}`, ""))
		v1NasInterfaceExpectsDeleteFileSystem(v1Interface, nil)
	})
	actual, err := controller.DeleteVolume(context.Background(), deleteVolumeRequest(), deleteVolumePV())
	assert.NoError(t, err)
	assert.Equal(t, &csi.DeleteVolumeResponse{}, actual)
}

func deleteVolumeRequest() *csi.DeleteVolumeRequest {
	return &csi.DeleteVolumeRequest{
		VolumeId: "volume-id",
	}
}

func deleteVolumePV() *corev1.PersistentVolume {
	return &corev1.PersistentVolume{
		Spec: corev1.PersistentVolumeSpec{
			StorageClassName: "mock-pv",
			PersistentVolumeSource: corev1.PersistentVolumeSource{
				CSI: &corev1.CSIPersistentVolumeSource{
					VolumeAttributes: map[string]string{
						"fileSystemId": "file-system-id",
						"server":       "nas.server.com",
						"deleteVolume": "true",
					},
				},
			},
		},
	}
}

func TestDeleteVolume_DeleteMountTargetError(t *testing.T) {
	controller := newTestFileSystemControllerWithExpects(t, func(v1Interface *interfaces.MockNasV1Interface) {
		v1NasInterfaceExpectsDescribeMountTargets(v1Interface, []string{"Active"}, nil)
		v1Interface.EXPECT().DeleteMountTarget(gomock.Any()).Return(&nas.DeleteMountTargetResponse{
			RequestId: "request-id",
		}, errors.New("delete mount target error"))
	})
	actual, err := controller.DeleteVolume(context.Background(), deleteVolumeRequest(), deleteVolumePV())
	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestDeleteVolume_DeleteFileSystemError(t *testing.T) {
	controller := newTestFileSystemControllerWithExpects(t, func(v1Interface *interfaces.MockNasV1Interface) {
		v1NasInterfaceExpectsDescribeMountTargets(v1Interface, []string{"Active"}, nil)
		v1NasInterfaceExpectsDeleteMountTarget(v1Interface, nil)
		v1Interface.EXPECT().DeleteFileSystem(gomock.Any()).Return(&nas.DeleteFileSystemResponse{
			RequestId: "request-id",
		}, errors.New("delete file system error"))
	})
	actual, err := controller.DeleteVolume(context.Background(), deleteVolumeRequest(), deleteVolumePV())
	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestGetNasVolumeOptions(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerZoneIdMetadataResponder()

	controller := newTestFileSystemController(t)
	tests := []struct {
		name     string
		args     *csi.CreateVolumeRequest
		expected *nasVolumeArgs
		wantErr  bool
	}{
		{
			name: "Default args",
			args: &csi.CreateVolumeRequest{
				Parameters: map[string]string{
					VpcID:     "vpc-id",
					VSwitchID: "vswitch-id",
				},
			},
			expected: &nasVolumeArgs{
				ProtocolType:    "NFS",
				StorageType:     "Performance",
				NetworkType:     "vpc",
				VpcID:           "vpc-id",
				VSwitchID:       "vswitch-id",
				AccessGroupName: "DEFAULT_VPC_GROUP_NAME",
			},
			wantErr: false,
		},
		{
			name: "Extreme file system type-Default args",
			args: createExtremeFileSystemRequest(),
			expected: &nasVolumeArgs{
				ProtocolType:    "NFS",
				StorageType:     "standard",
				FileSystemType:  "extreme",
				Capacity:        "100",
				EncryptType:     "0",
				NetworkType:     "vpc",
				VpcID:           "vpc-id",
				VSwitchID:       "vswitch-id",
				AccessGroupName: "DEFAULT_VPC_GROUP_NAME",
			},
			wantErr: false,
		},
		{
			name: "Region id-Non hangzhou finance",
			args: &csi.CreateVolumeRequest{
				Parameters: map[string]string{
					VpcID:     "vpc-id",
					VSwitchID: "vswitch-id",
					RegionID:  "cn-hangzhou-a",
				},
			},
			expected: &nasVolumeArgs{
				ProtocolType:    "NFS",
				StorageType:     "Performance",
				RegionID:        "cn-hangzhou-a",
				NetworkType:     "vpc",
				VpcID:           "vpc-id",
				VSwitchID:       "vswitch-id",
				AccessGroupName: "DEFAULT_VPC_GROUP_NAME",
			},
			wantErr: false,
		},
		{
			name: "Delete Volume-True",
			args: &csi.CreateVolumeRequest{
				Parameters: map[string]string{
					VpcID:        "vpc-id",
					VSwitchID:    "vswitch-id",
					DeleteVolume: "true",
				},
			},
			expected: &nasVolumeArgs{
				ProtocolType:    "NFS",
				StorageType:     "Performance",
				NetworkType:     "vpc",
				VpcID:           "vpc-id",
				VSwitchID:       "vswitch-id",
				AccessGroupName: "DEFAULT_VPC_GROUP_NAME",
				DeleteVolume:    true,
			},
			wantErr: false,
		},
		{
			name: "Delete Volume-False",
			args: &csi.CreateVolumeRequest{
				Parameters: map[string]string{
					VpcID:        "vpc-id",
					VSwitchID:    "vswitch-id",
					DeleteVolume: "false",
				},
			},
			expected: &nasVolumeArgs{
				ProtocolType:    "NFS",
				StorageType:     "Performance",
				NetworkType:     "vpc",
				VpcID:           "vpc-id",
				VSwitchID:       "vswitch-id",
				AccessGroupName: "DEFAULT_VPC_GROUP_NAME",
			},
			wantErr: false,
		},
		{
			name: "Invalid file system type",
			args: &csi.CreateVolumeRequest{
				Parameters: map[string]string{
					FileSystemType: "InvalidFileSystemType",
				},
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Invalid required bytes",
			args: &csi.CreateVolumeRequest{
				CapacityRange: &csi.CapacityRange{
					RequiredBytes: 1024,
				},
				Parameters: map[string]string{
					FileSystemType: "extreme",
				},
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Invalid network type",
			args: &csi.CreateVolumeRequest{
				Parameters: map[string]string{
					NetworkType: "InvalidNetworkType",
				},
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Absent vpcId",
			args: &csi.CreateVolumeRequest{
				Parameters: map[string]string{},
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Absent vSwitchId",
			args: &csi.CreateVolumeRequest{
				Parameters: map[string]string{
					VpcID: "",
				},
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Extreme file system type-Invalid storage type",
			args: &csi.CreateVolumeRequest{
				CapacityRange: &csi.CapacityRange{
					RequiredBytes: 100 * 1024 * 1024 * 1024,
				},
				Parameters: map[string]string{
					FileSystemType: "extreme",
					StorageType:    "InvalidStorageType",
				},
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Extreme file system type-Invalid encrypt type",
			args: &csi.CreateVolumeRequest{
				CapacityRange: &csi.CapacityRange{
					RequiredBytes: 100 * 1024 * 1024 * 1024,
				},
				Parameters: map[string]string{
					FileSystemType: "extreme",
					EncryptType:    "InvalidEncryptType",
				},
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Non-extreme file system type-Invalid storage type",
			args: &csi.CreateVolumeRequest{
				Parameters: map[string]string{
					StorageType: "InvalidStorageType",
				},
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Invalid protocol type",
			args: &csi.CreateVolumeRequest{
				Parameters: map[string]string{
					ProtocolType: "InvalidProtocolType",
				},
			},
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := controller.getNasVolumeOptions(tt.args)
			assert.Equal(t, tt.expected, actual)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestControllerExpandVolume(t *testing.T) {
	controller := newTestFileSystemController(t)
	actual, err := controller.ControllerExpandVolume(context.Background(), &csi.ControllerExpandVolumeRequest{}, &corev1.PersistentVolume{})
	assert.NoError(t, err)
	assert.Equal(t, &csi.ControllerExpandVolumeResponse{}, actual)
}
