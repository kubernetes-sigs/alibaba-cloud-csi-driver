// Copyright 2026 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package chart_test

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"slices"
	"strings"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	yamlutil "k8s.io/apimachinery/pkg/util/yaml"
)

type renderedResource struct {
	Kind     string            `json:"kind"`
	Metadata metav1.ObjectMeta `json:"metadata"`
	Spec     json.RawMessage   `json:"spec"`
	Data     map[string]string `json:"data"`
}

func renderChart(t *testing.T, options ...string) []renderedResource {
	t.Helper()
	if _, err := exec.LookPath("helm"); err != nil {
		t.Skip("helm is required")
	}
	args := append([]string{"template", "test", ".", "--namespace", "csi-test"}, options...)
	cmd := exec.CommandContext(t.Context(), "helm", args...)
	cmd.Env = append(os.Environ(), "KUBECONFIG=/dev/null")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("helm template: %v\n%s", err, stderr.String())
	}
	decoder := yamlutil.NewYAMLOrJSONDecoder(bytes.NewReader(output), 4096)
	var resources []renderedResource
	for {
		var resource renderedResource
		if err := decoder.Decode(&resource); err == io.EOF {
			break
		} else if err != nil {
			t.Fatal(err)
		}
		if resource.Kind != "" {
			resources = append(resources, resource)
		}
	}
	return resources
}

func findResource(t *testing.T, resources []renderedResource, kind, name string) renderedResource {
	t.Helper()
	for _, resource := range resources {
		if resource.Kind == kind && resource.Metadata.Name == name {
			return resource
		}
	}
	t.Fatalf("missing %s/%s", kind, name)
	return renderedResource{}
}

func container(t *testing.T, pod corev1.PodSpec, name string) corev1.Container {
	t.Helper()
	for _, c := range pod.Containers {
		if c.Name == name {
			return c
		}
	}
	t.Fatalf("missing container %s", name)
	return corev1.Container{}
}

func driverNames(c corev1.Container) []string {
	for _, arg := range c.Args {
		if strings.HasPrefix(arg, "--driver=") {
			return strings.Split(strings.TrimPrefix(arg, "--driver="), ",")
		}
	}
	return nil
}

func TestSubstrateDisabledDoesNotAddResources(t *testing.T) {
	for _, resource := range renderChart(t) {
		if strings.Contains(resource.Metadata.Name, "substrate") || resource.Metadata.Name == "ate-storage" {
			t.Fatalf("unexpected default resource: %s/%s", resource.Kind, resource.Metadata.Name)
		}
	}
}

