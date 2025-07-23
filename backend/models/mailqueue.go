
package models

import (
    "time"
)

type MailQueue struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    FromEmail   string    `json:"from_email" gorm:"not null"`
    ToEmail     string    `json:"to_email" gorm:"not null"`
    Subject     string    `json:"subject" gorm:"not null"`
    Body        string    `json:"body" gorm:"type:text"`
    Status      string    `json:"status" gorm:"default:'pending'"`
    Attempts    int       `json:"attempts" gorm:"default:0"`
    MaxAttempts int       `json:"max_attempts" gorm:"default:3"`
    ScheduledAt time.Time `json:"scheduled_at"`
    SentAt      *time.Time `json:"sent_at"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
