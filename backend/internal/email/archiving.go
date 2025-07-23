
package email

import "time"

// Email archiving system
type ArchiveManager struct {
    archives map[string]*EmailArchive
}

type EmailArchive struct {
    ID        string    `json:"id"`
    UserID    string    `json:"user_id"`
    Messages  []string  `json:"messages"`
    CreatedAt time.Time `json:"created_at"`
    Size      int64     `json:"size"`
}

func NewArchiveManager() *ArchiveManager {
    return &ArchiveManager{
        archives: make(map[string]*EmailArchive),
    }
}
