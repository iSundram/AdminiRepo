
package dns

// DNS template management
type TemplateManager struct {
    templates map[string]*DNSTemplate
}

type DNSTemplate struct {
    ID       string   `json:"id"`
    Name     string   `json:"name"`
    Records  []Record `json:"records"`
    Category string   `json:"category"`
}

func NewTemplateManager() *TemplateManager {
    return &TemplateManager{
        templates: make(map[string]*DNSTemplate),
    }
}
