package v1alpha1

const (
	// ControlPlaneRefKonnectID is the type for the KonnectID ControlPlaneRef.
	// It is used to reference a Konnect Control Plane entity by its ID on the Konnect platform.
	ControlPlaneRefKonnectID = "konnectID"
	// ControlPlaneRefKonnectNamespacedRef is the type for the KonnectNamespacedRef ControlPlaneRef.
	// It is used to reference a Konnect Control Plane entity inside the cluster
	// using a namespaced reference.
	ControlPlaneRefKonnectNamespacedRef = "konnectNamespacedRef"
)

// ControlPlaneRef is the schema for the ControlPlaneRef type.
// It is used to reference a Control Plane entity.
// +kubebuilder:validation:XValidation:rule="self.type == 'konnectNamespacedRef' ? has(self.konnectNamespacedRef) : true", message="when type is konnectNamespacedRef, konnectNamespacedRef must be set"
type ControlPlaneRef struct {
	// Type can be one of:
	// - konnectID
	// - konnectNamespacedRef
	// +kubebuilder:validation:Enum=konnectID;konnectNamespacedRef
	Type string `json:"type,omitempty"`

	// KonnectID is the schema for the KonnectID type.
	// This field is required when the Type is konnectID.
	// +optional
	KonnectID *string `json:"konnectID,omitempty"`

	// KonnectNamespacedRef is a reference to a Konnect Control Plane entity inside the cluster.
	// It contains the name of the Konnect Control Plane.
	// This field is required when the Type is konnectNamespacedRef.
	KonnectNamespacedRef *KonnectNamespacedRef `json:"konnectNamespacedRef,omitempty"`
	// +optional
}

// KonnectNamespacedRef is the schema for the KonnectNamespacedRef type.
type KonnectNamespacedRef struct {
	// Name is the name of the Konnect Control Plane.
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// TODO: Implement cross namespace references:
	// https://github.com/Kong/kubernetes-configuration/issues/36
}
