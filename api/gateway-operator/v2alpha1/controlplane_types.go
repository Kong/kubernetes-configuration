/*
Copyright 2025 Kong Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v2alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"

	commonv1alpha1 "github.com/kong/kubernetes-configuration/api/common/v1alpha1"
	operatorv1beta1 "github.com/kong/kubernetes-configuration/api/gateway-operator/v1beta1"
)

func init() {
	SchemeBuilder.Register(&ControlPlane{}, &ControlPlaneList{})
}

// ControlPlane is the Schema for the controlplanes API
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=kocp,categories=kong;all
// +kubebuilder:printcolumn:name="Ready",description="The Resource is ready",type=string,JSONPath=`.status.conditions[?(@.type=='Ready')].status`
// +apireference:kgo:include
// +kong:channels=gateway-operator
type ControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ControlPlaneSpec   `json:"spec,omitempty"`
	Status ControlPlaneStatus `json:"status,omitempty"`
}

// ControlPlaneList contains a list of ControlPlane
//
// +kubebuilder:object:root=true
// +apireference:kgo:include
type ControlPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ControlPlane `json:"items"`
}

// ControlPlaneSpec defines the desired state of ControlPlane
//
// +apireference:kgo:include
type ControlPlaneSpec struct {
	ControlPlaneOptions `json:",inline"`

	// GatewayClass indicates the Gateway resources which this ControlPlane
	// should be responsible for configuring routes for (e.g. HTTPRoute,
	// TCPRoute, UDPRoute, TLSRoute, etc.).
	//
	// Required for the ControlPlane to have any effect: at least one Gateway
	// must be present for configuration to be pushed to the data-plane and
	// only Gateway resources can be used to identify data-plane entities.
	//
	// +optional
	GatewayClass *gatewayv1.ObjectName `json:"gatewayClass,omitempty"`

	// IngressClass enables support for the older Ingress resource and indicates
	// which Ingress resources this ControlPlane should be responsible for.
	//
	// Routing configured this way will be applied to the Gateway resources
	// indicated by GatewayClass.
	//
	// If omitted, Ingress resources will not be supported by the ControlPlane.
	//
	// +optional
	IngressClass *string `json:"ingressClass,omitempty"`
}

// ControlPlaneOptions indicates the specific information needed to
// deploy and connect a ControlPlane to a DataPlane object.
//
// +apireference:kgo:include
// +kubebuilder:validation:XValidation:message="Extension not allowed for ControlPlane",rule="has(self.extensions) ? self.extensions.all(e, (e.group == 'konnect.konghq.com' && e.kind == 'KonnectExtension') || (e.group == 'gateway-operator.konghq.com' && e.kind == 'DataPlaneMetricsExtension')) : true"
type ControlPlaneOptions struct {
	// DataPlane designates the target data plane to configure.
	//
	// It can be either a URL to an externally managed DataPlane (e.g. installed
	// independently with Helm) or a name of a DataPlane resource that is
	// managed by the operator.
	//
	// +required
	DataPlane ControlPlaneDataPlaneTarget `json:"dataplane,omitempty"`

	// Extensions provide additional or replacement features for the ControlPlane
	// resources to influence or enhance functionality.
	//
	// +optional
	// +kubebuilder:validation:MinItems=0
	// +kubebuilder:validation:MaxItems=2
	Extensions []commonv1alpha1.ExtensionRef `json:"extensions,omitempty"`

	// WatchNamespaces indicates the namespaces to watch for resources.
	//
	// +optional
	// +kubebuilder:default={type: all}
	WatchNamespaces *operatorv1beta1.WatchNamespaces `json:"watchNamespaces,omitempty"`

	// FeatureGates is a list of feature gates that are enabled for this ControlPlane.
	//
	// +optional
	// +listType=map
	// +listMapKey=name
	// +kubebuilder:validation:MaxItems=32
	FeatureGates []ControlPlaneFeatureGate `json:"featureGates,omitempty"`

	// Controllers defines the controllers that are enabled for this ControlPlane.
	//
	// +optional
	// +listType=map
	// +listMapKey=name
	// +kubebuilder:validation:MaxItems=32
	Controllers []ControlPlaneController `json:"controllers,omitempty"`

	// AdminAPI defines the configuration for the Kong Admin API.
	//
	// +optional
	AdminAPI *ControlPlaneAdminAPI `json:"adminAPI,omitempty"`
}

// ControlPlaneDataPlaneTarget defines the target for the DataPlane that the ControlPlane
// is responsible for configuring.
//
// +kubebuilder:validation:XValidation:message="URL has to be provided when type is set to url",rule="self.type != 'url' || has(self.url)"
// +kubebuilder:validation:XValidation:message="Name cannot be provided when type is set to url",rule="self.type != 'url' || !has(self.name)"
// +kubebuilder:validation:XValidation:message="Name has to be provided when type is set to name",rule="self.type != 'name' || has(self.name)"
// +kubebuilder:validation:XValidation:message="URL cannot be provided when type is set to name",rule="self.type != 'name' || !has(self.url)"
type ControlPlaneDataPlaneTarget struct {
	// Type indicates the type of the DataPlane target.
	//
	// +kubebuilder:validation:Enum=url;name
	// +kubebuilder:validation:Required
	Type ControlPlaneDataPlaneTargetType `json:"type,omitempty"`

	// URL is the URL of the DataPlane target. This is used for configuring
	// externally managed DataPlanes like those installed independently with Helm.
	//
	// +optional
	// +kubebuilder:validation:Pattern=`^https?://[a-zA-Z0-9.-]+(:[0-9]+)?(/.*)?$`
	URL string `json:"url,omitempty"`

	// Name is the name of the DataPlane to configure.
	//
	// +optional
	Name string `json:"name,omitempty"`
}

// ControlPlaneDataPlaneTargetType defines the type of the DataPlane target
// that the ControlPlane is responsible for configuring.
type ControlPlaneDataPlaneTargetType string

const (
	// ControlPlaneDataPlaneTargetURL indicates that the DataPlane target is a URL.
	ControlPlaneDataPlaneTargetURL ControlPlaneDataPlaneTargetType = "url"

	// ControlPlaneDataPlaneTargetName indicates that the DataPlane target is a name
	// of a DataPlane resource managed by the operator.
	ControlPlaneDataPlaneTargetName ControlPlaneDataPlaneTargetType = "name"
)

// ControlPlaneAdminAPI defines the configuration for the Kong Admin API that
// a ControlPlane when configuring the DataPlane.
type ControlPlaneAdminAPI struct {
	// Workspace indicates the Kong Workspace to use for the ControlPlane.
	// If left empty then no Kong workspace will be used.
	//
	// +optional
	Workspace string `json:"workspace,omitempty"`
}

// ControlPlaneController defines a controller state for the ControlPlane.
// It overrides the default behavior as defined in the deployed operator version.
//
// +apireference:kgo:include
type ControlPlaneController struct {
	// Name is the name of the controller.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name,omitempty"`

	// Enabled indicates whether the controller is enabled or not.
	//
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled,omitempty"`
}

// ControlPlaneFeatureGate defines a feature gate state for the ControlPlane.
// It overrides the default behavior as defined in the deployed operator version.
//
// +apireference:kgo:include
type ControlPlaneFeatureGate struct {
	// Name is the name of the feature gate.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name,omitempty"`

	// Enabled indicates whether the feature gate is enabled or not.
	//
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled,omitempty"`
}

// ControlPlaneStatus defines the observed state of ControlPlane
//
// +apireference:kgo:include
type ControlPlaneStatus struct {
	// Conditions describe the current conditions of the Gateway.
	//
	// +optional
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=8
	// +kubebuilder:default={{type: "Scheduled", status: "Unknown", reason:"NotReconciled", message:"Waiting for controller", lastTransitionTime: "1970-01-01T00:00:00Z"}}
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// GetConditions returns the ControlPlane Status Conditions
func (c *ControlPlane) GetConditions() []metav1.Condition {
	return c.Status.Conditions
}

// SetConditions sets the ControlPlane Status Conditions
func (c *ControlPlane) SetConditions(conditions []metav1.Condition) {
	c.Status.Conditions = conditions
}

// GetExtensions retrieves the ControlPlane Extensions
func (c *ControlPlane) GetExtensions() []commonv1alpha1.ExtensionRef {
	return c.Spec.Extensions
}
