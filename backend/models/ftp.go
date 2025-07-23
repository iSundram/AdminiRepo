
package models

import (
    "time"
)

type FTPAccount struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    AccountID uint      `json:"account_id" gorm:"not null"`
    Account   Account   `json:"account" gorm:"foreignKey:AccountID"`
    Username  string    `json:"username" gorm:"unique;not null"`
    Password  string    `json:"-" gorm:"not null"`
    Directory string    `json:"directory" gorm:"not null"`
    QuotaMB   int       `json:"quota_mb" gorm:"default:1024"`
    IsActive  bool      `json:"is_active" gorm:"default:true"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
