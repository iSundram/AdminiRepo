
package models

import (
    "time"
)

type Package struct {
    ID           uint      `json:"id" gorm:"primaryKey"`
    Name         string    `json:"name" gorm:"unique;not null"`
    Description  string    `json:"description"`
    DiskSpace    int64     `json:"disk_space"` // MB
    Bandwidth    int64     `json:"bandwidth"`  // MB
    MaxDomains   int       `json:"max_domains" gorm:"default:1"`
    MaxSubdomains int      `json:"max_subdomains" gorm:"default:10"`
    MaxEmails    int       `json:"max_emails" gorm:"default:10"`
    MaxDatabases int       `json:"max_databases" gorm:"default:1"`
    MaxFTP       int       `json:"max_ftp" gorm:"default:1"`
    Price        float64   `json:"price"`
    IsActive     bool      `json:"is_active" gorm:"default:true"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}
