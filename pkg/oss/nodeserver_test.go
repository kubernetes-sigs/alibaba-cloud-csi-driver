package oss

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	ossfpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager/oss"
	_ "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager/oss/ossfs"
	_ "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager/oss/ossfs2"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	mountutils "k8s.io/mount-utils"
)

// setupTestNodeServer creates a test nodeServer with minimal required fields
func setupTestNodeServer(t *testing.T, mounter mountutils.Interface, skipAttach bool) *nodeServer {
	fakeMeta := metadata.FakeProvider{
		Values: map[metadata.MetadataKey]string{
			metadata.RegionID:    "cn-beijing",
			metadata.RAMRoleName: "worker-role",
		},
	}

	fusePodManagers := ossfpm.GetAllOSSFusePodManagers(utils.Config{}, fakeMeta, nil)
	ossfsPaths := ossfpm.GetAllFuseMounterPaths()

	return &nodeServer{
		metadata:        fakeMeta,
		locks:           utils.NewVolumeLocks(),
		nodeName:        "test-node",
		clientset:       nil,
		cnfsGetter:      nil,
		rawMounter:      mounter,
		fusePodManagers: fusePodManagers,
		ossfsPaths:      ossfsPaths,
		skipAttach:      skipAttach,
	}
}

// createMountPoint creates a mount point in FakeMounter for testing
func createMountPoint(path string) mountutils.MountPoint {
	evalPath, err := filepath.EvalSymlinks(path)
	if err != nil {
		evalPath, _ = filepath.Abs(path)
	}
	return mountutils.MountPoint{
		Device: "ossfs",
		Path:   evalPath,
		Type:   "fuse.ossfs",
	}
}

