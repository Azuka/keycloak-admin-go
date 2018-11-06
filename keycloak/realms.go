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

package keycloak

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
