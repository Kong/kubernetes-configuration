package v1alpha1

import "github.com/kong/kubernetes-configuration/api/configuration/common"

// NOTE: types below should be aliases but controller-tools' controller-gen
// panics in 0.17.1 when used with aliases:
// https://github.com/kubernetes-sigs/controller-tools/issues/1136

// ControlPlaneRef is the schema for the ControlPlaneRef type.
// It is used to reference a Control Plane entity.
type ControlPlaneRef common.ControlPlaneRef

const (
	// ControlPlaneRefKonnectNamespacedRef is the type for the KonnectNamespacedRef ControlPlaneRef.
	ControlPlaneRefKonnectNamespacedRef = common.ControlPlaneRefKonnectNamespacedRef
)
