//go:generate gomodifytags -file $GOFILE -struct ClientRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct ResourceServerRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct PolicyRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct ScopeRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct ResourceRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct ProtocolMapperRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase

package keycloak

const (
	// PolicyEnforcementModeEnforcing marks policy enforcement as enforcing
	PolicyEnforcementModeEnforcing = "ENFORCING"
	// PolicyEnforcementModePermissive marks policy enforcement as permissive
	PolicyEnforcementModePermissive = "PERMISSIVE"
	// PolicyEnforcementModeDisabled marks policy enforcement as disabled
	PolicyEnforcementModeDisabled = "DISABLED"

	// DecisionstrategyAffirmative sets decision strategy to affirmative
	DecisionstrategyAffirmative = "AFFIRMATIVE"
	// DecisionstrategyUnanimous sets decision strategy to unanimous
	DecisionstrategyUnanimous = "UNANIMOUS"
	// DecisionstrategyConsensus sets decision strategy to consensus
	DecisionstrategyConsensus = "CONSENSUS"
)

// ClientRepresentation represents a client's configuration in a realm
type ClientRepresentation struct {
	Access                             string                         `json:"access,omitempty"`
	AdminURL                           string                         `json:"adminUrl,omitempty"`
	Attributes                         AttributeMap                   `json:"attributes,omitempty"`
	AuthenticationFlowBindingOverrides AttributeMap                   `json:"authenticationFlowBindingOverrides,omitempty"`
	AuthorizationServicesEnabled       *bool                          `json:"authorizationServicesEnabled,omitempty"`
	AuthorizationSettings              *ResourceServerRepresentation  `json:"authorizationSettings,omitempty"`
	BaseURL                            string                         `json:"baseURL,omitempty"`
	BearerOnly                         *bool                          `json:"bearerOnly,omitempty"`
	ClientAuthenticatorType            string                         `json:"clientAuthenticatorType,omitempty"`
	ClientID                           string                         `json:"clientID,omitempty"`
	ConsentRequired                    *bool                          `json:"consentRequired,omitempty"`
	DefaultClientScopes                []string                       `json:"defaultClientScopes,omitempty"`
	DefaultRoles                       []string                       `json:"defaultRoles,omitempty"`
	Description                        string                         `json:"description,omitempty"`
	DirectAccessGrantsEnabled          *bool                          `json:"directAccessGrantsEnabled,omitempty"`
	Enabled                            *bool                          `json:"enabled,omitempty"`
	FrontChannelLogout                 *bool                          `json:"frontChannelLogout,omitempty"`
	FullScopeAllowed                   *bool                          `json:"fullScopeAllowed,omitempty"`
	ID                                 string                         `json:"id,omitempty"`
	ImplicitFlowEnabled                *bool                          `json:"implicitFlowEnabled,omitempty"`
	Name                               string                         `json:"name,omitempty"`
	NodeRegistrationTimeout            *UnixTime                      `json:"nodeRegistrationTimeout,omitempty"`
	NotBefore                          *UnixTime                      `json:"notBefore,omitempty"`
	OptionalClientScopes               []string                       `json:"optionalClientScopes,omitempty"`
	Origin                             string                         `json:"origin,omitempty"`
	Protocol                           string                         `json:"protocol,omitempty"`
	ProtocolMappers                    []ProtocolMapperRepresentation `json:"protocolMappers,omitempty"`
	PublicClient                       *bool                          `json:"publicClient,omitempty"`
	RedirectURIs                       []string                       `json:"redirectURIs,omitempty"`
	RegisteredNodes                    AttributeMap                   `json:"registeredNodes,omitempty"`
	RegistrationAccessToken            string                         `json:"registrationAccessToken,omitempty"`
	RootURL                            string                         `json:"rootURL,omitempty"`
	Secret                             string                         `json:"secret,omitempty"`
	ServiceAccountsEnabled             *bool                          `json:"serviceAccountsEnabled,omitempty"`
	StandardFlowEnabled                *bool                          `json:"standardFlowEnabled,omitempty"`
	SurrogateAuthRequired              *bool                          `json:"surrogateAuthRequired,omitempty"`
	WebOrigins                         []string                       `json:"webOrigins,omitempty"`
}

// ResourceServerRepresentation represents the authorization settings for a realm client
type ResourceServerRepresentation struct {
	AllowRemoteResourceManagement *bool                    `json:"allowRemoteResourceManagement,omitempty"`
	ClientID                      string                   `json:"clientID,omitempty"`
	ID                            string                   `json:"id,omitempty"`
	Name                          string                   `json:"name,omitempty"`
	Policies                      []PolicyRepresentation   `json:"policies,omitempty"`
	PolicyEnforcementMode         string                   `json:"policyEnforcementMode,omitempty"`
	Resources                     []ResourceRepresentation `json:"resources,omitempty"`
	Scopes                        []ScopeRepresentation    `json:"scopes,omitempty"`
}

// PolicyRepresentation represents the policies attached to the
// resource server for a realm client
type PolicyRepresentation struct {
	Config           AttributeMap `json:"config,omitempty"`
	DecisionStrategy string       `json:"decisionStrategy,omitempty"`
	Description      string       `json:"description,omitempty"`
	ID               string       `json:"id,omitempty"`
	Logic            string       `json:"logic,omitempty"` //enum (POSITIVE, NEGATIVE)
	Name             string       `json:"name,omitempty"`
	Owner            string       `json:"owner,omitempty"`
	Policies         []string     `json:"policies,omitempty"`
	Resources        []string     `json:"resources,omitempty"`
	Scopes           []string     `json:"scopes,omitempty"`
	Type             string       `json:"type,omitempty"`
}

// ScopeRepresentation represents scopes defined for a
// resource server, user, or resource
type ScopeRepresentation struct {
	DisplayName string                   `json:"displayName,omitempty"`
	IconURI     string                   `json:"iconURI,omitempty"`
	ID          string                   `json:"id,omitempty"`
	Name        string                   `json:"name,omitempty"`
	Policies    []PolicyRepresentation   `json:"policies,omitempty"`
	Resources   []ResourceRepresentation `json:"resources,omitempty"`
}

// ResourceRepresentation represents resources attached to a scope
type ResourceRepresentation struct {
	ID                 string                `json:"id,omitempty"`
	Attributes         AttributeMap          `json:"attributes,omitempty"`
	DisplayName        string                `json:"displayName,omitempty"`
	IconURI            string                `json:"iconURI,omitempty"`
	Name               string                `json:"name,omitempty"`
	OwnerManagedAccess *bool                 `json:"ownerManagedAccess,omitempty"`
	Scopes             []ScopeRepresentation `json:"scopes,omitempty"`
	Type               string                `json:"type,omitempty"`
	URI                string                `json:"uri,omitempty"`
}

// ProtocolMapperRepresentation represents an individual protocol mapper on a realm client
type ProtocolMapperRepresentation struct {
	Config         AttributeMap `json:"config,omitempty"`
	ID             string       `json:"id,omitempty"`
	Name           string       `json:"name,omitempty"`
	Protocol       string       `json:"protocol,omitempty"`
	ProtocolMapper string       `json:"protocolMapper,omitempty"`
}
