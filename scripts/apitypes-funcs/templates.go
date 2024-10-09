package main

const (
	konnectFuncOutputFileName = "zz_generated_konnect_funcs.go"
	konnectFuncTemplate       = `package {{ .PackageVersion }}

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	konnectv1alpha1 "github.com/kong/kubernetes-configuration/api/konnect/v1alpha1"
)

// Code generated by scripts/apitypes-funcs/main.go; DO NOT EDIT.

{{- range .Types }}
{{ if .KonnectStatusType }}
func (obj *{{ .Type }}) initKonnectStatus() {
	obj.Status.Konnect = &{{ .KonnectStatusType }}{}
}

// GetKonnectStatus returns the Konnect status contained in the {{ .Type }} status.
func (obj *{{ .Type }}) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if obj.Status.Konnect == nil {
		return nil
	}
	return &obj.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the {{ .Type }} status.
func (obj *{{ .Type }}) GetKonnectID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the {{ .Type }} status.
func (obj *{{ .Type }}) SetKonnectID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ID = id
}

{{ end -}}
// GetTypeName returns the {{ .Type }} Kind name
func (obj {{ .Type }}) GetTypeName() string {
	return "{{ .Type }}"
}

// GetConditions returns the Status Conditions
func (obj *{{ .Type }}) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *{{ .Type }}) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

// GetControlPlaneID returns the ControlPlane ID in the {{ .Type }} status.
func (obj *{{ .Type }}) GetControlPlaneID() string {
	if obj.Status.Konnect == nil {
		return ""
	}
	return obj.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the {{ .Type }} status.
func (obj *{{ .Type }}) SetControlPlaneID(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ControlPlaneID = id
}
{{- if hasSuffix "KonnectEntityStatusWithControlPlaneAndConsumerRefs" .KonnectStatusType }}

func (obj *{{ .Type }}) SetKonnectConsumerIDInStatus(id string) {
	if obj.Status.Konnect == nil {
		obj.initKonnectStatus()
	}
	obj.Status.Konnect.ConsumerID = id
}

func (obj *{{ .Type }}) GetConsumerRefName() string {
	return obj.Spec.ConsumerRef.Name
}
{{- end }}

{{- end }}
`

	konnectFuncStandaloneTemplate = `package {{ .PackageVersion }}

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Code generated by scripts/apitypes-funcs/main.go; DO NOT EDIT.

{{- range .Types }}
{{ if .KonnectStatusType }}
// GetKonnectStatus returns the Konnect status contained in the {{ .Type }} status.
func (obj *{{ .Type }}) GetKonnectStatus() {{ .KonnectStatusType }} {
	return &obj.Status.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the {{ .Type }} status.
func (obj *{{ .Type }}) GetKonnectID() string {
	return obj.Status.ID
}

// SetKonnectID sets the Konnect ID in the {{ .Type }} status.
func (obj *{{ .Type }}) SetKonnectID(id string) {
	obj.Status.ID = id
}

{{- end }}
// GetTypeName returns the {{ .Type }} Kind name
func (obj {{ .Type }}) GetTypeName() string {
	return "{{ .Type }}"
}

// GetConditions returns the Status Conditions
func (obj *{{ .Type }}) GetConditions() []metav1.Condition {
	return obj.Status.Conditions
}

// SetConditions sets the Status Conditions
func (obj *{{ .Type }}) SetConditions(conditions []metav1.Condition) {
	obj.Status.Conditions = conditions
}

{{- end }}
`

	listFuncOutputFileNamme = "zz_generated_list_funcs.go"
	listFuncTemplate        = `package {{ .PackageVersion }}

// Code generated by scripts/apitypes-funcs/main.go; DO NOT EDIT.
{{- range .Types }}

// GetItems() returns the list of {{ .Type }} items.
func (obj {{ .Type }}List) GetItems() []{{ .Type }} {
	return obj.Items
}
{{- end }}
`
)
