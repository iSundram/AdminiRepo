
package ssl

import (
    "crypto/x509"
    "time"
)

// TransparencyManager handles Certificate Transparency
type TransparencyManager struct {
    logs     map[string]*CTLog
    entries  map[string]*CTEntry
    monitors map[string]*CTMonitor
}

// CTLog represents a Certificate Transparency log
type CTLog struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    URL         string    `json:"url"`
    PublicKey   string    `json:"public_key"`
    Status      string    `json:"status"`
    LastSync    time.Time `json:"last_sync"`
    EntryCount  int64     `json:"entry_count"`
}

// CTEntry represents a CT log entry
type CTEntry struct {
    ID          string           `json:"id"`
    LogID       string           `json:"log_id"`
    Index       int64            `json:"index"`
    Timestamp   time.Time        `json:"timestamp"`
    Certificate *x509.Certificate `json:"certificate"`
    Domain      string           `json:"domain"`
    Issuer      string           `json:"issuer"`
    Status      string           `json:"status"`
}

// CTMonitor represents a CT monitoring configuration
type CTMonitor struct {
    ID       string   `json:"id"`
    Domain   string   `json:"domain"`
    Alerts   bool     `json:"alerts"`
    Webhooks []string `json:"webhooks"`
    Created  time.Time `json:"created"`
}

// NewTransparencyManager creates a new transparency manager
func NewTransparencyManager() *TransparencyManager {
    return &TransparencyManager{
        logs:     make(map[string]*CTLog),
        entries:  make(map[string]*CTEntry),
        monitors: make(map[string]*CTMonitor),
    }
}

// AddLog adds a CT log
func (tm *TransparencyManager) AddLog(name, url, publicKey string) (*CTLog, error) {
    logID := generateLogID(name)
    
    log := &CTLog{
        ID:         logID,
        Name:       name,
        URL:        url,
        PublicKey:  publicKey,
        Status:     "active",
        LastSync:   time.Now(),
        EntryCount: 0,
    }
    
    tm.logs[logID] = log
    return log, nil
}

// MonitorDomain adds domain monitoring
func (tm *TransparencyManager) MonitorDomain(domain string, alerts bool) (*CTMonitor, error) {
    monitorID := generateMonitorID(domain)
    
    monitor := &CTMonitor{
        ID:       monitorID,
        Domain:   domain,
        Alerts:   alerts,
        Webhooks: make([]string, 0),
        Created:  time.Now(),
    }
    
    tm.monitors[monitorID] = monitor
    return monitor, nil
}

// SearchEntries searches CT entries for a domain
func (tm *TransparencyManager) SearchEntries(domain string) ([]*CTEntry, error) {
    var entries []*CTEntry
    
    for _, entry := range tm.entries {
        if entry.Domain == domain {
            entries = append(entries, entry)
        }
    }
    
    return entries, nil
}

func generateLogID(name string) string {
    return "log_" + name
}

func generateMonitorID(domain string) string {
    return "monitor_" + domain
}
