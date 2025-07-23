
package models

import (
    "time"
)

type Plugin struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Name        string    `json:"name" gorm:"unique;not null"`
    Version     string    `json:"version" gorm:"not null"`
    Description string    `json:"description" gorm:"type:text"`
    Author      string    `json:"author"`
    Path        string    `json:"path" gorm:"not null"`
    Config      string    `json:"config" gorm:"type:text"`
    IsEnabled   bool      `json:"is_enabled" gorm:"default:false"`
    IsInstalled bool      `json:"is_installed" gorm:"default:false"`
    InstalledAt *time.Time `json:"installed_at"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
