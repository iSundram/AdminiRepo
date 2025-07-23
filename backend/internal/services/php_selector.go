
package services

import (
    "fmt"
    "sort"
)

// PHPSelectorManager handles PHP version management
type PHPSelectorManager struct {
    versions     map[string]*PHPVersion
    domains      map[string]string // domain -> php version
    defaultVersion string
}

// PHPVersion represents a PHP version configuration
type PHPVersion struct {
    Version     string            `json:"version"`
    Path        string            `json:"path"`
    Modules     []string          `json:"modules"`
    Config      map[string]string `json:"config"`
    Status      string            `json:"status"`
    Description string            `json:"description"`
}

// NewPHPSelectorManager creates a new PHP selector manager
func NewPHPSelectorManager() *PHPSelectorManager {
    manager := &PHPSelectorManager{
        versions: make(map[string]*PHPVersion),
        domains:  make(map[string]string),
        defaultVersion: "8.1",
    }
    
    // Initialize with common PHP versions
    manager.initializeVersions()
    return manager
}

func (psm *PHPSelectorManager) initializeVersions() {
    versions := []PHPVersion{
        {
            Version:     "7.4",
            Path:        "/usr/bin/php74",
            Modules:     []string{"mysql", "gd", "curl", "zip", "xml"},
            Config:      make(map[string]string),
            Status:      "available",
            Description: "PHP 7.4 LTS",
        },
        {
            Version:     "8.0",
            Path:        "/usr/bin/php80",
            Modules:     []string{"mysql", "gd", "curl", "zip", "xml", "opcache"},
            Config:      make(map[string]string),
            Status:      "available",
            Description: "PHP 8.0",
        },
        {
            Version:     "8.1",
            Path:        "/usr/bin/php81",
            Modules:     []string{"mysql", "gd", "curl", "zip", "xml", "opcache"},
            Config:      make(map[string]string),
            Status:      "available",
            Description: "PHP 8.1 LTS",
        },
        {
            Version:     "8.2",
            Path:        "/usr/bin/php82",
            Modules:     []string{"mysql", "gd", "curl", "zip", "xml", "opcache"},
            Config:      make(map[string]string),
            Status:      "available",
            Description: "PHP 8.2",
        },
    }
    
    for _, version := range versions {
        psm.versions[version.Version] = &version
    }
}

// GetAvailableVersions returns all available PHP versions
func (psm *PHPSelectorManager) GetAvailableVersions() []*PHPVersion {
    versions := make([]*PHPVersion, 0, len(psm.versions))
    for _, version := range psm.versions {
        versions = append(versions, version)
    }
    
    sort.Slice(versions, func(i, j int) bool {
        return versions[i].Version < versions[j].Version
    })
    
    return versions
}

// SetDomainPHPVersion sets PHP version for a domain
func (psm *PHPSelectorManager) SetDomainPHPVersion(domain, version string) error {
    if _, exists := psm.versions[version]; !exists {
        return fmt.Errorf("PHP version %s not available", version)
    }
    
    psm.domains[domain] = version
    return nil
}

// GetDomainPHPVersion gets PHP version for a domain
func (psm *PHPSelectorManager) GetDomainPHPVersion(domain string) string {
    if version, exists := psm.domains[domain]; exists {
        return version
    }
    return psm.defaultVersion
}

// InstallModule installs a PHP module for a version
func (psm *PHPSelectorManager) InstallModule(version, module string) error {
    phpVersion, exists := psm.versions[version]
    if !exists {
        return fmt.Errorf("PHP version %s not found", version)
    }
    
    // Check if module is already installed
    for _, installedModule := range phpVersion.Modules {
        if installedModule == module {
            return nil // Already installed
        }
    }
    
    // TODO: Implement actual module installation
    phpVersion.Modules = append(phpVersion.Modules, module)
    return nil
}

// UpdateConfig updates PHP configuration for a version
func (psm *PHPSelectorManager) UpdateConfig(version string, config map[string]string) error {
    phpVersion, exists := psm.versions[version]
    if !exists {
        return fmt.Errorf("PHP version %s not found", version)
    }
    
    for key, value := range config {
        phpVersion.Config[key] = value
    }
    
    return nil
}
