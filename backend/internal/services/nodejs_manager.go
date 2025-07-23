
package services

import (
    "fmt"
    "sort"
)

// NodeJSManager handles Node.js version management
type NodeJSManager struct {
    versions map[string]*NodeJSVersion
    apps     map[string]*NodeJSApp
}

// NodeJSVersion represents a Node.js version
type NodeJSVersion struct {
    Version     string   `json:"version"`
    Path        string   `json:"path"`
    Status      string   `json:"status"`
    NPMVersion  string   `json:"npm_version"`
    Modules     []string `json:"modules"`
    Description string   `json:"description"`
}

// NodeJSApp represents a Node.js application
type NodeJSApp struct {
    ID          string            `json:"id"`
    Name        string            `json:"name"`
    Domain      string            `json:"domain"`
    Path        string            `json:"path"`
    Version     string            `json:"version"`
    Port        int               `json:"port"`
    Status      string            `json:"status"`
    Environment map[string]string `json:"environment"`
    StartScript string            `json:"start_script"`
}

// NewNodeJSManager creates a new Node.js manager
func NewNodeJSManager() *NodeJSManager {
    manager := &NodeJSManager{
        versions: make(map[string]*NodeJSVersion),
        apps:     make(map[string]*NodeJSApp),
    }
    
    manager.initializeVersions()
    return manager
}

func (njm *NodeJSManager) initializeVersions() {
    versions := []NodeJSVersion{
        {
            Version:     "16.20.0",
            Path:        "/usr/bin/node16",
            Status:      "available",
            NPMVersion:  "8.19.4",
            Modules:     []string{},
            Description: "Node.js 16 LTS",
        },
        {
            Version:     "18.17.0",
            Path:        "/usr/bin/node18",
            Status:      "available",
            NPMVersion:  "9.6.7",
            Modules:     []string{},
            Description: "Node.js 18 LTS",
        },
        {
            Version:     "20.5.0",
            Path:        "/usr/bin/node20",
            Status:      "available",
            NPMVersion:  "9.8.0",
            Modules:     []string{},
            Description: "Node.js 20 LTS",
        },
    }
    
    for _, version := range versions {
        njm.versions[version.Version] = &version
    }
}

// GetAvailableVersions returns all available Node.js versions
func (njm *NodeJSManager) GetAvailableVersions() []*NodeJSVersion {
    versions := make([]*NodeJSVersion, 0, len(njm.versions))
    for _, version := range njm.versions {
        versions = append(versions, version)
    }
    
    sort.Slice(versions, func(i, j int) bool {
        return versions[i].Version < versions[j].Version
    })
    
    return versions
}

// CreateApp creates a new Node.js application
func (njm *NodeJSManager) CreateApp(name, domain, path, version string, port int) (*NodeJSApp, error) {
    if _, exists := njm.versions[version]; !exists {
        return nil, fmt.Errorf("Node.js version %s not available", version)
    }
    
    appID := fmt.Sprintf("nodejs_%s_%d", domain, port)
    
    app := &NodeJSApp{
        ID:          appID,
        Name:        name,
        Domain:      domain,
        Path:        path,
        Version:     version,
        Port:        port,
        Status:      "created",
        Environment: make(map[string]string),
        StartScript: "npm start",
    }
    
    njm.apps[appID] = app
    return app, nil
}

// GetApp retrieves a Node.js application
func (njm *NodeJSManager) GetApp(appID string) (*NodeJSApp, bool) {
    app, exists := njm.apps[appID]
    return app, exists
}

// StartApp starts a Node.js application
func (njm *NodeJSManager) StartApp(appID string) error {
    app, exists := njm.apps[appID]
    if !exists {
        return fmt.Errorf("app not found: %s", appID)
    }
    
    // TODO: Implement actual app starting logic
    app.Status = "running"
    return nil
}

// StopApp stops a Node.js application
func (njm *NodeJSManager) StopApp(appID string) error {
    app, exists := njm.apps[appID]
    if !exists {
        return fmt.Errorf("app not found: %s", appID)
    }
    
    // TODO: Implement actual app stopping logic
    app.Status = "stopped"
    return nil
}

// ListApps returns all Node.js applications
func (njm *NodeJSManager) ListApps() map[string]*NodeJSApp {
    return njm.apps
}
