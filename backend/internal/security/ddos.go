
package security

import (
    "context"
    "log"
    "sync"
    "time"
)

// DDoSProtectionService manages DDoS protection
type DDoSProtectionService struct {
    enabled     bool
    rateLimits  map[string]*RateLimit
    blockedIPs  map[string]*BlockedIP
    mu          sync.RWMutex
    threshold   int
    timeWindow  time.Duration
}

// NewDDoSProtectionService creates a new DDoS protection service
func NewDDoSProtectionService() *DDoSProtectionService {
    return &DDoSProtectionService{
        enabled:    true,
        rateLimits: make(map[string]*RateLimit),
        blockedIPs: make(map[string]*BlockedIP),
        threshold:  100, // requests per minute
        timeWindow: time.Minute,
    }
}

// CheckRateLimit checks if IP is within rate limits
func (d *DDoSProtectionService) CheckRateLimit(ctx context.Context, ip string) (bool, error) {
    d.mu.RLock()
    defer d.mu.RUnlock()
    
    if !d.enabled {
        return true, nil
    }
    
    // Check if IP is blocked
    if _, blocked := d.blockedIPs[ip]; blocked {
        return false, fmt.Errorf("IP is blocked: %s", ip)
    }
    
    rateLimit, exists := d.rateLimits[ip]
    if !exists {
        d.rateLimits[ip] = &RateLimit{
            IP:           ip,
            RequestCount: 1,
            FirstRequest: time.Now(),
            LastRequest:  time.Now(),
        }
        return true, nil
    }
    
    // Reset counter if time window expired
    if time.Since(rateLimit.FirstRequest) > d.timeWindow {
        rateLimit.RequestCount = 1
        rateLimit.FirstRequest = time.Now()
        rateLimit.LastRequest = time.Now()
        return true, nil
    }
    
    rateLimit.RequestCount++
    rateLimit.LastRequest = time.Now()
    
    // Check if threshold exceeded
    if rateLimit.RequestCount > d.threshold {
        d.blockIP(ip, "Rate limit exceeded")
        return false, fmt.Errorf("rate limit exceeded for IP: %s", ip)
    }
    
    return true, nil
}

// blockIP blocks an IP address
func (d *DDoSProtectionService) blockIP(ip string, reason string) {
    d.blockedIPs[ip] = &BlockedIP{
        IP:        ip,
        Reason:    reason,
        Timestamp: time.Now().Format(time.RFC3339),
    }
    log.Printf("Blocked IP due to DDoS: %s (Reason: %s)", ip, reason)
}

// UnblockIP unblocks an IP address
func (d *DDoSProtectionService) UnblockIP(ctx context.Context, ip string) error {
    d.mu.Lock()
    defer d.mu.Unlock()
    
    delete(d.blockedIPs, ip)
    delete(d.rateLimits, ip)
    log.Printf("Unblocked IP: %s", ip)
    return nil
}

// GetDDoSStats returns DDoS protection statistics
func (d *DDoSProtectionService) GetDDoSStats(ctx context.Context) (*DDoSStats, error) {
    d.mu.RLock()
    defer d.mu.RUnlock()
    
    return &DDoSStats{
        Enabled:           d.enabled,
        BlockedIPs:        len(d.blockedIPs),
        RateLimitedIPs:    len(d.rateLimits),
        RequestThreshold:  d.threshold,
        TimeWindow:        d.timeWindow.String(),
        LastUpdate:        time.Now().Format(time.RFC3339),
    }, nil
}

type RateLimit struct {
    IP           string    `json:"ip"`
    RequestCount int       `json:"request_count"`
    FirstRequest time.Time `json:"first_request"`
    LastRequest  time.Time `json:"last_request"`
}

type DDoSStats struct {
    Enabled          bool   `json:"enabled"`
    BlockedIPs       int    `json:"blocked_ips"`
    RateLimitedIPs   int    `json:"rate_limited_ips"`
    RequestThreshold int    `json:"request_threshold"`
    TimeWindow       string `json:"time_window"`
    LastUpdate       string `json:"last_update"`
}
