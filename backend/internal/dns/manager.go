
package dns

// DNS management
type DNSManager struct {
    zones map[string]*Zone
}

type Zone struct {
    ID      string   `json:"id"`
    Name    string   `json:"name"`
    Records []Record `json:"records"`
    TTL     int      `json:"ttl"`
}

type Record struct {
    Name  string `json:"name"`
    Type  string `json:"type"`
    Value string `json:"value"`
    TTL   int    `json:"ttl"`
}

func NewDNSManager() *DNSManager {
    return &DNSManager{
        zones: make(map[string]*Zone),
    }
}
