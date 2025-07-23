
package backup

import (
    "archive/tar"
    "compress/gzip"
    "io"
    "log"
    "os"
    "path/filepath"
)

type RestoreManager struct {
    BackupPath string
    TargetPath string
}

type RestorePoint struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"`
    Size      int64     `json:"size"`
    Type      string    `json:"type"`
    Path      string    `json:"path"`
}

func NewRestoreManager(backupPath, targetPath string) *RestoreManager {
    return &RestoreManager{
        BackupPath: backupPath,
        TargetPath: targetPath,
    }
}

func (rm *RestoreManager) ListRestorePoints() ([]RestorePoint, error) {
    var points []RestorePoint
    
    err := filepath.Walk(rm.BackupPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        
        if filepath.Ext(path) == ".tar.gz" {
            points = append(points, RestorePoint{
                ID:        filepath.Base(path),
                Name:      info.Name(),
                CreatedAt: info.ModTime(),
                Size:      info.Size(),
                Type:      "full",
                Path:      path,
            })
        }
        
        return nil
    })
    
    return points, err
}

func (rm *RestoreManager) RestoreFromPoint(pointID string) error {
    log.Printf("Starting restore from point: %s", pointID)
    
    backupFile := filepath.Join(rm.BackupPath, pointID)
    return rm.extractTarGz(backupFile, rm.TargetPath)
}

func (rm *RestoreManager) extractTarGz(src, dest string) error {
    file, err := os.Open(src)
    if err != nil {
        return err
    }
    defer file.Close()
    
    gzr, err := gzip.NewReader(file)
    if err != nil {
        return err
    }
    defer gzr.Close()
    
    tr := tar.NewReader(gzr)
    
    for {
        header, err := tr.Next()
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }
        
        target := filepath.Join(dest, header.Name)
        
        switch header.Typeflag {
        case tar.TypeDir:
            if err := os.MkdirAll(target, 0755); err != nil {
                return err
            }
        case tar.TypeReg:
            f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
            if err != nil {
                return err
            }
            if _, err := io.Copy(f, tr); err != nil {
                f.Close()
                return err
            }
            f.Close()
        }
    }
    
    return nil
}
