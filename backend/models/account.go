
package models

import (
    "time"
)

type Account struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    UserID      uint      `json:"user_id" gorm:"not null"`
    User        User      `json:"user" gorm:"foreignKey:UserID"`
    PackageID   uint      `json:"package_id" gorm:"not null"`
    Package     Package   `json:"package" gorm:"foreignKey:PackageID"`
    DomainName  string    `json:"domain_name" gorm:"unique;not null"`
    Status      string    `json:"status" gorm:"default:'active'"`
    SuspendedAt *time.Time `json:"suspended_at"`
    ExpiresAt   time.Time `json:"expires_at"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
