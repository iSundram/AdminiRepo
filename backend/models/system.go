
package models

import (
    "time"
)

type SystemSetting struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Key       string    `json:"key" gorm:"unique;not null"`
    Value     string    `json:"value" gorm:"type:text"`
    Type      string    `json:"type" gorm:"default:'string'"`
    IsPublic  bool      `json:"is_public" gorm:"default:false"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type SystemLog struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Level     string    `json:"level" gorm:"not null"`
    Message   string    `json:"message" gorm:"type:text"`
    Context   string    `json:"context" gorm:"type:text"`
    UserID    *uint     `json:"user_id"`
    User      *User     `json:"user" gorm:"foreignKey:UserID"`
    IPAddress string    `json:"ip_address"`
    CreatedAt time.Time `json:"created_at"`
}
