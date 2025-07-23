
package models

import (
    "time"
)

type Database struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    AccountID uint      `json:"account_id" gorm:"not null"`
    Account   Account   `json:"account" gorm:"foreignKey:AccountID"`
    Name      string    `json:"name" gorm:"unique;not null"`
    Type      string    `json:"type" gorm:"default:'mysql'"`
    SizeMB    int64     `json:"size_mb" gorm:"default:0"`
    IsActive  bool      `json:"is_active" gorm:"default:true"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type DatabaseUser struct {
    ID         uint      `json:"id" gorm:"primaryKey"`
    DatabaseID uint      `json:"database_id" gorm:"not null"`
    Database   Database  `json:"database" gorm:"foreignKey:DatabaseID"`
    Username   string    `json:"username" gorm:"not null"`
    Password   string    `json:"-" gorm:"not null"`
    Privileges string    `json:"privileges" gorm:"default:'ALL'"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}
