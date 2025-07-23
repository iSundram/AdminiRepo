
package models

import (
    "time"
)

type Notification struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    UserID    uint      `json:"user_id" gorm:"not null"`
    User      User      `json:"user" gorm:"foreignKey:UserID"`
    Type      string    `json:"type" gorm:"not null"`
    Title     string    `json:"title" gorm:"not null"`
    Message   string    `json:"message" gorm:"type:text"`
    IsRead    bool      `json:"is_read" gorm:"default:false"`
    ReadAt    *time.Time `json:"read_at"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
