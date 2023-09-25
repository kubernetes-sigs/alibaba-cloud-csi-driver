{{- define "imageSpec" -}}
{{- $repo := index . 0 -}}
{{- $i := index . 1 -}}
{{ $repo }}/{{ $i.repo }}:{{ $i.tag }}
{{- end -}}

{{- define "enabledDrivers" -}}
    {{- $drivers := list -}}
    {{- range $key, $val := . }}
        {{- if $val.enabled -}}
        {{- $drivers = append $drivers $key -}}
        {{- end -}}
    {{- end -}}
    {{- $drivers | join "," -}}
{{- end -}}
