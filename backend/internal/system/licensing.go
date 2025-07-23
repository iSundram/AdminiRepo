
package system

import (
    "crypto/md5"
    "fmt"
    "time"
)

type License struct {
    Key        string    `json:"key"`
    Type       string    `json:"type"`
    ExpiresAt  time.Time `json:"expires_at"`
    MaxUsers   int       `json:"max_users"`
    MaxDomains int       `json:"max_domains"`
    IsValid    bool      `json:"is_valid"`
}

func ValidateLicense(key string) (*License, error) {
    // Demo license validation
    hash := fmt.Sprintf("%x", md5.Sum([]byte(key)))
    
    return &License{
        Key:        key,
        Type:       "Enterprise",
        ExpiresAt:  time.Now().AddDate(1, 0, 0),
        MaxUsers:   1000,
        MaxDomains: 10000,
        IsValid:    len(key) > 10,
    }, nil
}
