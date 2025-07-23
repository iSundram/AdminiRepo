
package models

import (
    "time"
)

type APIToken struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    UserID      uint      `json:"user_id" gorm:"not null"`
    User        User      `json:"user" gorm:"foreignKey:UserID"`
    Name        string    `json:"name" gorm:"not null"`
    Token       string    `json:"token" gorm:"unique;not null"`
    Permissions string    `json:"permissions" gorm:"type:text"`
    LastUsedAt  *time.Time `json:"last_used_at"`
    ExpiresAt   *time.Time `json:"expires_at"`
    IsActive    bool      `json:"is_active" gorm:"default:true"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
