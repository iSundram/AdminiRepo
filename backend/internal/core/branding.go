
package core

// Branding management
type BrandingManager struct {
    brands map[string]*Brand
}

type Brand struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Logo        string `json:"logo"`
    Colors      map[string]string `json:"colors"`
    CustomCSS   string `json:"custom_css"`
}

func NewBrandingManager() *BrandingManager {
    return &BrandingManager{
        brands: make(map[string]*Brand),
    }
}
