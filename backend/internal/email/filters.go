
package email

// Email filter management
type FilterManager struct {
    filters map[string]*EmailFilter
}

type EmailFilter struct {
    ID        string                 `json:"id"`
    Name      string                 `json:"name"`
    Rules     []FilterRule           `json:"rules"`
    Actions   []FilterAction         `json:"actions"`
    Enabled   bool                   `json:"enabled"`
}

type FilterRule struct {
    Field    string `json:"field"`
    Operator string `json:"operator"`
    Value    string `json:"value"`
}

type FilterAction struct {
    Type   string `json:"type"`
    Value  string `json:"value"`
}

func NewFilterManager() *FilterManager {
    return &FilterManager{
        filters: make(map[string]*EmailFilter),
    }
}
