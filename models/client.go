package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID                                         string `gorm:"primaryKey;column:id"`
	ClientName                                 string `gorm:"column:client_name"`
	ClientSecret                               string `gorm:"column:client_secret"`
	Scope                                      string `gorm:"column:scope"`
	Owner                                      string `gorm:"column:owner"`
	PolicyURI                                  string `gorm:"column:policy_uri"`
	TosURI                                     string `gorm:"column:tos_uri"`
	ClientURI                                  string
	LogoURI                                    string `gorm:"column:logo_uri"`
	ClientSecretExpiresAt                      int32  `gorm:"column:client_secret_expires_at"`
	SectorIdentifierURI                        string `gorm:"column:sector_identifier_uri"`
	Jwks                                       string `gorm:"column:jwks"`
	JwksURI                                    string `gorm:"column:jwks_uri"`
	TokenEndpointAuthMethod                    string `gorm:"column:token_endpoint_auth_method"`
	RequestObjectSigningAlg                    string `gorm:"column:request_object_signing_alg"`
	UserinfoSignedResponseAlg                  string `gorm:"column:userinfo_signed_response_alg"`
	SubjectType                                string `gorm:"column:subject_type"`
	PkDeprecated                               int32  `gorm:"column:pk_deprecated"`
	CreatedAt                                  time.Time
	UpdatedAt                                  time.Time
	FrontchannelLogoutURI                      string        `gorm:"column:frontchannel_logout_uri"`
	FrontchannelLogoutSessionRequired          bool          `gorm:"column:frontchannel_logout_session_required"`
	BackchannelLogoutURI                       string        `gorm:"column:backchannel_logout_uri"`
	BackchannelLogoutSessionRequired           bool          `gorm:"column:backchannel_logout_session_required"`
	Metadata                                   string        `gorm:"column:metadata"`
	TokenEndpointAuthSigningAlg                string        `gorm:"column:token_endpoint_auth_signing_alg"`
	AuthorizationCodeGrantAccessTokenLifespan  sql.NullInt64 `gorm:"column:authorization_code_grant_access_token_lifespan"`
	AuthorizationCodeGrantIDTokenLifespan      sql.NullInt64 `gorm:"column:authorization_code_grant_id_token_lifespan"`
	AuthorizationCodeGrantRefreshTokenLifespan sql.NullInt64 `gorm:"column:authorization_code_grant_refresh_token_lifespan"`
	ClientCredentialsGrantAccessTokenLifespan  sql.NullInt64 `gorm:"column:client_credentials_grant_access_token_lifespan"`
	ImplicitGrantAccessTokenLifespan           sql.NullInt64 `gorm:"column:implicit_grant_access_token_lifespan"`
	ImplicitGrantIDTokenLifespan               sql.NullInt64 `gorm:"column:implicit_grant_id_token_lifespan"`
	JwtBearerGrantAccessTokenLifespan          sql.NullInt64 `gorm:"column:jwt_bearer_grant_access_token_lifespan"`
	PasswordGrantAccessTokenLifespan           sql.NullInt64 `gorm:"column:password_grant_access_token_lifespan"`
	PasswordGrantRefreshTokenLifespan          sql.NullInt64 `gorm:"column:password_grant_refresh_token_lifespan"`
	RefreshTokenGrantIDTokenLifespan           sql.NullInt64 `gorm:"column:refresh_token_grant_id_token_lifespan"`
	RefreshTokenGrantAccessTokenLifespan       sql.NullInt64 `gorm:"column:refresh_token_grant_access_token_lifespan"`
	RefreshTokenGrantRefreshTokenLifespan      sql.NullInt64 `gorm:"column:refresh_token_grant_refresh_token_lifespan"`
	PK                                         uuid.UUID     `gorm:"column:pk"`
	RegistrationAccessTokenSignature           string        `gorm:"column:registration_access_token_signature"`
	NID                                        uuid.UUID     `gorm:"column:nid"`
	RedirectURIs                               []byte        `gorm:"column:redirect_uris"`
	GrantTypes                                 []byte        `gorm:"column:grant_types"`
	ResponseTypes                              []byte        `gorm:"column:response_types"`
	Audience                                   []byte        `gorm:"column:audience"`
	AllowedCORSOrigins                         []byte        `gorm:"column:allowed_cors_origins"`
	Contacts                                   []byte        `gorm:"column:contacts"`
	RequestURIs                                []byte        `gorm:"column:request_uris"`
	PostLogoutRedirectURIs                     []byte        `gorm:"column:post_logout_redirect_uris"`
	AccessTokenStrategy                        string        `gorm:"column:access_token_strategy"`
	SkipConsent                                bool          `gorm:"column:skip_consent"`
	SkipLogoutConsent                          *bool         `gorm:"column:skip_logout_consent"`
}
