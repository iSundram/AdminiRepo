
package files

// WebDisk management
type WebDiskManager struct {
    webdisks map[string]*WebDisk
}

type WebDisk struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Path     string `json:"path"`
    Size     int64  `json:"size"`
    Owner    string `json:"owner"`
}

func NewWebDiskManager() *WebDiskManager {
    return &WebDiskManager{
        webdisks: make(map[string]*WebDisk),
    }
}
