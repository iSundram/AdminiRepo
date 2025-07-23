
package config

// Secrets management
type SecretsManager struct {
    secrets map[string]string
}

func NewSecretsManager() *SecretsManager {
    return &SecretsManager{
        secrets: make(map[string]string),
    }
}

func (sm *SecretsManager) GetSecret(key string) string {
    return sm.secrets[key]
}
