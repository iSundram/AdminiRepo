
package dns

// DNS record management
type RecordManager struct {
    records map[string]*DNSRecord
}

type DNSRecord struct {
    ID     string `json:"id"`
    ZoneID string `json:"zone_id"`
    Name   string `json:"name"`
    Type   string `json:"type"`
    Value  string `json:"value"`
    TTL    int    `json:"ttl"`
}

func NewRecordManager() *RecordManager {
    return &RecordManager{
        records: make(map[string]*DNSRecord),
    }
}
