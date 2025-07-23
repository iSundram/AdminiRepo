
package auth

// External authentication providers (OAuth, SAML, etc.)
type ExternalProviders struct {
    providers map[string]Provider
}

type Provider struct {
    Name         string
    ClientID     string
    ClientSecret string
    RedirectURL  string
}

func NewExternalProviders() *ExternalProviders {
    return &ExternalProviders{
        providers: make(map[string]Provider),
    }
}
