// Code generated by "make api"; DO NOT EDIT.
package authmethods

type OidcAuthMethodAttributes struct {
	State                             string   `json:"state,omitempty"`
	Issuer                            string   `json:"issuer,omitempty"`
	ClientId                          string   `json:"client_id,omitempty"`
	ClientSecret                      string   `json:"client_secret,omitempty"`
	ClientSecretHmac                  string   `json:"client_secret_hmac,omitempty"`
	MaxAge                            uint32   `json:"max_age,omitempty"`
	SigningAlgorithms                 []string `json:"signing_algorithms,omitempty"`
	ApiUrlPrefix                      string   `json:"api_url_prefix,omitempty"`
	CallbackUrl                       string   `json:"callback_url,omitempty"`
	IdpCaCerts                        []string `json:"idp_ca_certs,omitempty"`
	AllowedAudiences                  []string `json:"allowed_audiences,omitempty"`
	ClaimsScopes                      []string `json:"claims_scopes,omitempty"`
	AccountClaimMaps                  []string `json:"account_claim_maps,omitempty"`
	DisableDiscoveredConfigValidation bool     `json:"disable_discovered_config_validation,omitempty"`
	DryRun                            bool     `json:"dry_run,omitempty"`
}
