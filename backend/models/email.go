
package models

import (
    "time"
)

type EmailAccount struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    DomainID  uint      `json:"domain_id" gorm:"not null"`
    Domain    Domain    `json:"domain" gorm:"foreignKey:DomainID"`
    Username  string    `json:"username" gorm:"not null"`
    Password  string    `json:"-" gorm:"not null"`
    QuotaMB   int       `json:"quota_mb" gorm:"default:1024"`
    UsedMB    int       `json:"used_mb" gorm:"default:0"`
    IsActive  bool      `json:"is_active" gorm:"default:true"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
