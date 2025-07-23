
package models

import (
    "time"
)

type DNSRecord struct {
    ID       uint      `json:"id" gorm:"primaryKey"`
    DomainID uint      `json:"domain_id" gorm:"not null"`
    Domain   Domain    `json:"domain" gorm:"foreignKey:DomainID"`
    Name     string    `json:"name" gorm:"not null"`
    Type     string    `json:"type" gorm:"not null"`
    Value    string    `json:"value" gorm:"not null"`
    TTL      int       `json:"ttl" gorm:"default:3600"`
    Priority int       `json:"priority" gorm:"default:0"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
