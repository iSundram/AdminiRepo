
package monitoring

// Anomaly detection for monitoring
type AnomalyMonitor struct {
    anomalies map[string]*Anomaly
}

type Anomaly struct {
    ID          string  `json:"id"`
    Metric      string  `json:"metric"`
    Value       float64 `json:"value"`
    Expected    float64 `json:"expected"`
    Confidence  float64 `json:"confidence"`
    Severity    string  `json:"severity"`
}

func NewAnomalyMonitor() *AnomalyMonitor {
    return &AnomalyMonitor{
        anomalies: make(map[string]*Anomaly),
    }
}
