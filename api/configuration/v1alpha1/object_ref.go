package v1alpha1

// TODO: change other types to use the generic `KongObjectRef` and move it to a common package to prevent possible import cycles.
//
// KongObjectRef is a reference to another object representing a Kong entity with deterministic type.
type KongObjectRef struct {
	// NOTE: the `Required` validation rule does not reject empty strings so we use `MinLength` to reject empty string here.
	//
	// Name is the name of the entity.
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`
	// TODO: handle cross namespace references.
}
