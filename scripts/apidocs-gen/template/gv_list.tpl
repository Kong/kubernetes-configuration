{{- define "gvList" -}}
{{- $groupVersions := . -}}

<!-- This document is generated by kong/kubernetes-configuration's 'generate.docs' make target, DO NOT EDIT -->

## Packages
{{- range $groupVersions }}
- {{ markdownRenderGVLink . }}
{{- end }}

{{ range $groupVersions }}
{{ template "gvDetails" . }}
{{ end }}

{{- end -}}
