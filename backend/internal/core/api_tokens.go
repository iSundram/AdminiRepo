
package core

import "time"

// API token management
type APITokenManager struct {
    tokens map[string]*APIToken
}

type APIToken struct {
    ID          string    `json:"id"`
    UserID      string    `json:"user_id"`
    Token       string    `json:"token"`
    Permissions []string  `json:"permissions"`
    ExpiresAt   time.Time `json:"expires_at"`
}

func NewAPITokenManager() *APITokenManager {
    return &APITokenManager{
        tokens: make(map[string]*APIToken),
    }
}
