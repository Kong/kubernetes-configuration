package v1alpha1

// EntityType is the type for all the entity types.
type EntityType string

const (
	// EntityTypeOrigin is the type for Origin entities.
	EntityTypeOrigin EntityType = "Origin"
	// EntityTypeMirror is the type for Mirror entities.
	EntityTypeMirror EntityType = "Mirror"
)
