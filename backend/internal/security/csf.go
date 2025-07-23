
package security

import (
    "context"
    "log"
    "net"
)

// CSFService manages ConfigServer Firewall
type CSFService struct {
    enabled     bool
    blockedIPs  map[string]*BlockedIP
    allowedIPs  map[string]*AllowedIP
}

// NewCSFService creates a new CSF service
func NewCSFService() *CSFService {
    return &CSFService{
        enabled:    true,
        blockedIPs: make(map[string]*BlockedIP),
        allowedIPs: make(map[string]*AllowedIP),
    }
}

// BlockIP blocks an IP address
func (c *CSFService) BlockIP(ctx context.Context, ip string, reason string) error {
    if net.ParseIP(ip) == nil {
        return fmt.Errorf("invalid IP address: %s", ip)
    }
    
    c.blockedIPs[ip] = &BlockedIP{
        IP:        ip,
        Reason:    reason,
        Timestamp: "2024-01-01T00:00:00Z",
    }
    
    log.Printf("Blocked IP: %s (Reason: %s)", ip, reason)
    return nil
}

// UnblockIP unblocks an IP address
func (c *CSFService) UnblockIP(ctx context.Context, ip string) error {
    delete(c.blockedIPs, ip)
    log.Printf("Unblocked IP: %s", ip)
    return nil
}

// AllowIP adds IP to whitelist
func (c *CSFService) AllowIP(ctx context.Context, ip string, description string) error {
    if net.ParseIP(ip) == nil {
        return fmt.Errorf("invalid IP address: %s", ip)
    }
    
    c.allowedIPs[ip] = &AllowedIP{
        IP:          ip,
        Description: description,
        Timestamp:   "2024-01-01T00:00:00Z",
    }
    
    log.Printf("Allowed IP: %s", ip)
    return nil
}

// GetBlockedIPs returns list of blocked IPs
func (c *CSFService) GetBlockedIPs(ctx context.Context) (map[string]*BlockedIP, error) {
    return c.blockedIPs, nil
}

type BlockedIP struct {
    IP        string `json:"ip"`
    Reason    string `json:"reason"`
    Timestamp string `json:"timestamp"`
}

type AllowedIP struct {
    IP          string `json:"ip"`
    Description string `json:"description"`
    Timestamp   string `json:"timestamp"`
}
