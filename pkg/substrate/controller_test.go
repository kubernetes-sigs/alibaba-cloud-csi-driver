/*
Copyright 2026 The Kubernetes Authors.
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

package substrate

import (
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestLogicalVolumeLifecycle(t *testing.T) {
	controller := &Controller{}
	req := &csi.CreateVolumeRequest{Name: testID, Parameters: actorMetadataFixture(), VolumeCapabilities: []*csi.VolumeCapability{publishFixture().VolumeCapability}, CapacityRange: &csi.CapacityRange{RequiredBytes: 1024}}
	created, err := controller.CreateVolume(t.Context(), req)
	require.NoError(t, err)
	require.Equal(t, testID, created.Volume.VolumeId)
	require.Equal(t, int64(1024), created.Volume.CapacityBytes)
	require.Equal(t, "true", created.Volume.VolumeContext["csi.alibabacloud.com/substrate-mode"])
	wantContext := actorMetadataFixture()
	wantContext["csi.alibabacloud.com/substrate-mode"] = "true"
	require.Equal(t, wantContext, created.Volume.VolumeContext)
	retried, err := controller.CreateVolume(t.Context(), req)
	require.NoError(t, err)
	require.Equal(t, created.Volume, retried.Volume)
	_, err = controller.DeleteVolume(t.Context(), &csi.DeleteVolumeRequest{VolumeId: testID})
	require.NoError(t, err)
	_, err = controller.DeleteVolume(t.Context(), &csi.DeleteVolumeRequest{VolumeId: testID})
	require.NoError(t, err)
}

func TestControllerDoesNotResolveBackendConfiguration(t *testing.T) {
	parameters := actorMetadataFixture()
	parameters[PublishRequestsAnnotation] = "not JSON"
	parameters["server"] = "not-a-storage-server"
	created, err := (&Controller{}).CreateVolume(t.Context(), &csi.CreateVolumeRequest{Name: testID, Parameters: parameters, VolumeCapabilities: []*csi.VolumeCapability{publishFixture().VolumeCapability}})
	require.NoError(t, err)
	wantContext := actorMetadataFixture()
	wantContext["csi.alibabacloud.com/substrate-mode"] = "true"
	require.Equal(t, wantContext, created.Volume.VolumeContext)
}

func TestControllerRejectsInvalidVolumeRequests(t *testing.T) {
	for _, tc := range []struct {
		name string
		req  *csi.CreateVolumeRequest
	}{
		{"empty", &csi.CreateVolumeRequest{}},
		{"snapshot", &csi.CreateVolumeRequest{Name: testID, Parameters: actorMetadataFixture(), VolumeContentSource: &csi.VolumeContentSource{}}},
		{"invalid capacity", &csi.CreateVolumeRequest{Name: testID, Parameters: actorMetadataFixture(), CapacityRange: &csi.CapacityRange{RequiredBytes: 2048, LimitBytes: 1024}}},
		{"negative capacity", &csi.CreateVolumeRequest{Name: testID, Parameters: actorMetadataFixture(), CapacityRange: &csi.CapacityRange{RequiredBytes: -1}}},
		{"negative limit", &csi.CreateVolumeRequest{Name: testID, Parameters: actorMetadataFixture(), CapacityRange: &csi.CapacityRange{LimitBytes: -1}}},
		{"secrets", &csi.CreateVolumeRequest{Name: testID, Parameters: actorMetadataFixture(), Secrets: map[string]string{"key": "not-allowed"}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := (&Controller{}).CreateVolume(t.Context(), tc.req)
			require.Equal(t, codes.InvalidArgument, status.Code(err))
		})
	}
}

func TestCreateAcceptsAnyVolumeName(t *testing.T) {
	// Substrate owns its logical naming and has not stabilized it, so the bridge
	// keys identity off the Actor parameters and, at publish, the mount target.
	name := "pvc-0b60f2e4-3c8e-4a12-b2a6-9d5e4c8a2f33"
	created, err := (&Controller{}).CreateVolume(t.Context(), &csi.CreateVolumeRequest{Name: name, Parameters: actorMetadataFixture()})
	require.NoError(t, err)
	require.Equal(t, name, created.Volume.VolumeId)
}

func TestCreateIgnoresDeclaredCapabilities(t *testing.T) { // A logical volume carries no backend, so the capability the caller declares
	// is not interpreted here: what gets mounted is the inner request validated
	// at publish.
	created, err := (&Controller{}).CreateVolume(t.Context(), &csi.CreateVolumeRequest{
		Name:               testID,
		Parameters:         actorMetadataFixture(),
		VolumeCapabilities: []*csi.VolumeCapability{{AccessType: &csi.VolumeCapability_Block{Block: &csi.VolumeCapability_BlockVolume{}}}},
	})
	require.NoError(t, err)
	require.Equal(t, testID, created.Volume.VolumeId)
}

func TestValidateVolumeCapabilitiesAnswersFilesystemOnly(t *testing.T) {
	controller := &Controller{}
	confirmed, err := controller.ValidateVolumeCapabilities(t.Context(), &csi.ValidateVolumeCapabilitiesRequest{
		VolumeId:           testID,
		VolumeCapabilities: []*csi.VolumeCapability{publishFixture().VolumeCapability},
		VolumeContext:      actorMetadataFixture(),
	})
	require.NoError(t, err)
	require.NotNil(t, confirmed.Confirmed)
	require.Equal(t, actorMetadataFixture(), confirmed.Confirmed.VolumeContext)
	unconfigured, err := controller.ValidateVolumeCapabilities(t.Context(), &csi.ValidateVolumeCapabilitiesRequest{VolumeId: testID})
	require.NoError(t, err)
	require.Nil(t, unconfigured.Confirmed)
	require.NotEmpty(t, unconfigured.Message)
	rejected, err := controller.ValidateVolumeCapabilities(t.Context(), &csi.ValidateVolumeCapabilitiesRequest{
		VolumeId: testID,
		VolumeCapabilities: []*csi.VolumeCapability{{
			AccessType: &csi.VolumeCapability_Block{Block: &csi.VolumeCapability_BlockVolume{}},
			AccessMode: &csi.VolumeCapability_AccessMode{Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER},
		}},
	})
	require.NoError(t, err)
	require.Nil(t, rejected.Confirmed)
	require.NotEmpty(t, rejected.Message)
}

func TestControllerRejectsMissingActorParameters(t *testing.T) {
	for _, tc := range []struct{ key, value string }{
		{"csi.alibabacloud.com/actor.uid", ""},
		{"csi.alibabacloud.com/actor.name", ""},
		{"csi.alibabacloud.com/actor.namespace", ""},
	} {
		t.Run(tc.key+"="+tc.value, func(t *testing.T) {
			parameters := actorMetadataFixture()
			parameters[tc.key] = tc.value
			_, err := (&Controller{}).CreateVolume(t.Context(), &csi.CreateVolumeRequest{Name: testID, Parameters: parameters, VolumeCapabilities: []*csi.VolumeCapability{publishFixture().VolumeCapability}})
			require.Equal(t, codes.InvalidArgument, status.Code(err))
		})
	}
	_, err := (&Controller{}).CreateVolume(t.Context(), &csi.CreateVolumeRequest{Name: testID, VolumeCapabilities: []*csi.VolumeCapability{publishFixture().VolumeCapability}})
	require.Equal(t, codes.InvalidArgument, status.Code(err))
}

func TestCreateEchoesOnlyActorIdentity(t *testing.T) {
	controller := &Controller{}
	parameters := actorMetadataFixture()
	parameters[PodUIDKey] = "a-pod-uid"
	request := &csi.CreateVolumeRequest{Name: testID, Parameters: parameters, VolumeCapabilities: []*csi.VolumeCapability{writablePublishFixture().VolumeCapability}}
	first, err := controller.CreateVolume(t.Context(), request)
	require.NoError(t, err)
	for key, value := range actorMetadataFixture() {
		require.Equal(t, value, first.Volume.VolumeContext[key])
	}
	require.NotContains(t, first.Volume.VolumeContext, PodUIDKey)
	parameters[PodUIDKey] = "another-pod-uid"
	second, err := controller.CreateVolume(t.Context(), request)
	require.NoError(t, err)
	require.Equal(t, first.Volume.VolumeContext, second.Volume.VolumeContext)
	// Declared identity is echoed, never cross-checked with the volume name: the
	// Actor API is the only authority, and it is consulted at publish.
	parameters["csi.alibabacloud.com/actor.name"] = "wrong-actor"
	parameters["csi.alibabacloud.com/actor.uid"] = "00000000-0000-0000-0000-000000000000"
	third, err := controller.CreateVolume(t.Context(), request)
	require.NoError(t, err)
	require.Equal(t, "wrong-actor", third.Volume.VolumeContext["csi.alibabacloud.com/actor.name"])
	require.Equal(t, "00000000-0000-0000-0000-000000000000", third.Volume.VolumeContext["csi.alibabacloud.com/actor.uid"])
}

func TestCheckCreateRequestRejectsWhatTheBridgeCannotSatisfy(t *testing.T) {
	for _, tc := range []struct {
		name   string
		mutate func(*csi.CreateVolumeRequest)
	}{
		{"content source", func(req *csi.CreateVolumeRequest) { req.VolumeContentSource = &csi.VolumeContentSource{} }},
		{"secrets", func(req *csi.CreateVolumeRequest) { req.Secrets = map[string]string{"key": "value"} }},
		{"negative capacity", func(req *csi.CreateVolumeRequest) { req.CapacityRange = &csi.CapacityRange{RequiredBytes: -1} }},
		{"required above limit", func(req *csi.CreateVolumeRequest) {
			req.CapacityRange = &csi.CapacityRange{RequiredBytes: 200, LimitBytes: 100}
		}},
		{"missing actor UID", func(req *csi.CreateVolumeRequest) { delete(req.Parameters, "csi.alibabacloud.com/actor.uid") }},
		{"missing actor name", func(req *csi.CreateVolumeRequest) { delete(req.Parameters, "csi.alibabacloud.com/actor.name") }},
		{"missing actor namespace", func(req *csi.CreateVolumeRequest) { delete(req.Parameters, "csi.alibabacloud.com/actor.namespace") }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			req := &csi.CreateVolumeRequest{Name: "any-volume-name", Parameters: actorMetadataFixture()}
			tc.mutate(req)
			require.Equal(t, codes.InvalidArgument, status.Code(checkCreateRequest(req)))
		})
	}
	require.NoError(t, checkCreateRequest(&csi.CreateVolumeRequest{Name: "any-volume-name", Parameters: actorMetadataFixture()}))
}

func TestActorAttributesCarriesIdentityOnly(t *testing.T) {
	parameters := actorMetadataFixture()
	parameters["provisioner-only-parameter"] = "must-not-travel-to-publish"
	require.Equal(t, map[string]string{
		"csi.alibabacloud.com/substrate-mode":  "true",
		"csi.alibabacloud.com/actor.uid":       testUID,
		"csi.alibabacloud.com/actor.name":      "actor",
		"csi.alibabacloud.com/actor.namespace": "storage-test",
	}, actorAttributes(parameters))
}
