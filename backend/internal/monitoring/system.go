
package monitoring

// System monitoring
type SystemMonitor struct {
    metrics map[string]*SystemMetric
}

type SystemMetric struct {
    Name      string  `json:"name"`
    Value     float64 `json:"value"`
    Unit      string  `json:"unit"`
    Timestamp int64   `json:"timestamp"`
}

func NewSystemMonitor() *SystemMonitor {
    return &SystemMonitor{
        metrics: make(map[string]*SystemMetric),
    }
}
