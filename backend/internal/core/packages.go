
package core

// Package management
type PackageManager struct {
    packages map[string]*Package
}

type Package struct {
    ID          string         `json:"id"`
    Name        string         `json:"name"`
    Price       float64        `json:"price"`
    Resources   map[string]int `json:"resources"`
    Features    []string       `json:"features"`
}

func NewPackageManager() *PackageManager {
    return &PackageManager{
        packages: make(map[string]*Package),
    }
}
