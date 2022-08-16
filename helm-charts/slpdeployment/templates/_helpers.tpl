{{/*
Expand the name of the operator.
*/}}
{{- define "slp.operatorName" -}}
opa-slp-operator
{{- end }}

{{/*
Expand the name of the chart.
*/}}
{{- define "slp.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "slp.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "slp.labels" -}}
{{ include "slp.selectorLabels" . }}
app.kubernetes.io/managed-by: {{ include "slp.operatorName" . }}-{{ .Chart.Version }}
app.kubernetes.io/name: {{ .Values.namespace }}-{{ include "slp.name" . }}
app.kubernetes.io/version: {{ .Values.image.tag | default .Chart.AppVersion }}
helm.sh/chart: {{ include "slp.chart" . }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "slp.selectorLabels" -}}
app: {{ .Values.namespace }}-{{ include "slp.name" . }}
version: {{ .Values.image.tag | default .Chart.AppVersion }}
system-type: istio
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "slp.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "slp.name" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}
