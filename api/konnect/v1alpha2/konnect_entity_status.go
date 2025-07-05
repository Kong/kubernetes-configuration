package v1alpha2

// KonnectEntityStatus represents the status of a Konnect entity.
// +apireference:kgo:include
type KonnectEntityStatus struct {
	// ID is the unique identifier of the Konnect entity as assigned by Konnect API.
	// If it's unset (empty string), it means the Konnect entity hasn't been created yet.
	//
	// +optional
	ID string `json:"id,omitempty"`

	// ServerURL is the URL of the Konnect server in which the entity exists.
	//
	// +optional
	ServerURL string `json:"serverURL,omitempty"`

	// OrgID is ID of Konnect Org that this entity has been created in.
	//
	// +optional
	OrgID string `json:"organizationID,omitempty"`
}

// GetOrgID returns the OrgID field of the KonnectEntityStatus struct.
func (in *KonnectEntityStatus) GetOrgID() string {
	if in == nil {
		return ""
	}
	return in.OrgID
}

// SetOrgID sets the OrgID field of the KonnectEntityStatus struct.
func (in *KonnectEntityStatus) SetOrgID(id string) {
	in.OrgID = id
}

// GetKonnectID returns the ID field of the KonnectEntityStatus struct.
func (in *KonnectEntityStatus) GetKonnectID() string {
	if in == nil {
		return ""
	}
	return in.ID
}

// SetKonnectID sets the ID field of the KonnectEntityStatus struct.
func (in *KonnectEntityStatus) SetKonnectID(id string) {
	in.ID = id
}

// GetServerURL returns the server URL of the KonnectEntityStatus struct.
func (in *KonnectEntityStatus) GetServerURL() string {
	if in == nil {
		return ""
	}
	return in.ServerURL
}

// SetServerURL sets the server URL of the KonnectEntityStatus struct.
func (in *KonnectEntityStatus) SetServerURL(s string) {
	in.ServerURL = s
}
