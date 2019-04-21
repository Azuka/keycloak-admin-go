//go:generate gomodifytags -file $GOFILE -struct MultivaluedHashMap -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct RealmRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct AuthenticationFlowRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct AuthenticationExecutionExportRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct AuthenticatorConfigRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct ClientScopeRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct IdentityProviderMapperRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct IdentityProviderRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct RequiredActionProviderRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct RoleRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct RoleComposites -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct ScopeMappingRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct UserFederationMapperRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct UserFederationProviderRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct GroupRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct ClientRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct ResourceServerRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct PolicyRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct ScopeRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct ResourceRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct ProtocolMapperRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct UserRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct UserConsentRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct CredentialRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct FederatedIdentityRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct UserSessionRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate gomodifytags -file $GOFILE -struct GroupRepresentation -add-options json=omitempty -add-tags json -w -transform camelcase

package keycloak

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

// AttributeMap represents a map of attributes
type AttributeMap map[string]interface{}

// UnixTime is an alias for a date time from Keycloak
// which comes in as an int32
type UnixTime time.Time

// MarshalJSON lets UnixTime implement the json.Marshaler interface
func (t UnixTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).UnixNano()/int64(time.Millisecond), 10)), nil
}

// UnmarshalJSON lets UnixTime implement the json.Unmarshaler interface
func (t *UnixTime) UnmarshalJSON(s []byte) error {
	r := strings.Replace(string(s), `"`, ``, -1)

	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(0, q*int64(time.Millisecond))
	return nil
}

func (t UnixTime) String() string {
	return time.Time(t).String()
}

// MultivaluedHashMap multivalued map
// easyjson:json
type MultivaluedHashMap struct {
	Empty      bool    `json:"empty,omitempty"`
	LoadFactor float64 `json:"loadFactor,omitempty"`
	Threshold  int32   `json:"threshold,omitempty"`
}