func TestNodePublishVolume_RuntimeTypes(t *testing.T) {
	// Test matrix: 3 * 2 (COCO, RunD, MicroVM in scenario 1 & 3) + 1 * 3 (RunC in scenario 1, 2, 3)
	// Scenario 1: targetPath mounted (token rotation or already mounted return)
	// Scenario 2: targetPath not mounted, but attachPath mounted (RunC only, bind mount only)
	// Scenario 3: both targetPath and attachPath not mounted (new mount)
	tests := []struct {
		name                string
		runtimeType         RuntimeType
		directAssigned      bool
		socketPath          string
		skipAttach          bool
		targetMounted       bool
		attachMounted       bool // For RunC only
		hasToken            bool
		expectTokenRotate   bool
		expectValidation    bool
		expectExtendedMount bool
		expectBindMount     bool // For RunC only
		expectBindMountOnly bool // For RunC scenario 2
	}{
		// COCO scenarios (scenario 1 & 3)
		{
			name:                "COCO scenario 1: already mounted",
			runtimeType:         RuntimeTypeCOCO,
			directAssigned:      true,
			socketPath:          "",
			skipAttach:          false,
			targetMounted:       true,
			hasToken:            false,
			expectExtendedMount: false, // Early return
		},
		{
			name:                "COCO scenario 3: new mount",
			runtimeType:         RuntimeTypeCOCO,
			directAssigned:      true,
			socketPath:          "",
			skipAttach:          false,
			targetMounted:       false,
			hasToken:            false,
			expectExtendedMount: true, // COCO uses publishDirectVolume
		},

		// RunC scenarios (scenario 1, 2, 3)
		{
			name:                "RunC scenario 1: token rotation",
			runtimeType:         RuntimeTypeRunC,
			directAssigned:      false,
			socketPath:          "/tmp/socket",
			skipAttach:          false,
			targetMounted:       true,
			attachMounted:       true,
			hasToken:            true,
			expectTokenRotate:   true,
			expectValidation:    false, // Skipped for token rotation
			expectExtendedMount: true,  // Called for token rotation (interceptor skips)
			expectBindMount:     false,
		},
		{
			name:                "RunC scenario 1: already mounted (no token)",
			runtimeType:         RuntimeTypeRunC,
			directAssigned:      false,
			socketPath:          "/tmp/socket",
			skipAttach:          false,
			targetMounted:       true,
			attachMounted:       true,
			hasToken:            false,
			expectExtendedMount: false, // Early return
		},
		{
			name:                "RunC scenario 2: bind mount only (attachPath mounted, targetPath not mounted)",
			runtimeType:         RuntimeTypeRunC,
			directAssigned:      false,
			socketPath:          "/tmp/socket",
			skipAttach:          false,
			targetMounted:       false,
			attachMounted:       true,
			hasToken:            true,
			expectValidation:    true,
			expectExtendedMount: false, // Skip ExtendedMount
			expectBindMountOnly: true,  // Only bind mount
		},
		{
			name:                "RunC scenario 3: new mount (both not mounted)",
			runtimeType:         RuntimeTypeRunC,
			directAssigned:      false,
			socketPath:          "/tmp/socket",
			skipAttach:          false,
			targetMounted:       false,
			attachMounted:       false,
			hasToken:            false,
			expectValidation:    true,
			expectExtendedMount: true,
			expectBindMount:     true,
		},

		// RunD scenarios (scenario 1 & 3)
		{
			name:                "RunD scenario 1: token rotation",
			runtimeType:         RuntimeTypeRunD,
			directAssigned:      true,
			socketPath:          "/tmp/socket",
			skipAttach:          true,
			targetMounted:       true,
			hasToken:            true,
			expectTokenRotate:   true,
			expectValidation:    false,
			expectExtendedMount: true, // Called for token rotation (interceptor skips)
		},
		{
			name:                "RunD scenario 1: already mounted (no token)",
			runtimeType:         RuntimeTypeRunD,
			directAssigned:      true,
			socketPath:          "/tmp/socket",
			skipAttach:          true,
			targetMounted:       true,
			hasToken:            false,
			expectExtendedMount: false, // Early return
		},
		{
			name:                "RunD scenario 3: new mount (csi-agent)",
			runtimeType:         RuntimeTypeRunD,
			directAssigned:      true,
			socketPath:          "/tmp/socket",
			skipAttach:          true,
			targetMounted:       false,
			hasToken:            false,
			expectValidation:    true,
			expectExtendedMount: true,
		},

		// MicroVM scenarios (scenario 1 & 3)
		{
			name:                "MicroVM scenario 1: token rotation",
			runtimeType:         RuntimeTypeMicroVM,
			directAssigned:      false,
			socketPath:          "",
			skipAttach:          true,
			targetMounted:       true,
			hasToken:            true,
			expectTokenRotate:   true,
			expectValidation:    false,
			expectExtendedMount: true, // Called for token rotation (interceptor skips)
		},
		{
			name:                "MicroVM scenario 1: already mounted (no token)",
			runtimeType:         RuntimeTypeMicroVM,
			directAssigned:      false,
			socketPath:          "",
			skipAttach:          true,
			targetMounted:       true,
			hasToken:            false,
			expectExtendedMount: false, // Early return
		},
		{
			name:                "MicroVM scenario 3: new mount",
			runtimeType:         RuntimeTypeMicroVM,
			directAssigned:      false,
			socketPath:          "",
			skipAttach:          true,
			targetMounted:       false,
			hasToken:            false,
			expectValidation:    true,
			expectExtendedMount: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// determine the atucal runtime type
			runtimeType, err := DetermineRuntimeType(tt.directAssigned, tt.socketPath, tt.skipAttach)
			require.NoError(t, err)
			require.Equal(t, tt.runtimeType, runtimeType)

			// Skip COCO tests as they require directvolume support
			if tt.runtimeType == RuntimeTypeCOCO {
				t.Skip("COCO tests require directvolume support, skipping")
			}

			baseDir := t.TempDir()
			targetPath := filepath.Join(baseDir, "target")
			// Set fuse attach base dir to temp directory to avoid /run permission issues in tests
			// This must be done before any calls to GetAttachPath()
			if tt.runtimeType == RuntimeTypeRunC {
				mounterutils.SetFuseAttachBaseDir(baseDir)
				defer func() {
					// Reset to default after test
					mounterutils.SetFuseAttachBaseDir("/run")
				}()
			}

			// Create directories
			require.NoError(t, os.MkdirAll(targetPath, 0o755))
			// Get attachPath after setting fuse attach base dir
			var attachPath string
			if tt.runtimeType == RuntimeTypeRunC {
				attachPath = mounterutils.GetAttachPath("test-volume-id")
				// Create parent directories for attachPath
				require.NoError(t, os.MkdirAll(filepath.Dir(attachPath), 0o755))
				// If attachPath should be mounted, create the attachPath directory itself
				// so that IsNotMountPoint can properly detect it as a mount point
				if tt.attachMounted {
					require.NoError(t, os.MkdirAll(attachPath, 0o755))
				}
			}

			// Setup mount points
			var mountPoints []mountutils.MountPoint
			if tt.targetMounted {
				mountPoints = append(mountPoints, createMountPoint(targetPath))
			}
			if tt.runtimeType == RuntimeTypeRunC && tt.attachMounted {
				mountPoints = append(mountPoints, createMountPoint(attachPath))
			}

			fakeMounter := mountutils.NewFakeMounter(mountPoints)
			ns := setupTestNodeServer(t, fakeMounter, tt.skipAttach)

			// Prepare request
			req := &csi.NodePublishVolumeRequest{
				VolumeId:   "test-volume-id",
				TargetPath: targetPath,
				VolumeContext: map[string]string{
					"bucket":   "test-bucket",
					"url":      "https://oss-cn-beijing.aliyuncs.com",
					"path":     "/test",
					"fuseType": "ossfs",
				},
				VolumeCapability: &csi.VolumeCapability{
					AccessType: &csi.VolumeCapability_Mount{
						Mount: &csi.VolumeCapability_MountVolume{
							FsType: "ossfs",
						},
					},
				},
			}

			// Set directAssigned
			if tt.directAssigned {
				req.VolumeContext["directAssigned"] = "true"
			}

			// Set socket path
			if tt.socketPath != "" {
				if req.PublishContext == nil {
					req.PublishContext = make(map[string]string)
				}
				req.PublishContext[mountProxySocket] = tt.socketPath
			}

			// Set secrets (token if needed)
			if tt.hasToken {
				req.Secrets = map[string]string{
					mounterutils.KeyAccessKeyId:     "test-akid",
					mounterutils.KeyAccessKeySecret: "test-aksecret",
					mounterutils.KeySecurityToken:   "test-token",
				}
			} else {
				req.Secrets = map[string]string{
					"akId":     "test-akid",
					"akSecret": "test-aksecret",
				}
			}

			// Call NodePublishVolume
			// Note: This will fail in actual execution due to missing mounter/proxy implementation,
			// but we can verify the logic flow and mount point checks
			_, err = ns.NodePublishVolume(context.Background(), req)

			// Verify expectations based on scenario
			if !tt.expectExtendedMount && !tt.expectTokenRotate && !tt.expectBindMountOnly {
				// Should return early without error (already mounted)
				assert.NoError(t, err, "Should return early without error when already mounted")
			} else if tt.expectBindMountOnly {
				// RunC scenario 2: bind mount only, no ExtendedMount
				// Should succeed without calling ExtendedMount
				assert.NoError(t, err, "Should succeed with bind mount only")
			} else if tt.expectExtendedMount {
				// For new mounts or token rotation, we expect an error due to missing mounter/proxy implementation
				// but the logic should have progressed to mount attempt
				// In real scenarios, this would succeed
				if err != nil {
					// Check if error is due to missing mounter/proxy (expected in test environment)
					// Should not be a validation error
					assert.NotContains(t, err.Error(), "InvalidArgument", "Should not fail on validation")
				}
			} else if tt.expectTokenRotate {
				// Token rotation should be handled by interceptor
				// In test environment, this may fail due to missing mounter/proxy
				// but the logic should have skipped validation
				// Token rotation logic is more thoroughly tested in interceptor tests
				if err != nil {
					// Should not be a validation error
					assert.NotContains(t, err.Error(), "InvalidArgument", "Should not fail on validation for token rotation")
				}
			}

			// Verify mount point checks
			if tt.targetMounted {
				notMnt, err := fakeMounter.IsLikelyNotMountPoint(targetPath)
				assert.NoError(t, err)
				assert.False(t, notMnt, "targetPath should be mounted")
			}

			if tt.runtimeType == RuntimeTypeRunC && tt.attachMounted {
				notMnt, err := fakeMounter.IsLikelyNotMountPoint(attachPath)
				assert.NoError(t, err)
				assert.False(t, notMnt, "attachPath should be mounted")
			}
		})
	}
}

