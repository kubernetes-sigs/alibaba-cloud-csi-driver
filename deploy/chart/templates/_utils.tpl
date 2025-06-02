{{- define "registry" -}}
{{- $vpc := (eq .deploy.network "vpc") -}}
{{- if (eq .deploy.network nil) -}}
    {{- $vpc = .deploy.ecs -}}
{{- end -}}
{{- if $vpc -}}
    {{- if (ne .images.registryVPC nil) -}}
        {{- .images.registryVPC -}}
    {{- else if (eq .deploy.regionID nil) -}}
        registry-cn-hangzhou.ack.aliyuncs.com
    {{- else -}}
        registry-{{ .deploy.regionID }}-vpc.aliyuncs.com
    {{- end -}}
{{- else -}}
    {{- if (ne .images.registry nil) -}}
        {{- .images.registry -}}
    {{- else if (eq .deploy.regionID nil) -}}
        registry-cn-hangzhou.ack.aliyuncs.com
    {{- else -}}
        registry-{{ .deploy.regionID }}.aliyuncs.com
    {{- end -}}
{{- end -}}
{{- end -}}


{{- define "imageSpec" -}}
{{- $v := index . 0 -}}
{{- $container := index . 1 -}}
{{- $cv := get $v.images $container -}}
{{ printf "%s/%s:%s" (include "registry" $v) $cv.repo $cv.tag | quote }}
{{- end -}}

{{- define "enabledPlugins" -}}
    {{- $drivers := list -}}
    {{- $csi := . -}}
    {{- range $key := tuple "disk" "nas" "oss" }}
        {{- if (index $csi $key).enabled -}}
            {{- $drivers = append $drivers $key -}}
        {{- end -}}
    {{- end -}}
    {{- $drivers | join "," -}}
{{- end -}}

{{- define "enabledControllers" -}}
    {{- $drivers := list -}}
    {{- $csi := . -}}
    {{- range $key := tuple "disk" "nas" "oss" "bmcpfs" }}
        {{- $val := index $csi $key -}}
        {{- if and $val.enabled $val.controller.enabled -}}
            {{- $drivers = append $drivers $key -}}
        {{- end -}}
    {{- end -}}
    {{- $drivers | join "," -}}
{{- end -}}

{{- define "akEnv" -}}
{{- if .enabled -}}
- name: ACCESS_KEY_ID
  valueFrom:
    secretKeyRef:
      name: {{ .secretName }}
      key: {{ default "id" .idKey }}
- name: ACCESS_KEY_SECRET
  valueFrom:
    secretKeyRef:
      name: {{ .secretName }}
      key: {{ default "secret" .secretKey }}
{{- end -}}
{{- end -}}

{{- define "kubeletDirEnv" -}}
{{- $d := clean . -}}
{{- if ne $d "/var/lib/kubelet" -}}
- name: KUBELET_ROOT_DIR
  value: {{ $d | quote }}
{{- end -}}
{{- end -}}

{{- define "networkEnv" -}}
{{- if (ne .network nil) -}}
- name: ALIBABA_CLOUD_NETWORK_TYPE
  value: {{ .network | quote }}
{{- else if .ecs -}}
- name: ALIBABA_CLOUD_NETWORK_TYPE
  value: "vpc"
{{- end -}}
{{- end -}}
