
package models

import (
    "time"
)

type User struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Username    string    `json:"username" gorm:"unique;not null"`
    Email       string    `json:"email" gorm:"unique;not null"`
    Password    string    `json:"-" gorm:"not null"`
    FirstName   string    `json:"first_name"`
    LastName    string    `json:"last_name"`
    Role        string    `json:"role" gorm:"default:'user'"`
    IsActive    bool      `json:"is_active" gorm:"default:true"`
    LastLogin   *time.Time `json:"last_login"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    DeletedAt   *time.Time `json:"deleted_at" gorm:"index"`
}