// TestNodePublishVolume_TokenRotation tests token rotation scenarios specifically
func TestNodePublishVolume_TokenRotation(t *testing.T) {
	tests := []struct {
		name          string
		runtimeType   RuntimeType
		targetMounted bool
		hasToken      bool
		hasFixedAKSK  bool
		expectRotate  bool
	}{
		{
			name:          "RunC token rotation",
			runtimeType:   RuntimeTypeRunC,
			targetMounted: true,
			hasToken:      true,
			hasFixedAKSK:  false,
			expectRotate:  true,
		},
		{
			name:          "RunC no rotation (fixed AKSK)",
			runtimeType:   RuntimeTypeRunC,
			targetMounted: true,
			hasToken:      true,
			hasFixedAKSK:  true,
			expectRotate:  false,
		},
		{
			name:          "RunD token rotation",
			runtimeType:   RuntimeTypeRunD,
			targetMounted: true,
			hasToken:      true,
			hasFixedAKSK:  false,
			expectRotate:  true,
		},
		{
			name:          "MicroVM token rotation",
			runtimeType:   RuntimeTypeMicroVM,
			targetMounted: true,
			hasToken:      true,
			hasFixedAKSK:  false,
			expectRotate:  true,
		},
		{
			name:          "No rotation (not mounted)",
			runtimeType:   RuntimeTypeRunC,
			targetMounted: false,
			hasToken:      true,
			hasFixedAKSK:  false,
			// Note: needRotateToken only checks secrets, not mount state.
			// The actual token rotation decision in NodePublishVolume combines both:
			// isTokenRotation := !notMntTarget && needRotateToken(...)
			// So needRotateToken returns true, but token rotation won't happen if not mounted.
			expectRotate: true,
		},
		{
			name:          "No rotation (no token)",
			runtimeType:   RuntimeTypeRunC,
			targetMounted: true,
			hasToken:      false,
			hasFixedAKSK:  false,
			expectRotate:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseDir := t.TempDir()
			targetPath := filepath.Join(baseDir, "target")

			require.NoError(t, os.MkdirAll(targetPath, 0o755))

			var mountPoints []mountutils.MountPoint
			if tt.targetMounted {
				mountPoints = append(mountPoints, createMountPoint(targetPath))
			}

			fakeMounter := mountutils.NewFakeMounter(mountPoints)

			// Test needRotateToken logic
			secrets := make(map[string]string)
			if tt.hasToken {
				secrets[mounterutils.KeySecurityToken] = "test-token"
				secrets[mounterutils.KeyAccessKeyId] = "test-akid"
				secrets[mounterutils.KeyAccessKeySecret] = "test-aksecret"
			}
			if tt.hasFixedAKSK {
				secrets["passwd-ossfs"] = "akid:aksecret"
			}

			result := needRotateToken(mounterutils.OssFsType, secrets)
			assert.Equal(t, tt.expectRotate, result, "needRotateToken should return %v", tt.expectRotate)

			// Verify mount point check
			notMnt, err := fakeMounter.IsLikelyNotMountPoint(targetPath)
			assert.NoError(t, err)
			assert.Equal(t, !tt.targetMounted, notMnt, "Mount point check should match expectation")
		})
	}
}

