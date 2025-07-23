
package email

// Email routing management
type RoutingManager struct {
    routes map[string]*EmailRoute
}

type EmailRoute struct {
    ID          string   `json:"id"`
    Domain      string   `json:"domain"`
    Destination string   `json:"destination"`
    Priority    int      `json:"priority"`
    Active      bool     `json:"active"`
}

func NewRoutingManager() *RoutingManager {
    return &RoutingManager{
        routes: make(map[string]*EmailRoute),
    }
}
