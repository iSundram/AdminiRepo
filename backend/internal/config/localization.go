
package config

// Localization and internationalization
type LocalizationManager struct {
    languages map[string]*Language
}

type Language struct {
    Code        string            `json:"code"`
    Name        string            `json:"name"`
    Translations map[string]string `json:"translations"`
}

func NewLocalizationManager() *LocalizationManager {
    return &LocalizationManager{
        languages: make(map[string]*Language),
    }
}