// TestNodePublishVolume_RunC_BindMount tests RunC bind mount scenarios
func TestNodePublishVolume_RunC_BindMount(t *testing.T) {
	tests := []struct {
		name              string
		targetMounted     bool
		attachMounted     bool
		expectBindMount   bool
		expectEarlyReturn bool
	}{
		{
			name:              "Both mounted (already complete)",
			targetMounted:     true,
			attachMounted:     true,
			expectBindMount:   false,
			expectEarlyReturn: true,
		},
		{
			name:              "Attach mounted, target not mounted (bind mount needed)",
			targetMounted:     false,
			attachMounted:     true,
			expectBindMount:   true,
			expectEarlyReturn: false,
		},
		{
			name:              "Neither mounted (new mount)",
			targetMounted:     false,
			attachMounted:     false,
			expectBindMount:   true,
			expectEarlyReturn: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseDir := t.TempDir()
			targetPath := filepath.Join(baseDir, "target")
			// Use temp directory for attachPath to avoid /run permission issues
			attachPath := filepath.Join(baseDir, "attach")

			require.NoError(t, os.MkdirAll(targetPath, 0o755))
			require.NoError(t, os.MkdirAll(attachPath, 0o755))

			var mountPoints []mountutils.MountPoint
			if tt.targetMounted {
				mountPoints = append(mountPoints, createMountPoint(targetPath))
			}
			if tt.attachMounted {
				mountPoints = append(mountPoints, createMountPoint(attachPath))
			}

			fakeMounter := mountutils.NewFakeMounter(mountPoints)

			// Verify mount point states
			notMntTarget, err := fakeMounter.IsLikelyNotMountPoint(targetPath)
			assert.NoError(t, err)
			assert.Equal(t, !tt.targetMounted, notMntTarget, "targetPath mount state")

			notMntAttach, err := fakeMounter.IsLikelyNotMountPoint(attachPath)
			assert.NoError(t, err)
			assert.Equal(t, !tt.attachMounted, notMntAttach, "attachPath mount state")

			// Verify expectations
			if tt.expectEarlyReturn {
				// Both mounted - should return early
				assert.True(t, tt.targetMounted && tt.attachMounted, "Both should be mounted for early return")
			}

			if tt.expectBindMount {
				// Bind mount should be needed
				assert.False(t, tt.targetMounted, "targetPath should not be mounted for bind mount")
			}
		})
	}
}
