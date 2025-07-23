
package dns

// DNS zone management
type ZoneManager struct {
    zones map[string]*DNSZone
}

type DNSZone struct {
    ID       string `json:"id"`
    Domain   string `json:"domain"`
    Status   string `json:"status"`
    Serial   int    `json:"serial"`
}

func NewZoneManager() *ZoneManager {
    return &ZoneManager{
        zones: make(map[string]*DNSZone),
    }
}
