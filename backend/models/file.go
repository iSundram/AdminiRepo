
package models

import (
    "time"
)

type File struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    AccountID uint      `json:"account_id" gorm:"not null"`
    Account   Account   `json:"account" gorm:"foreignKey:AccountID"`
    Name      string    `json:"name" gorm:"not null"`
    Path      string    `json:"path" gorm:"not null"`
    Size      int64     `json:"size"`
    MimeType  string    `json:"mime_type"`
    Hash      string    `json:"hash"`
    IsPublic  bool      `json:"is_public" gorm:"default:false"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
