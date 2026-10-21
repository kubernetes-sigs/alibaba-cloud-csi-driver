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

	for _, flag := range []string{"--substrate-api-endpoint=k8s:///ate-system/api:443", "--substrate-api-server-name=api.ate-system.svc", "--substrate-api-ca-file=/run/podidentity.podcert.ate.dev/trust-bundle.pem", "--substrate-api-token-file=/run/ateapi/token"} {
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

	findResource(t, resources, "CSIDriverConfig", "substrate.csi.alibabacloud.com")
	findResource(t, resources, "CSIDriverConfig", "nasplugin.csi.alibabacloud.com")
	var service corev1.ServiceSpec
	if err := json.Unmarshal(findResource(t, resources, "Service", "csi-provisioner").Spec, &service); err != nil {
		t.Fatal(err)
	}
	if service.Selector["app"] != "csi-provisioner" {
		t.Fatal("NAS and bridge must share the existing controller service")
	}
	// Named ports rather than a bare count: a count silently accepts one driver's
	// endpoint going missing while another is duplicated.
	gotPorts := map[string]int32{}
	for _, port := range service.Ports {
		gotPorts[port.Name] = port.Port
	}
	for name, want := range map[string]int32{"nas-grpc-tls": 443, "bridge-grpc-tls": 444} {
		if gotPorts[name] != want {
			t.Errorf("controller service port %s = %d, want %d (all ports: %v)", name, gotPorts[name], want, gotPorts)
		}
	}
	config := findResource(t, resources, "ConfigMap", "csi-provisioner-substrate")
	for _, socket := range []string{"/csi/nasplugin.csi.alibabacloud.com/csi.sock", "/csi/substrate.csi.alibabacloud.com/csi.sock"} {
		if !strings.Contains(config.Data["envoy.yaml"], socket) {
			t.Fatalf("missing backend %s", socket)
		}
	}
	// Parsed rather than counted. A bare count of "require_client_certificate: true"
	// silently accepts one listener gaining mTLS while another loses it, and had to be
	// bumped by hand every time an endpoint was added.
	var bootstrap struct {
		Node struct {
			ID      string `json:"id"`
			Cluster string `json:"cluster"`
		} `json:"node"`
		StaticResources struct {
			Listeners []struct {
				Name         string `json:"name"`
				FilterChains []struct {
					TransportSocket struct {
						TypedConfig struct {
							RequireClientCertificate bool           `json:"require_client_certificate"`
							CommonTLSContext         map[string]any `json:"common_tls_context"`
						} `json:"typed_config"`
					} `json:"transport_socket"`
				} `json:"filter_chains"`
			} `json:"listeners"`
		} `json:"static_resources"`
	}
	if err := yamlutil.Unmarshal([]byte(config.Data["envoy.yaml"]), &bootstrap); err != nil {
		t.Fatalf("parsing envoy.yaml: %v", err)
	}
	if len(bootstrap.StaticResources.Listeners) == 0 {
		t.Fatal("envoy.yaml declares no listeners")
	}
	// SDS will not initialize without a node identity: Envoy exits at startup with
	// "TlsCertificateSdsApi: node 'id' and 'cluster' are required". A bootstrap that
	// only ever used inline certificates needs no node block, so this is easy to drop
	// when moving to SDS and only shows up as a CrashLoopBackOff sidecar.
	if bootstrap.Node.ID == "" || bootstrap.Node.Cluster == "" {
		t.Errorf("envoy.yaml node is %+v; SDS requires both id and cluster", bootstrap.Node)
	}
	for _, listener := range bootstrap.StaticResources.Listeners {
		if len(listener.FilterChains) != 1 {
			t.Fatalf("listener %s has %d filter chains, want exactly 1", listener.Name, len(listener.FilterChains))
		}
		tlsContext := listener.FilterChains[0].TransportSocket.TypedConfig
		if !tlsContext.RequireClientCertificate {
			t.Errorf("listener %s does not require a client certificate", listener.Name)
		}
		// Both the serving certificate and the client CA must arrive through SDS.
		// Envoy silently ignores watched_directory on an inline tls_certificates or
		// validation_context in the bootstrap, so an inline entry is read once at
		// startup and never again. kubelet rotates the projected podCertificate
		// roughly daily; from the moment the startup leaf expires every mTLS
		// handshake fails with "x509: certificate has expired" while the Pod keeps
		// reporting Ready, so nothing detects it and the driver stays unreachable
		// until it is restarted.
		for _, inline := range []string{"tls_certificates", "validation_context"} {
			if _, found := tlsContext.CommonTLSContext[inline]; found {
				t.Errorf("listener %s carries an inline %s entry; certificate rotation would be silently broken, deliver it through SDS instead", listener.Name, inline)
			}
		}
		for _, sds := range []string{"tls_certificate_sds_secret_configs", "validation_context_sds_secret_config"} {
			if _, found := tlsContext.CommonTLSContext[sds]; !found {
				t.Errorf("listener %s is missing %s", listener.Name, sds)
			}
		}
	}
	// The SDS resources are where watched_directory actually takes effect, and each
	// one has to be reachable as a file: the ConfigMap is mounted whole at /etc/envoy
	// with no items filter, so a key and a path_config_source that disagrees with it
	// is a startup failure rather than a silent one.
	sdsKeys := 0
	for key, body := range config.Data {
		if !strings.HasPrefix(key, "sds-") {
			continue
		}
		sdsKeys++
		if !strings.Contains(body, "watched_directory") {
			t.Errorf("SDS resource %s is missing watched_directory; rotation would be silently broken", key)
		}
		if !strings.Contains(config.Data["envoy.yaml"], "/etc/envoy/"+key) {
			t.Errorf("SDS resource %s is never referenced as /etc/envoy/%s", key, key)
		}
	}
	if sdsKeys != 2 {
		t.Errorf("found %d SDS resources, want exactly the serving certificate and the client CA", sdsKeys)
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
	config := findResource(t, resources, "CSIDriverConfig", "substrate.csi.alibabacloud.com")
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
		{"substrate.createDriverConfig=false", "CSIDriverConfig", "substrate.csi.alibabacloud.com"},
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
		if resource.Kind == "CSIDriverConfig" && resource.Metadata.Name == "nasplugin.csi.alibabacloud.com" {
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
