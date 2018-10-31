package keycloak

// RoleRepresentation represents role
type RoleRepresentation struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type MappingRepresentation struct {
	RealmMappings  []RoleRepresentation `json:"realmMappings,omitempty"`
	ClientMappings bool                 `json:"clientMappings,omitempty"`
}
