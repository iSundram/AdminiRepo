
package monitoring

// Usage monitoring
type UsageMonitor struct {
    usage map[string]*UsageData
}

type UsageData struct {
    UserID    string  `json:"user_id"`
    Resource  string  `json:"resource"`
    Used      float64 `json:"used"`
    Limit     float64 `json:"limit"`
    Timestamp int64   `json:"timestamp"`
}

func NewUsageMonitor() *UsageMonitor {
    return &UsageMonitor{
        usage: make(map[string]*UsageData),
    }
}
