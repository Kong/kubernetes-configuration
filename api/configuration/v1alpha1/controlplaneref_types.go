package v1alpha1

import "github.com/kong/kubernetes-configuration/api/configuration/common"

// ControlPlaneRef is the schema for the ControlPlaneRef type.
// It is used to reference a Control Plane entity.
type ControlPlaneRef = common.ControlPlaneRef

// KonnectNamespacedRef is the schema for the KonnectNamespacedRef type.
type KonnectNamespacedRef = common.KonnectNamespacedRef

const (
	// ControlPlaneRefKonnectID is the type for the KonnectID ControlPlaneRef.
	// It is used to reference a Konnect Control Plane entity by its ID on the Konnect platform.
	ControlPlaneRefKonnectID = common.ControlPlaneRefKonnectID
	// ControlPlaneRefKonnectNamespacedRef is the type for the KonnectNamespacedRef ControlPlaneRef.
	// It is used to reference a Konnect Control Plane entity inside the cluster
	// using a namespaced reference.
	ControlPlaneRefKonnectNamespacedRef = common.ControlPlaneRefKonnectNamespacedRef
	// ControlPlaneRefKIC is the type for KIC ControlPlaneRef.
	// It is used to reference a KIC as Control Plane.
	ControlPlaneRefKIC = common.ControlPlaneRefKIC
)
