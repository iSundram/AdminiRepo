
package models

import (
    "time"
)

type Backup struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    AccountID   uint      `json:"account_id" gorm:"not null"`
    Account     Account   `json:"account" gorm:"foreignKey:AccountID"`
    Name        string    `json:"name" gorm:"not null"`
    Type        string    `json:"type" gorm:"not null"` // full, incremental, files, database
    SizeMB      int64     `json:"size_mb"`
    Path        string    `json:"path"`
    RemotePath  string    `json:"remote_path"`
    Status      string    `json:"status" gorm:"default:'pending'"`
    Progress    int       `json:"progress" gorm:"default:0"`
    StartedAt   *time.Time `json:"started_at"`
    CompletedAt *time.Time `json:"completed_at"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
