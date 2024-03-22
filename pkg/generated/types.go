// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package generated

import (
	"encoding/json"
	"fmt"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

const (
	Oauth2AuthenticationScopes = "oauth2Authentication.Scopes"
)

// Defines values for AuthMethod.
const (
	ClientSecretBasic AuthMethod = "client_secret_basic"
	ClientSecretPost  AuthMethod = "client_secret_post"
)

// Defines values for Claim.
const (
	ClaimAud           Claim = "aud"
	ClaimEmail         Claim = "email"
	ClaimEmailVerified Claim = "email_verified"
	ClaimExp           Claim = "exp"
	ClaimFamilityName  Claim = "famility_name"
	ClaimGivenName     Claim = "given_name"
	ClaimIat           Claim = "iat"
	ClaimIss           Claim = "iss"
	ClaimLocale        Claim = "locale"
	ClaimName          Claim = "name"
	ClaimPicture       Claim = "picture"
	ClaimSub           Claim = "sub"
)

// Defines values for CodeChallengeMethod.
const (
	Plain CodeChallengeMethod = "plain"
	S256  CodeChallengeMethod = "S256"
)

// Defines values for GrantType.
const (
	AuthorizationCode GrantType = "authorization_code"
	RefreshToken      GrantType = "refresh_token"
)

// Defines values for Oauth2ErrorError.
const (
	AccessDenied            Oauth2ErrorError = "access_denied"
	Conflict                Oauth2ErrorError = "conflict"
	Forbidden               Oauth2ErrorError = "forbidden"
	InvalidClient           Oauth2ErrorError = "invalid_client"
	InvalidGrant            Oauth2ErrorError = "invalid_grant"
	InvalidRequest          Oauth2ErrorError = "invalid_request"
	InvalidScope            Oauth2ErrorError = "invalid_scope"
	MethodNotAllowed        Oauth2ErrorError = "method_not_allowed"
	NotFound                Oauth2ErrorError = "not_found"
	ServerError             Oauth2ErrorError = "server_error"
	TemporarilyUnavailable  Oauth2ErrorError = "temporarily_unavailable"
	UnauthorizedClient      Oauth2ErrorError = "unauthorized_client"
	UnsupportedGrantType    Oauth2ErrorError = "unsupported_grant_type"
	UnsupportedMediaType    Oauth2ErrorError = "unsupported_media_type"
	UnsupportedResponseType Oauth2ErrorError = "unsupported_response_type"
)

// Defines values for ResponseType.
const (
	ResponseTypeCode             ResponseType = "code"
	ResponseTypeCodeIdToken      ResponseType = "code id_token"
	ResponseTypeCodeToken        ResponseType = "code token"
	ResponseTypeCodeTokenIdToken ResponseType = "code token id_token"
	ResponseTypeIdToken          ResponseType = "id_token"
	ResponseTypeNone             ResponseType = "none"
	ResponseTypeToken            ResponseType = "token"
	ResponseTypeTokenIdToken     ResponseType = "token id_token"
)

// Defines values for Scope.
const (
	ScopeEmail   Scope = "email"
	ScopeOpenid  Scope = "openid"
	ScopeProfile Scope = "profile"
)

// Defines values for SigningAlgorithm.
const (
	ES512 SigningAlgorithm = "ES512"
)

// AuthMethod Supported authentication methods.
type AuthMethod string

// Claim Supported claims.
type Claim string

// CodeChallengeMethod Supported code challenge methods.
type CodeChallengeMethod string

// GrantType Supported grant type.
type GrantType string

// JsonWebKey JSON web key. See the relevant JWKS documentation for further details.
type JsonWebKey = map[string]interface{}

// JsonWebKeySet JSON web key set. This data type is defined by an external 3rd party standards
// committee. Consult the relevant documentation for further details.
type JsonWebKeySet struct {
	Keys *[]JsonWebKey `json:"keys,omitempty"`
}

// LoginRequestOptions Login request options.
type LoginRequestOptions struct {
	// Email The user's email address.
	Email *string `json:"email"`

	// Provider The explcit provider type.
	Provider *string `json:"provider"`

	// State The state string supplied by the authorization endpoint.
	State string `json:"state"`
}

// Oauth2Error Generic error message.
type Oauth2Error struct {
	// Error A terse error string expanding on the HTTP error code. Errors are based on the OAuth2 specification, but are expanded with proprietary status codes for APIs other than those specified by OAuth2.
	Error Oauth2ErrorError `json:"error"`

	// ErrorDescription Verbose message describing the error.
	ErrorDescription string `json:"error_description"`
}

// Oauth2ErrorError A terse error string expanding on the HTTP error code. Errors are based on the OAuth2 specification, but are expanded with proprietary status codes for APIs other than those specified by OAuth2.
type Oauth2ErrorError string

// Oauth2Provider An oauth2 provider.
type Oauth2Provider struct {
	ClientID    string `json:"clientID"`
	DisplayName string `json:"displayName"`
	Issuer      string `json:"issuer"`
	Name        string `json:"name"`
}

// Oauth2Providers A list of oauth2 providers.
type Oauth2Providers = []Oauth2Provider

// OpenidConfiguration OpenID configuration.
type OpenidConfiguration struct {
	// AuthorizationEndpoint The oauth2 endpoint that initiates authentication.
	AuthorizationEndpoint string `json:"authorization_endpoint"`

	// ClaimsSupported A list of supported claims
	ClaimsSupported []Claim `json:"claims_supported"`

	// CodeChallengeMethodsSupported A list of code challenge methods supported.
	CodeChallengeMethodsSupported []CodeChallengeMethod `json:"code_challenge_methods_supported"`

	// GrantTypesSupported A list of supported grants for the token endpoint.
	GrantTypesSupported []GrantType `json:"grant_types_supported"`

	// IdTokenSigningAlgValuesSupported A list of signing algorithms supported for ID tokens.
	IdTokenSigningAlgValuesSupported []SigningAlgorithm `json:"id_token_signing_alg_values_supported"`

	// Issuer The OpenID Issuer (iss field).
	Issuer string `json:"issuer"`

	// JwksUri The oauth2 endpoint that exposes public signing keys for token validation.
	JwksUri string `json:"jwks_uri"`

	// ResponseTypesSupported A list of supported response types that can be requested for the authorization endpoint.
	ResponseTypesSupported []ResponseType `json:"response_types_supported"`

	// ScopesSupported A list of supported oauth2 scopes.
	ScopesSupported []Scope `json:"scopes_supported"`

	// TokenEndpoint The oauth2 endpoint that is used to exchange an authentication code for tokens.
	TokenEndpoint string `json:"token_endpoint"`

	// TokenEndpointAuthMethodsSupported A list of supported authentication methods for the token endpoint.
	TokenEndpointAuthMethodsSupported []AuthMethod `json:"token_endpoint_auth_methods_supported"`

	// UserinfoEndpoint The oidc endpoint used to get information about an access token's user.
	UserinfoEndpoint string `json:"userinfo_endpoint"`
}

// Organization An organization.
type Organization struct {
	// Domain An email domain that identifies all users.
	Domain *string `json:"domain,omitempty"`

	// Name A unique organization name.
	Name string `json:"name"`

	// ProviderName An identity provider for the organization. This must be an oauth2 provider name.
	ProviderName *string `json:"providerName,omitempty"`
}

// Organizations A list of organizations.
type Organizations = []Organization

// ResponseType Supported response types.
type ResponseType string

// Scope Supported scopes.
type Scope string

// SigningAlgorithm Supported signing algorithms.
type SigningAlgorithm string

// Token Oauth2 token result.
type Token struct {
	// AccessToken The opaque access token.
	AccessToken string `json:"access_token"`

	// ExpiresIn The time in seconds the token will last for.
	ExpiresIn int `json:"expires_in"`

	// IdToken An OIDC ID token.
	IdToken *string `json:"id_token,omitempty"`

	// TokenType How the access token is to be presented to the resource server.
	TokenType string `json:"token_type"`
}

// TokenRequestOptions oauth2 token endpoint.
type TokenRequestOptions struct {
	// ClientId Client ID. Required with the "code" grant type.
	ClientId *string `json:"client_id"`

	// Code Authorization code. Required with the "code" grant type.
	Code *string `json:"code"`

	// CodeVerifier Client code verifier. Required with the "code" grant type.
	CodeVerifier *string `json:"code_verifier"`

	// GrantType Supported grant type.  Must be either "code" or "password".
	GrantType string `json:"grant_type"`

	// Password Resource owner password. Required with the "password" grant type.
	Password *string `json:"password"`

	// RedirectUri Client redirect URI. Required with the "code" grant type.
	RedirectUri *string `json:"redirect_uri"`

	// Username Resource owner username. Required with the "password" grant type.
	Username *string `json:"username"`
	union    json.RawMessage
}

// TokenRequestOptions0 defines model for .
type TokenRequestOptions0 struct {
	GrantType *interface{} `json:"grant_type,omitempty"`
}

// TokenRequestOptions1 defines model for .
type TokenRequestOptions1 struct {
	GrantType *interface{} `json:"grant_type,omitempty"`
}

// GroupidParameter defines model for groupidParameter.
type GroupidParameter = string

// OrganizationParameter defines model for organizationParameter.
type OrganizationParameter = string

// BadRequestResponse Generic error message.
type BadRequestResponse = Oauth2Error

// ConflictResponse Generic error message.
type ConflictResponse = Oauth2Error

// ForbiddenResponse Generic error message.
type ForbiddenResponse = Oauth2Error

// InternalServerErrorResponse Generic error message.
type InternalServerErrorResponse = Oauth2Error

// JwksResponse JSON web key set. This data type is defined by an external 3rd party standards
// committee. Consult the relevant documentation for further details.
type JwksResponse = JsonWebKeySet

// Oauth2ProvidersResponse A list of oauth2 providers.
type Oauth2ProvidersResponse = Oauth2Providers

// OpenidConfigurationResponse OpenID configuration.
type OpenidConfigurationResponse = OpenidConfiguration

// OrganizationsResponse A list of organizations.
type OrganizationsResponse = Organizations

// TokenResponse Oauth2 token result.
type TokenResponse = Token

// UnauthorizedResponse Generic error message.
type UnauthorizedResponse = Oauth2Error

// UserinfoResponse defines model for userinfoResponse.
type UserinfoResponse interface{}

// CreateGroupRequest defines model for createGroupRequest.
type CreateGroupRequest = map[string]interface{}

// CreateOrganizationRequest An organization.
type CreateOrganizationRequest = Organization

// UpdateGroupRequest defines model for updateGroupRequest.
type UpdateGroupRequest = map[string]interface{}

// UpdateOrganizationRequest An organization.
type UpdateOrganizationRequest = Organization

// PostApiV1OrganizationsOrganizationGroupsJSONBody defines parameters for PostApiV1OrganizationsOrganizationGroups.
type PostApiV1OrganizationsOrganizationGroupsJSONBody = map[string]interface{}

// PutApiV1OrganizationsOrganizationGroupsGroupidJSONBody defines parameters for PutApiV1OrganizationsOrganizationGroupsGroupid.
type PutApiV1OrganizationsOrganizationGroupsGroupidJSONBody = map[string]interface{}

// PostApiV1OrganizationsJSONRequestBody defines body for PostApiV1Organizations for application/json ContentType.
type PostApiV1OrganizationsJSONRequestBody = Organization

// PutApiV1OrganizationsOrganizationJSONRequestBody defines body for PutApiV1OrganizationsOrganization for application/json ContentType.
type PutApiV1OrganizationsOrganizationJSONRequestBody = Organization

// PostApiV1OrganizationsOrganizationGroupsJSONRequestBody defines body for PostApiV1OrganizationsOrganizationGroups for application/json ContentType.
type PostApiV1OrganizationsOrganizationGroupsJSONRequestBody = PostApiV1OrganizationsOrganizationGroupsJSONBody

// PutApiV1OrganizationsOrganizationGroupsGroupidJSONRequestBody defines body for PutApiV1OrganizationsOrganizationGroupsGroupid for application/json ContentType.
type PutApiV1OrganizationsOrganizationGroupsGroupidJSONRequestBody = PutApiV1OrganizationsOrganizationGroupsGroupidJSONBody

// PostOauth2V2LoginFormdataRequestBody defines body for PostOauth2V2Login for application/x-www-form-urlencoded ContentType.
type PostOauth2V2LoginFormdataRequestBody = LoginRequestOptions

// PostOauth2V2TokenFormdataRequestBody defines body for PostOauth2V2Token for application/x-www-form-urlencoded ContentType.
type PostOauth2V2TokenFormdataRequestBody = TokenRequestOptions

// AsTokenRequestOptions0 returns the union data inside the TokenRequestOptions as a TokenRequestOptions0
func (t TokenRequestOptions) AsTokenRequestOptions0() (TokenRequestOptions0, error) {
	var body TokenRequestOptions0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTokenRequestOptions0 overwrites any union data inside the TokenRequestOptions as the provided TokenRequestOptions0
func (t *TokenRequestOptions) FromTokenRequestOptions0(v TokenRequestOptions0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTokenRequestOptions0 performs a merge with any union data inside the TokenRequestOptions, using the provided TokenRequestOptions0
func (t *TokenRequestOptions) MergeTokenRequestOptions0(v TokenRequestOptions0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsTokenRequestOptions1 returns the union data inside the TokenRequestOptions as a TokenRequestOptions1
func (t TokenRequestOptions) AsTokenRequestOptions1() (TokenRequestOptions1, error) {
	var body TokenRequestOptions1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTokenRequestOptions1 overwrites any union data inside the TokenRequestOptions as the provided TokenRequestOptions1
func (t *TokenRequestOptions) FromTokenRequestOptions1(v TokenRequestOptions1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTokenRequestOptions1 performs a merge with any union data inside the TokenRequestOptions, using the provided TokenRequestOptions1
func (t *TokenRequestOptions) MergeTokenRequestOptions1(v TokenRequestOptions1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

func (t TokenRequestOptions) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	if err != nil {
		return nil, err
	}
	object := make(map[string]json.RawMessage)
	if t.union != nil {
		err = json.Unmarshal(b, &object)
		if err != nil {
			return nil, err
		}
	}

	if t.ClientId != nil {
		object["client_id"], err = json.Marshal(t.ClientId)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'client_id': %w", err)
		}
	}

	if t.Code != nil {
		object["code"], err = json.Marshal(t.Code)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'code': %w", err)
		}
	}

	if t.CodeVerifier != nil {
		object["code_verifier"], err = json.Marshal(t.CodeVerifier)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'code_verifier': %w", err)
		}
	}

	object["grant_type"], err = json.Marshal(t.GrantType)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'grant_type': %w", err)
	}

	if t.Password != nil {
		object["password"], err = json.Marshal(t.Password)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'password': %w", err)
		}
	}

	if t.RedirectUri != nil {
		object["redirect_uri"], err = json.Marshal(t.RedirectUri)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'redirect_uri': %w", err)
		}
	}

	if t.Username != nil {
		object["username"], err = json.Marshal(t.Username)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'username': %w", err)
		}
	}
	b, err = json.Marshal(object)
	return b, err
}