func TestSubstrateEnabledWiresControllerNodeAndTrust(t *testing.T) {
	resources := renderChart(t, "--set", "enableSubstrate=true", "--set", "deploy.featureGates=AlinasMountProxy=true")
	for _, resource := range resources {
		if (resource.Kind == "Deployment" && resource.Metadata.Name == "csi-substrate-controller") || (resource.Kind == "DaemonSet" && resource.Metadata.Name == "csi-plugin-substrate") || (resource.Kind == "ServiceAccount" && resource.Metadata.Name == "alicloud-csi-substrate") {
			t.Fatalf("unexpected standalone bridge resource %s/%s", resource.Kind, resource.Metadata.Name)
		}
	}
	var deployment appsv1.DeploymentSpec
	if err := json.Unmarshal(findResource(t, resources, "Deployment", "csi-provisioner").Spec, &deployment); err != nil {
		t.Fatal(err)
	}
	driver := container(t, deployment.Template.Spec, "csi-provisioner")
	if !slices.Contains(driverNames(driver), "nas") || !slices.Contains(driverNames(driver), "substrate") {
		t.Fatal("controller must keep NAS and add bridge")
	}
	args := strings.Join(driver.Args, " ")
	for _, flag := range []string{"substrate", "--run-controller-service=true", "--run-node-service=false"} {
		if !strings.Contains(args, flag) {
			t.Errorf("missing controller flag %s", flag)
		}
	}
	if strings.Contains(args, "--substrate-api-") {
		t.Fatal("controller must not require the Actor API")
	}
	for _, volume := range deployment.Template.Spec.Volumes {
		if volume.Name == "substrate-api-token" {
			t.Fatal("controller must not project an Actor API token")
		}
	}
	container(t, deployment.Template.Spec, "substrate-controller-proxy")
	if deployment.Template.Spec.ServiceAccountName != "alicloud-csi-provisioner" {
		t.Fatal("controller service account must be preserved")
	}
	var daemon appsv1.DaemonSetSpec
	if err := json.Unmarshal(findResource(t, resources, "DaemonSet", "csi-plugin").Spec, &daemon); err != nil {
		t.Fatal(err)
	}
	node := container(t, daemon.Template.Spec, "csi-plugin")
	for _, flag := range []string{"--substrate-api-endpoint=api.ate-system.svc:443", "--substrate-api-ca-file=/run/podidentity.podcert.ate.dev/trust-bundle.pem", "--substrate-api-token-file=/run/ateapi/token"} {
		if !slices.Contains(node.Args, flag) {
			t.Errorf("missing node flag %s", flag)
		}
	}
	if !slices.Contains(driverNames(node), "nas") || !slices.Contains(driverNames(node), "substrate") {
		t.Fatal("node must keep NAS and add bridge")
	}
	if strings.Contains(strings.Join(node.Args, " "), "--mount-proxy-sock=") {
		t.Fatal("Substrate must not override the shared OSS proxy flag")
	}
	if strings.Contains(strings.Join(node.Args, " "), "--nas-mount-proxy-sock=") {
		t.Fatal("Substrate must not introduce a NAS-specific socket override")
	}
	if !slices.Contains(node.Args, "--feature-gates=AlinasMountProxy=true") {
		t.Fatal("the existing NAS feature gate must be preserved")
	}
	if daemon.Template.Spec.ServiceAccountName != "alicloud-csi-node" {
		t.Fatal("node service account must be preserved")
	}
	for _, init := range daemon.Template.Spec.InitContainers {
		if strings.Contains(strings.Join(init.Args, " "), "substrate.csi.alibabacloud.com") {
			t.Fatal("bridge is not an init driver")
		}
	}
	registrar := container(t, daemon.Template.Spec, "substrate-driver-registrar")
	if !strings.Contains(strings.Join(registrar.Args, " "), "/var/lib/kubelet/csi-plugins/substrate.csi.alibabacloud.com/csi.sock") {
		t.Fatal("registrar path mismatch")
	}
	propagated, persistent, ca, token, clientCert := false, false, false, false, false
	for _, m := range node.VolumeMounts {
		if m.Name == "substrate-actors" && m.MountPropagation != nil && *m.MountPropagation == corev1.MountPropagationBidirectional {
			propagated = true
		}
	}
	for _, v := range daemon.Template.Spec.Volumes {
		if v.Name == "substrate-plugin-dir" && v.HostPath != nil && v.HostPath.Path == "/var/lib/kubelet/csi-plugins/substrate.csi.alibabacloud.com" {
			persistent = true
		}
		if v.Projected != nil {
			for _, s := range v.Projected.Sources {
				if s.ClusterTrustBundle != nil && s.ClusterTrustBundle.SignerName != nil && *s.ClusterTrustBundle.SignerName == "servicedns.podcert.ate.dev/identity" && s.ClusterTrustBundle.Path == "trust-bundle.pem" {
					ca = true
				}
				if s.ServiceAccountToken != nil && s.ServiceAccountToken.Audience == "api.ate-system.svc" {
					token = true
				}
				if s.PodCertificate != nil && s.PodCertificate.SignerName == "podidentity.podcert.ate.dev/identity" {
					clientCert = true
				}
			}
		}
	}
	if !propagated || !persistent || !ca || !token || !clientCert {
		t.Fatal("shared workload needs actor propagation, persistent state, CA, token and PodCertificate")
	}
	findResource(t, resources, "CSIDriver", "substrate.csi.alibabacloud.com")
	findResource(t, resources, "StorageClass", "ate-storage")
	findResource(t, resources, "CSIDriverConfig", "substrate-csi-bridge")
	findResource(t, resources, "CSIDriverConfig", "substrate-nas")
	var service corev1.ServiceSpec
	if err := json.Unmarshal(findResource(t, resources, "Service", "csi-provisioner").Spec, &service); err != nil {
		t.Fatal(err)
	}
	if service.Selector["app"] != "csi-provisioner" || len(service.Ports) != 2 {
		t.Fatal("NAS and bridge must share the existing controller service")
	}
	config := findResource(t, resources, "ConfigMap", "csi-provisioner-substrate")
	for _, socket := range []string{"/csi/nasplugin.csi.alibabacloud.com/csi.sock", "/csi/substrate.csi.alibabacloud.com/csi.sock"} {
		if !strings.Contains(config.Data["envoy.yaml"], socket) {
			t.Fatalf("missing backend %s", socket)
		}
	}
	if strings.Count(config.Data["envoy.yaml"], "require_client_certificate: true") != 2 {
		t.Fatal("both controller endpoints require mTLS")
	}
}

