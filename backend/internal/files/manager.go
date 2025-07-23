
package files

// File management system
type FileManager struct {
    files map[string]*File
}

type File struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Path     string `json:"path"`
    Size     int64  `json:"size"`
    MimeType string `json:"mime_type"`
    Owner    string `json:"owner"`
}

func NewFileManager() *FileManager {
    return &FileManager{
        files: make(map[string]*File),
    }
}