func (t *TokenRequestOptions) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	if err != nil {
		return err
	}
	object := make(map[string]json.RawMessage)
	err = json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["client_id"]; found {
		err = json.Unmarshal(raw, &t.ClientId)
		if err != nil {
			return fmt.Errorf("error reading 'client_id': %w", err)
		}
	}

	if raw, found := object["code"]; found {
		err = json.Unmarshal(raw, &t.Code)
		if err != nil {
			return fmt.Errorf("error reading 'code': %w", err)
		}
	}

	if raw, found := object["code_verifier"]; found {
		err = json.Unmarshal(raw, &t.CodeVerifier)
		if err != nil {
			return fmt.Errorf("error reading 'code_verifier': %w", err)
		}
	}

	if raw, found := object["grant_type"]; found {
		err = json.Unmarshal(raw, &t.GrantType)
		if err != nil {
			return fmt.Errorf("error reading 'grant_type': %w", err)
		}
	}

	if raw, found := object["password"]; found {
		err = json.Unmarshal(raw, &t.Password)
		if err != nil {
			return fmt.Errorf("error reading 'password': %w", err)
		}
	}

	if raw, found := object["redirect_uri"]; found {
		err = json.Unmarshal(raw, &t.RedirectUri)
		if err != nil {
			return fmt.Errorf("error reading 'redirect_uri': %w", err)
		}
	}

	if raw, found := object["username"]; found {
		err = json.Unmarshal(raw, &t.Username)
		if err != nil {
			return fmt.Errorf("error reading 'username': %w", err)
		}
	}

	return err
}
