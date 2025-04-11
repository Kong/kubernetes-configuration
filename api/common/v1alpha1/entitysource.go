package v1alpha1

// EntitySource is the type for all the entity types.
type EntitySource string

const (
	// EntityTypeOrigin is the type for Origin entities.
	EntityTypeOrigin EntitySource = "Origin"
	// EntityTypeMirror is the type for Mirror entities.
	EntityTypeMirror EntitySource = "Mirror"
)
