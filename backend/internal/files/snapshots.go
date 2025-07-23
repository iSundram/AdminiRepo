
package files

import "time"

// File snapshot management
type SnapshotManager struct {
    snapshots map[string]*Snapshot
}

type Snapshot struct {
    ID        string    `json:"id"`
    Path      string    `json:"path"`
    Size      int64     `json:"size"`
    CreatedAt time.Time `json:"created_at"`
    Type      string    `json:"type"`
}

func NewSnapshotManager() *SnapshotManager {
    return &SnapshotManager{
        snapshots: make(map[string]*Snapshot),
    }
}
