{{- define "imageSpec" -}}
{{- $v := index . 0 -}}
{{- $container := index . 1 -}}
{{- $cv := get $v $container -}}
{{ printf "%s/%s:%s" ($cv.registry | default $v.registry) $cv.repo $cv.tag | quote }}
{{- end -}}

{{- define "workerImageSpec" -}}
{{- $v := index . 0 -}}
{{- $container := index . 1 -}}
{{- $cv := get $v $container -}}
{{ printf "%s/%s:%s" ($cv.registry | default $v.workerRegistry | default $v.registry) $cv.repo $cv.tag | quote }}
{{- end -}}

{{- define "enabledPlugins" -}}
    {{- $drivers := list -}}
    {{- range $key, $val := . }}
        {{- if $val.enabled -}}
        {{- $drivers = append $drivers $key -}}
        {{- end -}}
    {{- end -}}
    {{- $drivers | join "," -}}
{{- end -}}

{{- define "enabledControllers" -}}
    {{- $drivers := list -}}
    {{- range $key, $val := . }}
        {{- if and $val.enabled (ne $val.controller.enabled false) -}}
        {{- $drivers = append $drivers $key -}}
        {{- end -}}
    {{- end -}}
    {{- $drivers | join "," -}}
{{- end -}}
