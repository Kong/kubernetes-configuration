package v1alpha1

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// KongCustomEntityKind is the kind name for the KongCustomEntity resource.
	KongCustomEntityKind = "KongCustomEntity"
)

// KongEntityScope defines the scope of the Kong entity.
type KongEntityScope string

// KongCustomEntity defines a "custom" Kong entity that KIC cannot support the entity type directly.
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:resource:shortName=kce,categories=kong-ingress-controller;kong,path=kongcustomentities,singular=kongcustomentity
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Entity Type",type=string,JSONPath=`.spec.type`,description="type of the Kong entity"
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`,description="Age"
// +kubebuilder:printcolumn:name="Programmed",type=string,JSONPath=`.status.conditions[?(@.type=="Programmed")].status`
// +kubebuilder:validation:XValidation:rule="self.spec.type == oldSelf.spec.type",message="The spec.type field is immutable"
// +apireference:kic:include
// +kong:channels=ingress-controller
type KongCustomEntity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KongCustomEntitySpec `json:"spec"`

	// Status stores the reconciling status of the resource.
	Status KongCustomEntityStatus `json:"status,omitempty"`
}

// KongCustomEntitySpec defines the specification of the KongCustomEntity.
// +kubebuilder:validation:XValidation:rule="!(self.type in ['services','routes','upstreams','targets','plugins','consumers','consumer_groups'])",message="The type field cannot be one of the known Kong entity types"
// +apireference:kic:include
type KongCustomEntitySpec struct {
	// EntityType is the type of the Kong entity. The type is used in generating declarative configuration.
	EntityType string `json:"type"`
	// Fields defines the fields of the Kong entity itself.
	Fields apiextensionsv1.JSON `json:"fields"`
	// ControllerName specifies the controller that should reconcile it, like ingress class.
	ControllerName string `json:"controllerName"`

	// ParentRef references the kubernetes resource it attached to when its scope is "attached".
	// Currently only KongPlugin/KongClusterPlugin allowed. This will make the custom entity to be attached
	// to the entity(service/route/consumer) where the plugin is attached.
	ParentRef *ObjectReference `json:"parentRef,omitempty"`
}

// ObjectReference defines reference of a kubernetes object.
// +apireference:kic:include
type ObjectReference struct {
	// Group defines the API group of the referred object.
	Group *string `json:"group,omitempty"`
	// Kind defines the kind of the referred object.
	// +kubebuilder:validation:Enum=KongPlugin;KongClusterPlugin
	Kind *string `json:"kind,omitempty"`
	// Empty namespace means the same namespace of the owning object.
	Namespace *string `json:"namespace,omitempty"`
	// Name defines the name of the referred object.
	Name string `json:"name"`
}

// KongCustomEntityStatus defines the status of the KongCustomEntity.
// +apireference:kic:include
type KongCustomEntityStatus struct {
	// Conditions describe the current conditions of the KongCustomEntityStatus.
	//
	// Known condition types are:
	//
	// * "Programmed"
	//
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=8
	// +kubebuilder:default={{type: "Programmed", status: "Unknown", reason:"Pending", message:"Waiting for controller", lastTransitionTime: "1970-01-01T00:00:00Z"}}
	Conditions []metav1.Condition `json:"conditions"`
}

// +kubebuilder:object:root=true

// KongCustomEntityList contains a list of KongCustomEntity.
type KongCustomEntityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KongCustomEntity `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KongCustomEntity{}, &KongCustomEntityList{})
}
