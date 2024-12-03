package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID                                         string
	ClientName                                 string
	ClientSecret                               string
	Scope                                      string
	Owner                                      string
	PolicyURI                                  string
	TosURI                                     string
	ClientURI                                  string
	LogoURI                                    string
	ClientSecretExpiresAt                      int32
	SectorIdentifierURI                        string
	Jwks                                       string
	JwksURI                                    string
	TokenEndpointAuthMethod                    string
	RequestObjectSigningAlg                    string
	UserinfoSignedResponseAlg                  string
	SubjectType                                string
	PkDeprecated                               int32
	CreatedAt                                  time.Time
	UpdatedAt                                  time.Time
	FrontchannelLogoutURI                      string
	FrontchannelLogoutSessionRequired          bool
	BackchannelLogoutURI                       string
	BackchannelLogoutSessionRequired           bool
	Metadata                                   string
	TokenEndpointAuthSigningAlg                string
	AuthorizationCodeGrantAccessTokenLifespan  sql.NullInt64
	AuthorizationCodeGrantIDTokenLifespan      sql.NullInt64
	AuthorizationCodeGrantRefreshTokenLifespan sql.NullInt64
	ClientCredentialsGrantAccessTokenLifespan  sql.NullInt64
	ImplicitGrantAccessTokenLifespan           sql.NullInt64
	ImplicitGrantIDTokenLifespan               sql.NullInt64
	JwtBearerGrantAccessTokenLifespan          sql.NullInt64
	PasswordGrantAccessTokenLifespan           sql.NullInt64
	PasswordGrantRefreshTokenLifespan          sql.NullInt64
	RefreshTokenGrantIDTokenLifespan           sql.NullInt64
	RefreshTokenGrantAccessTokenLifespan       sql.NullInt64
	RefreshTokenGrantRefreshTokenLifespan      sql.NullInt64
	PK                                         uuid.UUID
	RegistrationAccessTokenSignature           string
	NID                                        uuid.UUID
	RedirectURIs                               []byte
	GrantTypes                                 []byte
	ResponseTypes                              []byte
	Audience                                   []byte
	AllowedCORSOrigins                         []byte
	Contacts                                   []byte
	RequestURIs                                []byte
	PostLogoutRedirectURIs                     []byte
	AccessTokenStrategy                        string
	SkipConsent                                bool
	SkipLogoutConsent                          *bool
}
