
package monitoring

import "time"

// Alert management
type AlertManager struct {
    alerts map[string]*Alert
}

type Alert struct {
    ID          string    `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Severity    string    `json:"severity"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"created_at"`
}

func NewAlertManager() *AlertManager {
    return &AlertManager{
        alerts: make(map[string]*Alert),
    }
}
