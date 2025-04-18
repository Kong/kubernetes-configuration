/*
Copyright 2021 Kong, Inc.

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

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	SchemeBuilder.Register(&UDPIngress{}, &UDPIngressList{})
}

// UDPIngressList contains a list of UDPIngress.
//
// +kubebuilder:object:root=true
// +apireference:kic:include
type UDPIngressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UDPIngress `json:"items"`
}

// UDPIngress is the Schema for the udpingresses API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=kong-ingress-controller;kong
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Address",type=string,JSONPath=`.status.loadBalancer.ingress[*].ip`,description="Address of the load balancer"
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`,description="Age"
// +apireference:kic:include
// +kong:channels=ingress-controller
type UDPIngress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the UDPIngress specification.
	Spec   UDPIngressSpec   `json:"spec,omitempty"`
	Status UDPIngressStatus `json:"status,omitempty"`
}

// UDPIngressSpec defines the desired state of UDPIngress.
// +apireference:kic:include
type UDPIngressSpec struct {
	// A list of rules used to configure the Ingress.
	Rules []UDPIngressRule `json:"rules,omitempty"`
}

// UDPIngressStatus defines the observed state of UDPIngress.
// +apireference:kic:include
type UDPIngressStatus struct {
	// LoadBalancer contains the current status of the load-balancer.
	LoadBalancer corev1.LoadBalancerStatus `json:"loadBalancer,omitempty"`
}
