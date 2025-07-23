
package models

import (
    "time"
)

type Reseller struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    UserID      uint      `json:"user_id" gorm:"unique;not null"`
    User        User      `json:"user" gorm:"foreignKey:UserID"`
    CompanyName string    `json:"company_name"`
    MaxAccounts int       `json:"max_accounts" gorm:"default:10"`
    UsedAccounts int      `json:"used_accounts" gorm:"default:0"`
    Commission  float64   `json:"commission" gorm:"default:10.0"`
    IsActive    bool      `json:"is_active" gorm:"default:true"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
