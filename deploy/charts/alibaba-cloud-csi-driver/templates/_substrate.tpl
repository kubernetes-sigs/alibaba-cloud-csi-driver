{{- define "substrate.drivers" -}}
{{- .base -}}{{- if .enabled -}}{{- if .base }},{{ end -}}substrate{{- end -}}
{{- end -}}

{{- define "substrate.commonVolumes" -}}
- name: substrate-identity
  projected:
    sources:
      - podCertificate:
          signerName: podidentity.podcert.ate.dev/identity
          keyType: ECDSAP256
          credentialBundlePath: credential-bundle.pem
{{- if .node }}
      - clusterTrustBundle:
          signerName: servicedns.podcert.ate.dev/identity
          labelSelector:
            matchLabels: {podcert.ate.dev/canarying: live}
          path: trust-bundle.pem
{{- end }}
      - clusterTrustBundle:
          signerName: podidentity.podcert.ate.dev/identity
          labelSelector:
            matchLabels: {podcert.ate.dev/canarying: live}
          path: client-trust-bundle.pem
{{- if .node }}
- name: substrate-api-token
  projected:
    sources:
      - serviceAccountToken:
          audience: {{ required "substrate.apiAudience is required" .Values.substrate.apiAudience | quote }}
          expirationSeconds: 3600
          path: token
{{- end }}
{{- end -}}

{{- define "substrate.commonMounts" -}}
- name: substrate-identity
  mountPath: /run/podidentity.podcert.ate.dev
  readOnly: true
{{- if .node }}
- name: substrate-api-token
  mountPath: /run/ateapi
  readOnly: true
{{- end }}
{{- end -}}

{{- define "substrate.apiArgs" -}}
- --substrate-api-endpoint={{ required "substrate.apiEndpoint is required" .Values.substrate.apiEndpoint }}
- --substrate-api-ca-file=/run/podidentity.podcert.ate.dev/trust-bundle.pem
- --substrate-api-token-file=/run/ateapi/token
{{- end -}}

{{- define "substrate.validate" -}}
{{- if or (not (hasPrefix "/" .Values.substrate.actorRoot)) (eq (clean .Values.substrate.actorRoot) "/") -}}
{{- fail "substrate.actorRoot must be an absolute dedicated directory, not /" -}}
{{- end -}}
{{- end -}}
