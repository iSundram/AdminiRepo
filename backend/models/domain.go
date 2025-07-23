
package models

import (
    "time"
)

type Domain struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    AccountID uint      `json:"account_id" gorm:"not null"`
    Account   Account   `json:"account" gorm:"foreignKey:AccountID"`
    Name      string    `json:"name" gorm:"unique;not null"`
    Type      string    `json:"type" gorm:"default:'addon'"`
    IsMain    bool      `json:"is_main" gorm:"default:false"`
    Status    string    `json:"status" gorm:"default:'active'"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
