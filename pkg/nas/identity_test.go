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

package nas

import (
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/require"
)

func TestSubstrateActorIdentityRemainsCredentialResource(t *testing.T) {
	context := map[string]string{
		"csi.alibabacloud.com/substrate-mode":  "true",
		"csi.alibabacloud.com/actor.uid":       "actor-id",
		"csi.alibabacloud.com/actor.name":      "actor-name",
		"csi.alibabacloud.com/actor.namespace": "actor-space",
		"csi.storage.k8s.io/pod.uid":           "worker-id",
		"authType":                             "agent-identity", "sandboxCredProviderName": "provider",
	}
	opt, _, err := parseVolumeContext(context)
	require.NoError(t, err)
	require.Equal(t, "actor-id", opt.SandboxId)
	options := appendJWTAuthOptions([]string{"tls"}, opt)
	resolved := jwtauth.ResolveOpts(mounterutils.IndexMountOptions(options))
	require.Equal(t, "actor-id", resolved.SandboxId)
	require.True(t, resolved.SubstrateMode)
	require.Equal(t, "worker-id", context["csi.storage.k8s.io/pod.uid"])
	context["sandboxId"] = "worker-id"
	_, _, err = parseVolumeContext(context)
	require.Error(t, err)
	delete(context, "sandboxId")
	context["csi.alibabacloud.com/actor.uid"] = ""
	_, _, err = parseVolumeContext(context)
	require.Error(t, err)
}