// RealmRepresentation represents a realm
type RealmRepresentation struct {
	AccessCodeLifespan                  int                                    `json:"accessCodeLifespan,omitempty"`
	AccessCodeLifespanLogin             int                                    `json:"accessCodeLifespanLogin,omitempty"`
	AccessCodeLifespanUserAction        int                                    `json:"accessCodeLifespanUserAction,omitempty"`
	AccessTokenLifespan                 int                                    `json:"accessTokenLifespan,omitempty"`
	AccessTokenLifespanForImplicitFlow  int                                    `json:"accessTokenLifespanForImplicitFlow,omitempty"`
	AccountTheme                        string                                 `json:"accountTheme,omitempty"`
	ActionTokenGeneratedByAdminLifespan int                                    `json:"actionTokenGeneratedByAdminLifespan,omitempty"`
	ActionTokenGeneratedByUserLifespan  int                                    `json:"actionTokenGeneratedByUserLifespan,omitempty"`
	AdminEventsDetailsEnabled           *bool                                  `json:"adminEventsDetailsEnabled,omitempty"`
	AdminEventsEnabled                  *bool                                  `json:"adminEventsEnabled,omitempty"`
	AdminTheme                          string                                 `json:"adminTheme,omitempty"`
	Attributes                          AttributeMap                           `json:"attributes,omitempty"`
	AuthenticationFlows                 []AuthenticationFlowRepresentation     `json:"authenticationFlows,omitempty"`
	AuthenticatorConfig                 []AuthenticatorConfigRepresentation    `json:"authenticatorConfig,omitempty"`
	BrowserFlow                         string                                 `json:"browserFlow,omitempty"`
	BrowserSecurityHeaders              AttributeMap                           `json:"browserSecurityHeaders,omitempty"`
	BruteForceProtected                 *bool                                  `json:"bruteForceProtected,omitempty"`
	ClientAuthenticationFlow            string                                 `json:"clientAuthenticationFlow,omitempty"`
	ClientScopeMappings                 AttributeMap                           `json:"clientScopeMappings,omitempty"`
	ClientScopes                        []ClientScopeRepresentation            `json:"clientScopes,omitempty"`
	Clients                             []ClientRepresentation                 `json:"clients,omitempty"`
	Components                          MultivaluedHashMap                     `json:"components,omitempty"`
	DefaultDefaultClientScopes          []string                               `json:"defaultDefaultClientScopes,omitempty"`
	DefaultGroups                       []string                               `json:"defaultGroups,omitempty"`
	DefaultLocale                       string                                 `json:"defaultLocale,omitempty"`
	DefaultOptionalClientScopes         []string                               `json:"defaultOptionalClientScopes,omitempty"`
	DefaultRoles                        []string                               `json:"defaultRoles,omitempty"`
	DirectGrantFlow                     string                                 `json:"directGrantFlow,omitempty"`
	DisplayName                         string                                 `json:"displayName,omitempty"`
	DisplayNameHTML                     string                                 `json:"displayNameHtml,omitempty"`
	DockerAuthenticationFlow            string                                 `json:"dockerAuthenticationFlow,omitempty"`
	DuplicateEmailsAllowed              *bool                                  `json:"duplicateEmailsAllowed,omitempty"`
	EditUsernameAllowed                 *bool                                  `json:"editUsernameAllowed,omitempty"`
	EmailTheme                          string                                 `json:"emailTheme,omitempty"`
	Enabled                             *bool                                  `json:"enabled,omitempty"`
	EnabledEventTypes                   []string                               `json:"enabledEventTypes,omitempty"`
	EventsEnabled                       *bool                                  `json:"eventsEnabled,omitempty"`
	EventsExpiration                    int                                    `json:"eventsExpiration,omitempty"`
	EventsListeners                     []string                               `json:"eventsListeners,omitempty"`
	FailureFactor                       int                                    `json:"failureFactor,omitempty"`
	FederatedUsers                      []UserRepresentation                   `json:"federatedUsers,omitempty"`
	Groups                              []GroupRepresentation                  `json:"groups,omitempty"`
	ID                                  string                                 `json:"id,omitempty"`
	IdentityProviderMappers             []IdentityProviderMapperRepresentation `json:"identityProviderMappers,omitempty"`
	IdentityProviders                   []IdentityProviderRepresentation       `json:"identityProviders,omitempty"`
	InternationalizationEnabled         *bool                                  `json:"internationalizationEnabled,omitempty"`
	KeycloakVersion                     string                                 `json:"keycloakVersion,omitempty"`
	LoginTheme                          string                                 `json:"loginTheme,omitempty"`
	LoginWithEmailAllowed               *bool                                  `json:"loginWithEmailAllowed,omitempty"`
	MaxDeltaTimeSeconds                 int                                    `json:"maxDeltaTimeSeconds,omitempty"`
	MaxFailureWaitSeconds               int                                    `json:"maxFailureWaitSeconds,omitempty"`
	MinimumQuickLoginWaitSeconds        int                                    `json:"minimumQuickLoginWaitSeconds,omitempty"`
	NotBefore                           int                                    `json:"notBefore,omitempty"`
	OfflineSessionIdleTimeout           int                                    `json:"offlineSessionIdleTimeout,omitempty"`
	OtpPolicyAlgorithm                  string                                 `json:"otpPolicyAlgorithm,omitempty"`
	OtpPolicyDigits                     int                                    `json:"otpPolicyDigits,omitempty"`
	OtpPolicyLookAheadWindow            int                                    `json:"otpPolicyLookAheadWindow,omitempty"`
	OtpPolicyPeriod                     int                                    `json:"otpPolicyPeriod,omitempty"`
	OtpPolicyType                       string                                 `json:"otpPolicyType,omitempty"`
	OtpSupportedApplications            []string                               `json:"otpSupportedApplications,omitempty"`
	PasswordPolicy                      string                                 `json:"passwordPolicy,omitempty"`
	PermanentLockout                    *bool                                  `json:"permanentLockout,omitempty"`
	ProtocolMappers                     []ProtocolMapperRepresentation         `json:"protocolMappers,omitempty"`
	QuickLoginCheckMilliSeconds         int                                    `json:"quickLoginCheckMilliSeconds,omitempty"`
	Realm                               string                                 `json:"realm,omitempty"`
	RefreshTokenMaxReuse                int                                    `json:"refreshTokenMaxReuse,omitempty"`
	RegistrationAllowed                 *bool                                  `json:"registrationAllowed,omitempty"`
	RegistrationEmailAsUsername         *bool                                  `json:"registrationEmailAsUsername,omitempty"`
	RegistrationFlow                    string                                 `json:"registrationFlow,omitempty"`
	RememberMe                          *bool                                  `json:"rememberMe,omitempty"`
	RequiredActions                     []RequiredActionProviderRepresentation `json:"requiredActions,omitempty"`
	ResetCredentialsFlow                string                                 `json:"resetCredentialsFlow,omitempty"`
	ResetPasswordAllowed                *bool                                  `json:"resetPasswordAllowed,omitempty"`
	RevokeRefreshToken                  *bool                                  `json:"revokeRefreshToken,omitempty"`
	Roles                               RolesRepresentation                    `json:"roles,omitempty"`
	ScopeMappings                       []ScopeMappingRepresentation           `json:"scopeMappings,omitempty"`
	SMTPServer                          AttributeMap                           `json:"smtpServer,omitempty"`
	SslRequired                         string                                 `json:"sslRequired,omitempty"`
	SsoSessionIdleTimeout               int                                    `json:"ssoSessionIdleTimeout,omitempty"`
	SsoSessionMaxLifespan               int                                    `json:"ssoSessionMaxLifespan,omitempty"`
	SupportedLocales                    []string                               `json:"supportedLocales,omitempty"`
	UserFederationMappers               []UserFederationMapperRepresentation   `json:"userFederationMappers,omitempty"`
	UserFederationProviders             []UserFederationProviderRepresentation `json:"userFederationProviders,omitempty"`
	UserManagedAccessAllowed            *bool                                  `json:"userManagedAccessAllowed,omitempty"`
	Users                               []UserRepresentation                   `json:"users,omitempty"`
	VerifyEmail                         *bool                                  `json:"verifyEmail,omitempty"`
	WaitIncrementSeconds                int                                    `json:"waitIncrementSeconds,omitempty"`
}

