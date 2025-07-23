
package config

// Theme management
type ThemeManager struct {
    themes map[string]*Theme
}

type Theme struct {
    ID          string            `json:"id"`
    Name        string            `json:"name"`
    Colors      map[string]string `json:"colors"`
    Fonts       map[string]string `json:"fonts"`
    CustomCSS   string            `json:"custom_css"`
}

func NewThemeManager() *ThemeManager {
    return &ThemeManager{
        themes: make(map[string]*Theme),
    }
}
