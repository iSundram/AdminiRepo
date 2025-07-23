
package auth

// Single Sign-On module
type SSOManager struct {
    providers map[string]SSOProvider
}

type SSOProvider struct {
    Name     string
    Type     string
    Config   map[string]string
    Enabled  bool
}

func NewSSOManager() *SSOManager {
    return &SSOManager{
        providers: make(map[string]SSOProvider),
    }
}