// AuthenticationFlowRepresentation for representing Flows
type AuthenticationFlowRepresentation struct {
	Alias                    string                                        `json:"alias,omitempty"`
	AuthenticationExecutions []AuthenticationExecutionExportRepresentation `json:"authenticationExecutions,omitempty"`
	BuiltIn                  *bool                                         `json:"builtIn,omitempty"`
	Description              string                                        `json:"description,omitempty"`
	ID                       string                                        `json:"id,omitempty"`
	ProviderID               string                                        `json:"providerID,omitempty"`
	TopLevel                 *bool                                         `json:"topLevel,omitempty"`
}

// AuthenticationExecutionExportRepresentation for Authenticator Execution
type AuthenticationExecutionExportRepresentation struct {
	Authenticator       string `json:"authenticator,omitempty"`
	AuthenticatorConfig string `json:"authenticatorConfig,omitempty"`
	AuthenticatorFlow   *bool  `json:"authenticatorFlow,omitempty"`
	AutheticatorFlow    *bool  `json:"autheticatorFlow,omitempty"`
	FlowAlias           string `json:"flowAlias,omitempty"`
	Priority            int    `json:"priority,omitempty"`
	Requirement         string `json:"requirement,omitempty"`
	UserSetupAllowed    *bool  `json:"userSetupAllowed,omitempty"`
}

// AuthenticatorConfigRepresentation Authenticator Config
type AuthenticatorConfigRepresentation struct {
	Alias  string       `json:"alias,omitempty"`
	Config AttributeMap `json:"config,omitempty"`
	ID     string       `json:"id,omitempty"`
}

// ClientScopeRepresentation Client Scope
type ClientScopeRepresentation struct {
	Attributes      AttributeMap                   `json:"attributes,omitempty"`
	Description     string                         `json:"description,omitempty"`
	ID              string                         `json:"id,omitempty"`
	Name            string                         `json:"name,omitempty"`
	Protocol        string                         `json:"protocol,omitempty"`
	ProtocolMappers []ProtocolMapperRepresentation `json:"protocolMappers,omitempty"`
}

// IdentityProviderMapperRepresentation Identity Provider Mapper
type IdentityProviderMapperRepresentation struct {
	Config                 AttributeMap `json:"config,omitempty"`
	ID                     string       `json:"id,omitempty"`
	IdentityProviderAlias  string       `json:"identityProviderAlias,omitempty"`
	IdentityProviderMapper string       `json:"identityProviderMapper,omitempty"`
	Name                   string       `json:"name,omitempty"`
}