func TestSubstrateCustomPaths(t *testing.T) {
	resources := renderChart(t, "--set", "enableSubstrate=true", "--set", "deploy.kubeletRootDir=/custom/kubelet", "--set", "substrate.actorRoot=/custom/actors")
	var daemon appsv1.DaemonSetSpec
	if err := json.Unmarshal(findResource(t, resources, "DaemonSet", "csi-plugin").Spec, &daemon); err != nil {
		t.Fatal(err)
	}
	driver := container(t, daemon.Template.Spec, "csi-plugin")
	if !strings.Contains(strings.Join(driver.Args, " "), "--substrate-actor-root=/custom/actors") {
		t.Fatal("actor root override not forwarded")
	}
	config := findResource(t, resources, "CSIDriverConfig", "substrate-csi-bridge")
	if !strings.Contains(string(config.Spec), "unix:///custom/kubelet/csi-plugins/substrate.csi.alibabacloud.com/csi.sock") {
		t.Fatal("node socket override mismatch")
	}
}

func TestSubstrateRejectsUnsafeActorRoot(t *testing.T) {
	if _, err := exec.LookPath("helm"); err != nil {
		t.Skip("helm is required")
	}
	cmd := exec.CommandContext(t.Context(), "helm", "template", "test", ".", "--set", "enableSubstrate=true", "--set", "substrate.actorRoot=/")
	cmd.Env = append(os.Environ(), "KUBECONFIG=/dev/null")
	if err := cmd.Run(); err == nil {
		t.Fatal("root filesystem must not be an actor root")
	}
}

func TestSubstrateComponentSwitches(t *testing.T) {
	for _, tc := range []struct{ option, forbiddenKind, forbiddenName string }{
		{"controller.enabled=false", "Deployment", "csi-provisioner"},
		{"plugin.enabled=false", "DaemonSet", "csi-plugin"},
		{"substrate.storageClass.create=false", "StorageClass", "ate-storage"},
		{"substrate.createDriverConfig=false", "CSIDriverConfig", "substrate-csi-bridge"},
	} {
		t.Run(tc.option, func(t *testing.T) {
			for _, resource := range renderChart(t, "--set", "enableSubstrate=true", "--set", tc.option) {
				if resource.Kind == tc.forbiddenKind && resource.Metadata.Name == tc.forbiddenName {
					t.Fatalf("unexpected %s/%s", resource.Kind, resource.Metadata.Name)
				}
			}
		})
	}
}

