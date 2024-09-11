/*
Copyright 2024 Kong, Inc.

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

package v1alpha1

import (
	sdkkonnectgocomp "github.com/Kong/sdk-konnect-go/models/components"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	konnectv1alpha1 "github.com/kong/kubernetes-configuration/api/konnect/v1alpha1"
)

// KongUpstream is the schema for Upstream API which defines a Kong Upstream.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Programmed",description="The Resource is Programmed on Konnect",type=string,JSONPath=`.status.conditions[?(@.type=='Programmed')].status`
// +kubebuilder:validation:XValidation:rule="!has(oldSelf.spec.controlPlaneRef) || has(self.spec.controlPlaneRef)", message="controlPlaneRef is required once set"
// +kubebuilder:validation:XValidation:rule="(!self.status.conditions.exists(c, c.type == 'Programmed' && c.status == 'True')) ? true : oldSelf.spec.controlPlaneRef == self.spec.controlPlaneRef", message="spec.controlPlaneRef is immutable when an entity is already Programmed"
type KongUpstream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KongUpstreamSpec `json:"spec"`

	// +kubebuilder:default={conditions: {{type: "Programmed", status: "Unknown", reason:"Pending", message:"Waiting for controller", lastTransitionTime: "1970-01-01T00:00:00Z"}}}
	Status KongUpstreamStatus `json:"status,omitempty"`
}

func (s *KongUpstream) initKonnectStatus() {
	s.Status.Konnect = &konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef{}
}

// GetKonnectStatus returns the Konnect status contained in the KongUpstream status.
func (s *KongUpstream) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
	if s.Status.Konnect == nil {
		return nil
	}
	return &s.Status.Konnect.KonnectEntityStatus
}

// GetKonnectID returns the Konnect ID in the KongUpstream status.
func (s *KongUpstream) GetKonnectID() string {
	if s.Status.Konnect == nil {
		return ""
	}
	return s.Status.Konnect.ID
}

// SetKonnectID sets the Konnect ID in the KongUpstream status.
func (s *KongUpstream) SetKonnectID(id string) {
	if s.Status.Konnect == nil {
		s.initKonnectStatus()
	}
	s.Status.Konnect.ID = id
}

// GetControlPlaneID returns the ControlPlane ID in the KongUpstream status.
func (s *KongUpstream) GetControlPlaneID() string {
	if s.Status.Konnect == nil {
		return ""
	}
	return s.Status.Konnect.ControlPlaneID
}

// SetControlPlaneID sets the ControlPlane ID in the KongUpstream status.
func (s *KongUpstream) SetControlPlaneID(id string) {
	if s.Status.Konnect == nil {
		s.initKonnectStatus()
	}
	s.Status.Konnect.ControlPlaneID = id
}

// GetTypeName returns the KongUpstream Kind name
func (s KongUpstream) GetTypeName() string {
	return "KongUpstream"
}

// GetConditions returns the Status Conditions
func (s *KongUpstream) GetConditions() []metav1.Condition {
	return s.Status.Conditions
}

// SetConditions sets the Status Conditions
func (s *KongUpstream) SetConditions(conditions []metav1.Condition) {
	s.Status.Conditions = conditions
}

// KongUpstreamSpec defines specification of a Kong Route.
type KongUpstreamSpec struct {
	// ControlPlaneRef is a reference to a ControlPlane this KongUpstream is associated with.
	// +optional
	ControlPlaneRef *ControlPlaneRef `json:"controlPlaneRef,omitempty"`

	KongUpstreamAPISpec `json:",inline"`
}

// KongUpstreamAPISpec defines specification of a Kong Service.
type KongUpstreamAPISpec struct {
	// Which load balancing algorithm to use.
	Algorithm *sdkkonnectgocomp.UpstreamAlgorithm `default:"round-robin" json:"algorithm"`
	// If set, the certificate to be used as client certificate while TLS handshaking to the upstream server.
	ClientCertificate *sdkkonnectgocomp.UpstreamClientCertificate `json:"client_certificate,omitempty"`
	// What to use as hashing input if the primary `hash_on` does not return a hash (eg. header is missing, or no Consumer identified). Not available if `hash_on` is set to `cookie`.
	HashFallback *sdkkonnectgocomp.HashFallback `default:"none" json:"hash_fallback"`
	// The header name to take the value from as hash input. Only required when `hash_fallback` is set to `header`.
	HashFallbackHeader *string `json:"hash_fallback_header,omitempty"`
	// The name of the query string argument to take the value from as hash input. Only required when `hash_fallback` is set to `query_arg`.
	HashFallbackQueryArg *string `json:"hash_fallback_query_arg,omitempty"`
	// The name of the route URI capture to take the value from as hash input. Only required when `hash_fallback` is set to `uri_capture`.
	HashFallbackURICapture *string `json:"hash_fallback_uri_capture,omitempty"`
	// What to use as hashing input. Using `none` results in a weighted-round-robin scheme with no hashing.
	HashOn *sdkkonnectgocomp.HashOn `default:"none" json:"hash_on"`
	// The cookie name to take the value from as hash input. Only required when `hash_on` or `hash_fallback` is set to `cookie`. If the specified cookie is not in the request, Kong will generate a value and set the cookie in the response.
	HashOnCookie *string `json:"hash_on_cookie,omitempty"`
	// The cookie path to set in the response headers. Only required when `hash_on` or `hash_fallback` is set to `cookie`.
	HashOnCookiePath *string `default:"/" json:"hash_on_cookie_path"`
	// The header name to take the value from as hash input. Only required when `hash_on` is set to `header`.
	HashOnHeader *string `json:"hash_on_header,omitempty"`
	// The name of the query string argument to take the value from as hash input. Only required when `hash_on` is set to `query_arg`.
	HashOnQueryArg *string `json:"hash_on_query_arg,omitempty"`
	// The name of the route URI capture to take the value from as hash input. Only required when `hash_on` is set to `uri_capture`.
	HashOnURICapture *string                        `json:"hash_on_uri_capture,omitempty"`
	Healthchecks     *sdkkonnectgocomp.Healthchecks `json:"healthchecks,omitempty"`
	// The hostname to be used as `Host` header when proxying requests through Kong.
	HostHeader *string `json:"host_header,omitempty"`
	// This is a hostname, which must be equal to the `host` of a Service.
	Name *string `json:"name,omitempty"`
	// The number of slots in the load balancer algorithm. If `algorithm` is set to `round-robin`, this setting determines the maximum number of slots. If `algorithm` is set to `consistent-hashing`, this setting determines the actual number of slots in the algorithm. Accepts an integer in the range `10`-`65536`.
	Slots *int64 `default:"10000" json:"slots"`
	// An optional set of strings associated with the Upstream for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the balancer will use SRV hostname(if DNS Answer has SRV record) as the proxy upstream `Host`.
	UseSrvName *bool `default:"false" json:"use_srv_name"`
}

// KongUpstreamStatus represents the current status of the Kong Service resource.
type KongUpstreamStatus struct {
	// Konnect contains the Konnect entity status.
	// +optional
	Konnect *konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef `json:"konnect,omitempty"`

	// Conditions describe the status of the Konnect entity.
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=8
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true

// KongUpstreamList contains a list of Kong Services.
type KongUpstreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KongUpstream `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KongUpstream{}, &KongUpstreamList{})
}
