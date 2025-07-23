
package backup

import (
    "crypto/md5"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "time"
)

type IncrementalBackup struct {
    BaseDir   string
    LastRun   time.Time
    FileIndex map[string]string // filename -> hash
}

type FileChange struct {
    Path      string `json:"path"`
    Type      string `json:"type"` // added, modified, deleted
    Hash      string `json:"hash"`
    Size      int64  `json:"size"`
    ModTime   time.Time `json:"mod_time"`
}

func NewIncrementalBackup(baseDir string) *IncrementalBackup {
    return &IncrementalBackup{
        BaseDir:   baseDir,
        FileIndex: make(map[string]string),
    }
}

func (ib *IncrementalBackup) ScanChanges() ([]FileChange, error) {
    var changes []FileChange
    
    err := filepath.Walk(ib.BaseDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        
        if info.IsDir() {
            return nil
        }
        
        // Calculate file hash
        hash, err := ib.calculateFileHash(path)
        if err != nil {
            return err
        }
        
        // Check if file changed
        if oldHash, exists := ib.FileIndex[path]; !exists {
            changes = append(changes, FileChange{
                Path:    path,
                Type:    "added",
                Hash:    hash,
                Size:    info.Size(),
                ModTime: info.ModTime(),
            })
        } else if oldHash != hash {
            changes = append(changes, FileChange{
                Path:    path,
                Type:    "modified",
                Hash:    hash,
                Size:    info.Size(),
                ModTime: info.ModTime(),
            })
        }
        
        ib.FileIndex[path] = hash
        return nil
    })
    
    return changes, err
}

func (ib *IncrementalBackup) calculateFileHash(filename string) (string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer file.Close()
    
    hash := md5.New()
    if _, err := io.Copy(hash, file); err != nil {
        return "", err
    }
    
    return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
