
package models

import (
    "time"
)

type SSLCertificate struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    DomainID    uint      `json:"domain_id" gorm:"not null"`
    Domain      Domain    `json:"domain" gorm:"foreignKey:DomainID"`
    Type        string    `json:"type" gorm:"default:'letsencrypt'"`
    Certificate string    `json:"certificate" gorm:"type:text"`
    PrivateKey  string    `json:"private_key" gorm:"type:text"`
    IssuedBy    string    `json:"issued_by"`
    IssuedAt    time.Time `json:"issued_at"`
    ExpiresAt   time.Time `json:"expires_at"`
    IsActive    bool      `json:"is_active" gorm:"default:true"`
    AutoRenew   bool      `json:"auto_renew" gorm:"default:true"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
