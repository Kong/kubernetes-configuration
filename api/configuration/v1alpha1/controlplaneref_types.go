package v1alpha1

import "github.com/kong/kubernetes-configuration/api/configuration/common"

// NOTE: types below should be aliases but controller-tools' controller-gen
// panics in 0.17.1 when used with aliases:
// https://github.com/kubernetes-sigs/controller-tools/issues/1136

// ControlPlaneRef is the schema for the ControlPlaneRef type.
// It is used to reference a Control Plane entity.
type ControlPlaneRef common.ControlPlaneRef

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
