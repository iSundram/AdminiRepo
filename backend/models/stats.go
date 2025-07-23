
package models

import (
    "time"
)

type UsageStats struct {
    ID            uint      `json:"id" gorm:"primaryKey"`
    AccountID     uint      `json:"account_id" gorm:"not null"`
    Account       Account   `json:"account" gorm:"foreignKey:AccountID"`
    Date          time.Time `json:"date" gorm:"type:date"`
    DiskUsageMB   int64     `json:"disk_usage_mb"`
    BandwidthMB   int64     `json:"bandwidth_mb"`
    EmailsSent    int       `json:"emails_sent"`
    EmailsReceived int      `json:"emails_received"`
    Visitors      int       `json:"visitors"`
    PageViews     int       `json:"page_views"`
    CreatedAt     time.Time `json:"created_at"`
}
