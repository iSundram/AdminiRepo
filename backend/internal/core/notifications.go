
package core

import "time"

// Notification management
type NotificationManager struct {
    notifications map[string]*Notification
}

type Notification struct {
    ID        string    `json:"id"`
    UserID    string    `json:"user_id"`
    Title     string    `json:"title"`
    Message   string    `json:"message"`
    Type      string    `json:"type"`
    Read      bool      `json:"read"`
    CreatedAt time.Time `json:"created_at"`
}

func NewNotificationManager() *NotificationManager {
    return &NotificationManager{
        notifications: make(map[string]*Notification),
    }
}
