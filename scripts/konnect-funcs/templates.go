package main

const (
	konnectFuncOutputFileName = "zz_generated_konnect_funcs.go"
	konnectFuncTemplate       = `package {{ .Package }}

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	konnectv1alpha1 "github.com/kong/kubernetes-configuration/api/konnect/v1alpha1"
)

// Code generated by scripts/konnect-funcs/main.go; DO NOT EDIT.

{{ range .Types }}
func ({{ .ReceiverName }} *{{ .Type }}) initKonnectStatus() {
	{{ .ReceiverName }}.Status.Konnect = &{{ .KonnectStatusType }}{}
}

// GetKonnectStatus returns the Konnect status contained in the {{ .Type }} status.
func ({{ .ReceiverName }} *{{ .Type }}) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if {{ .ReceiverName }}.Status.Konnect == nil {
		return nil
	}
	return &{{ .ReceiverName }}.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the {{ .Type }} status.
func ({{ .ReceiverName }} *{{ .Type }}) GetKonnectID() string {
	if {{ .ReceiverName }}.Status.Konnect == nil {
		return ""
	}
	return {{ .ReceiverName }}.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the {{ .Type }} status.
func ({{ .ReceiverName }} *{{ .Type }}) SetKonnectID(id string) {
	if {{ .ReceiverName }}.Status.Konnect == nil {
		{{ .ReceiverName }}.initKonnectStatus()
	}
	{{ .ReceiverName }}.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the {{ .Type }} status.
func ({{ .ReceiverName }} *{{ .Type }}) GetControlPlaneID() string {
	if {{ .ReceiverName }}.Status.Konnect == nil {
		return ""
	}
	return {{ .ReceiverName }}.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the {{ .Type }} status.
func ({{ .ReceiverName }} *{{ .Type }}) SetControlPlaneID(id string) {
	if {{ .ReceiverName }}.Status.Konnect == nil {
		{{ .ReceiverName }}.initKonnectStatus()
	}
	{{ .ReceiverName }}.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the {{ .Type }} Kind name
func ({{ .ReceiverName }} {{ .Type }}) GetTypeName() string {
	return "{{ .Type }}"
}

// GetConditions returns the Status Conditions
func ({{ .ReceiverName }} *{{ .Type }}) GetConditions() []metav1.Condition {
	return {{ .ReceiverName }}.Status.Conditions
}

// SetConditions sets the Status Conditions
func ({{ .ReceiverName }} *{{ .Type }}) SetConditions(conditions []metav1.Condition) {
	{{ .ReceiverName }}.Status.Conditions = conditions
}

{{- end }}
`
)
