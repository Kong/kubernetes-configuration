package v1alpha1

// KonnectConfiguration is the Schema for the KonnectConfiguration API.
// +kubebuilder:object:generate=false
// +apireference:kgo:include
type KonnectConfiguration struct {
	// APIAuthConfigurationRef is the reference to the API Auth Configuration
	// that should be used for this Konnect Configuration.
	//
	// +kubebuilder:validation:Required
	APIAuthConfigurationRef KonnectAPIAuthConfigurationRef `json:"authRef"`

	// NOTE: Place for extending the KonnectConfiguration object.
	// This is a good place to add fields like "class" which could reference a cluster-wide
	// configuration for Konnect (similar to what Gateway API's GatewayClass).
}

// KonnectEntityOptions stores the options of entities specific to Konnect.
// +apireference:kgo:include
type KonnectEntityOptions struct {
	Adopt *KonnectAdoptOptions `json:"adopt,omitempty"`
}

// KonnectAdoptOptions stores the options for adopting existing Konnect entities.
// +apireference:kgo:include
type KonnectAdoptOptions struct {
	ID string `json:"id"`
}
