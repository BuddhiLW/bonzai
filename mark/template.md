# Name

{{ summary . }}

# Usage

{{ usage . | indent 4 }}

{{if .Cmds -}}
# Commands

{{ commands . | indent 4 }}

{{ end -}}
{{- if .Long -}}
# Description

{{ long . }}

{{ end }}
