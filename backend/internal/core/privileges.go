
package core

// Privilege management
type PrivilegeManager struct {
    privileges map[string]*Privilege
}

type Privilege struct {
    ID          string   `json:"id"`
    Name        string   `json:"name"`
    Description string   `json:"description"`
    Permissions []string `json:"permissions"`
}

func NewPrivilegeManager() *PrivilegeManager {
    return &PrivilegeManager{
        privileges: make(map[string]*Privilege),
    }
}