// IdentityProviderRepresentation Identity Provider
type IdentityProviderRepresentation struct {
	AddReadTokenRoleOnCreate  *bool        `json:"addReadTokenRoleOnCreate,omitempty"`
	Alias                     string       `json:"alias,omitempty"`
	Config                    AttributeMap `json:"config,omitempty"`
	DisplayName               string       `json:"displayName,omitempty"`
	Enabled                   *bool        `json:"enabled,omitempty"`
	FirstBrokerLoginFlowAlias string       `json:"firstBrokerLoginFlowAlias,omitempty"`
	InternalID                string       `json:"internalID,omitempty"`
	LinkOnly                  *bool        `json:"linkOnly,omitempty"`
	PostBrokerLoginFlowAlias  string       `json:"postBrokerLoginFlowAlias,omitempty"`
	ProviderID                string       `json:"providerID,omitempty"`
	StoreToken                *bool        `json:"storeToken,omitempty"`
	TrustEmail                *bool        `json:"trustEmail,omitempty"`
}

// RequiredActionProviderRepresentation Required Action Provider
type RequiredActionProviderRepresentation struct {
	Alias         string       `json:"alias,omitempty"`
	Config        AttributeMap `json:"config,omitempty"`
	DefaultAction *bool        `json:"defaultAction,omitempty"`
	Enabled       *bool        `json:"enabled,omitempty"`
	Name          string       `json:"name,omitempty"`
	ProviderID    string       `json:"providerID,omitempty"`
}

// RolesRepresentation Roles Representation
type RolesRepresentation struct {
	Client AttributeMap         `json:"client,omitempty"`
	Realm  []RoleRepresentation `json:"realm,omitempty"`
}

// RoleRepresentation Role
type RoleRepresentation struct {
	ClientRole  *bool          `json:"clientRole,omitempty"`
	Composite   *bool          `json:"composite,omitempty"`
	Composites  RoleComposites `json:"composites,omitempty"`
	ContainerID string         `json:"containerID,omitempty"`
	Description string         `json:"description,omitempty"`
	ID          string         `json:"id,omitempty"`
	Name        string         `json:"name,omitempty"`
}

// RoleComposites known in keycloak as a "RoleRepresentations-Composites" in
// in the source it is just an inner-class.
type RoleComposites struct {
	Client AttributeMap `json:"client,omitempty"`
	Realm  []string     `json:"realm,omitempty"`
}

// ScopeMappingRepresentation Scope Mapping
type ScopeMappingRepresentation struct {
	Client      string   `json:"client,omitempty"`
	ClientScope string   `json:"clientScope,omitempty"`
	Roles       []string `json:"roles,omitempty"`
	Self        string   `json:"self,omitempty"`
}

// UserFederationMapperRepresentation User Federation
type UserFederationMapperRepresentation struct {
	Config                        AttributeMap `json:"config,omitempty"`
	FederationMapperType          string       `json:"federationMapperType,omitempty"`
	FederationProviderDisplayName string       `json:"federationProviderDisplayName,omitempty"`
	ID                            string       `json:"id,omitempty"`
	Name                          string       `json:"name,omitempty"`
}

// UserFederationProviderRepresentation User federation provider
type UserFederationProviderRepresentation struct {
	ChangedSyncPeriod int32        `json:"changedSyncPeriod,omitempty"`
	Config            AttributeMap `json:"config,omitempty"`
	DisplayName       string       `json:"displayName,omitempty"`
	FullSyncPeriod    int32        `json:"fullSyncPeriod,omitempty"`
	ID                string       `json:"id,omitempty"`
	LastSync          int          `json:"lastSync,omitempty"`
	Priority          int32        `json:"priority,omitempty"`
	ProviderName      string       `json:"providerName,omitempty"`
}

type MappingRepresentation struct {
	RealmMappings  []RoleRepresentation `json:"realmMappings,omitempty"`
	ClientMappings bool                 `json:"clientMappings,omitempty"`
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
	FederationLink         *url.URL                          `json:"federationLink,omitempty"`
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
