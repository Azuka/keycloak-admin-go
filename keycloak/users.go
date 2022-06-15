//go:generate gomodifytags -file $GOFILE -struct UserRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct UserConsentRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct CredentialRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct FederatedIdentityRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct UserSessionRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct GroupRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase

package keycloak

// UserConsentRepresentation represents client consents
type UserConsentRepresentation struct {
	ClientID            string    `json:"clientId,omitempty"`
	CreatedDate         *UnixTime `json:"createdDate,omitempty"`
	GrantedClientScopes []string  `json:"grantedClientScopes,omitempty"`
	LastUpdatedDate     *UnixTime `json:"lastUpdatedDate,omitempty"`
}

// CredentialRepresentation represents credentials for a user or client
type CredentialRepresentation struct {
	Algorithm         string    `json:"algorithm,omitempty"`
	Counter           int32     `json:"counter,omitempty"`
	CreatedDate       *UnixTime `json:"createdDate,omitempty"`
	Device            string    `json:"device,omitempty"`
	Digits            int32     `json:"digits,omitempty"`
	HashIterations    int32     `json:"hashIterations,omitempty"`
	HashedSaltedValue string    `json:"hashedSaltedValue,omitempty"`
	Period            int32     `json:"period,omitempty"`
	Salt              string    `json:"salt,omitempty"`
	Temporary         *bool     `json:"temporary,omitempty"`
	Type              string    `json:"type,omitempty"`
	Value             string    `json:"value,omitempty"`
}

// FederatedIdentityRepresentation represents a federated identity
type FederatedIdentityRepresentation struct {
	IdentityProvider string `json:"identityProvider,omitempty"`
	UserID           string `json:"userId,omitempty"`
	UserName         string `json:"userName,omitempty"`
}

// UserRepresentation represents a realm user in Keycloak
type UserRepresentation struct {
	Access                 AttributeMap                      `json:"access,omitempty"`
	Attributes             AttributeMap                      `json:"attributes,omitempty"`
	ClientRoles            AttributeMap                      `json:"clientRoles,omitempty"`
	ClientConsents         []UserConsentRepresentation       `json:"clientConsents,omitempty"`
	CreatedTimestamp       *UnixTime                         `json:"createdTimestamp,omitempty"`
	Credentials            []CredentialRepresentation        `json:"credentials,omitempty"`
	DisableCredentialTypes []string                          `json:"disableCredentialTypes,omitempty"`
	Email                  string                            `json:"email,omitempty"`
	EmailVerified          *bool                             `json:"emailVerified,omitempty"`
	Enabled                *bool                             `json:"enabled,omitempty"`
	FederatedIdentities    []FederatedIdentityRepresentation `json:"federatedIdentities,omitempty"`
	FederationLink         string                            `json:"federationLink,omitempty"`
	FirstName              string                            `json:"firstName,omitempty"`
	Groups                 []string                          `json:"groups,omitempty"`
	ID                     string                            `json:"id,omitempty"`
	LastName               string                            `json:"lastName,omitempty"`
	NotBefore              *UnixTime                         `json:"notBefore,omitempty"`
	Origin                 string                            `json:"origin,omitempty"`
	RealmRoles             []string                          `json:"realmRoles,omitempty"`
	RequiredActions        []string                          `json:"requiredActions,omitempty"`
	Self                   string                            `json:"self,omitempty"`
	ServiceAccountClientID string                            `json:"serviceAccountClientId,omitempty"`
	Username               string                            `json:"username,omitempty"`
}

// UserSessionRepresentation is a single session for a user
type UserSessionRepresentation struct {
	Clients    AttributeMap `json:"clients,omitempty"`
	ID         string       `json:"id,omitempty"`
	IPAddress  string       `json:"ipAddress,omitempty"`
	LastAccess *UnixTime    `json:"lastAccess,omitempty"`
	Start      *UnixTime    `json:"start,omitempty"`
	UserID     string       `json:"userID,omitempty"`
	UserName   string       `json:"userName,omitempty"`
}

// GroupRepresentation represents a single user group in a realm
type GroupRepresentation struct {
	Access      AttributeMap          `json:"access,omitempty"`
	Attributes  AttributeMap          `json:"attributes,omitempty"`
	ClientRoles AttributeMap          `json:"clientRoles,omitempty"`
	ID          string                `json:"id,omitempty"`
	Name        string                `json:"name,omitempty"`
	Path        string                `json:"path,omitempty"`
	RealmRoles  []string              `json:"realmRoles,omitempty"`
	SubGroups   []GroupRepresentation `json:"subGroups,omitempty"`
}