func TestSubstrateWithoutNativeControllers(t *testing.T) {
	resources := renderChart(t, "--set", "enableSubstrate=true", "--set", "csi.nas.enabled=false,csi.disk.enabled=false,csi.oss.enabled=false,csi.bmcpfs.enabled=false,csi.ens.enabled=false")
	for _, resource := range resources {
		if resource.Kind == "CSIDriverConfig" && resource.Metadata.Name == "substrate-nas" {
			t.Fatal("disabled NAS must not have a controller config")
		}
	}
	var deployment appsv1.DeploymentSpec
	if err := json.Unmarshal(findResource(t, resources, "Deployment", "csi-provisioner").Spec, &deployment); err != nil {
		t.Fatal(err)
	}
	if names := driverNames(container(t, deployment.Template.Spec, "csi-provisioner")); !slices.Equal(names, []string{"substrate"}) {
		t.Fatalf("unexpected driver list %v", names)
	}
	var service corev1.ServiceSpec
	if err := json.Unmarshal(findResource(t, resources, "Service", "csi-provisioner").Spec, &service); err != nil {
		t.Fatal(err)
	}
	if len(service.Ports) != 1 || service.Ports[0].Port != 444 || service.Ports[0].TargetPort.IntVal != 8444 {
		t.Fatal("bridge service port must remain stable without NAS")
	}
}

func TestSubstrateControllerOnlyDoesNotRequireActorAPI(t *testing.T) {
	resources := renderChart(t, "--set", "enableSubstrate=true", "--set", "plugin.enabled=false", "--set", "substrate.apiEndpoint=", "--set", "substrate.apiAudience=")
	var deployment appsv1.DeploymentSpec
	if err := json.Unmarshal(findResource(t, resources, "Deployment", "csi-provisioner").Spec, &deployment); err != nil {
		t.Fatal(err)
	}
	container(t, deployment.Template.Spec, "substrate-controller-proxy")
	certificate, clientTrust := false, false
	for _, volume := range deployment.Template.Spec.Volumes {
		if volume.Projected == nil {
			continue
		}
		for _, source := range volume.Projected.Sources {
			if source.ServiceAccountToken != nil {
				t.Fatal("controller must not project an Actor API token")
			}
			if source.ClusterTrustBundle != nil {
				if source.ClusterTrustBundle.Path == "trust-bundle.pem" {
					t.Fatal("controller must not require the Actor API server trust bundle")
				}
				clientTrust = clientTrust || source.ClusterTrustBundle.Path == "client-trust-bundle.pem"
			}
			certificate = certificate || source.PodCertificate != nil
		}
	}
	if !certificate || !clientTrust {
		t.Fatal("controller must retain certificate and client trust for inbound mTLS")
	}
}

func TestSubstrateEnvoyUsesStandardImageSelection(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []string
		want string
	}{
		{"Beijing VPC", []string{"--set", "deploy.regionID=cn-beijing", "--set", "deploy.network=vpc"}, "registry-cn-beijing-vpc.ack.aliyuncs.com/acs/envoy:v1.39-latest"},
		{"Beijing public", []string{"--set", "deploy.regionID=cn-beijing", "--set", "deploy.network=internet"}, "registry-cn-beijing.ack.aliyuncs.com/acs/envoy:v1.39-latest"},
		{"custom image", []string{"--set", "deploy.network=vpc", "--set", "images.registryVPC=registry.example.com", "--set", "images.envoy.repo=team/envoy", "--set", "images.envoy.tag=reviewed"}, "registry.example.com/team/envoy:reviewed"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			args := append([]string{"--set", "enableSubstrate=true"}, tc.args...)
			resources := renderChart(t, args...)
			var deployment appsv1.DeploymentSpec
			if err := json.Unmarshal(findResource(t, resources, "Deployment", "csi-provisioner").Spec, &deployment); err != nil {
				t.Fatal(err)
			}
			if got := container(t, deployment.Template.Spec, "substrate-controller-proxy").Image; got != tc.want {
				t.Fatalf("Envoy image %q, want %q", got, tc.want)
			}
		})
	}
}
