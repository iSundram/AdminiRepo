
package services

import (
    "fmt"
    "time"
)

// DatabaseManager handles database management
type DatabaseManager struct {
    databases map[string]*Database
    users     map[string]*DatabaseUser
}

// Database represents a database
type Database struct {
    Name      string    `json:"name"`
    Type      string    `json:"type"` // mysql, postgresql, mongodb
    Size      int64     `json:"size"`
    Owner     string    `json:"owner"`
    CreatedAt time.Time `json:"created_at"`
    Status    string    `json:"status"`
    Config    map[string]string `json:"config"`
}

// DatabaseUser represents a database user
type DatabaseUser struct {
    Username    string            `json:"username"`
    Type        string            `json:"type"`
    Databases   []string          `json:"databases"`
    Permissions map[string]string `json:"permissions"`
    CreatedAt   time.Time         `json:"created_at"`
    Status      string            `json:"status"`
}

// NewDatabaseManager creates a new database manager
func NewDatabaseManager() *DatabaseManager {
    return &DatabaseManager{
        databases: make(map[string]*Database),
        users:     make(map[string]*DatabaseUser),
    }
}

// CreateDatabase creates a new database
func (dm *DatabaseManager) CreateDatabase(name, dbType, owner string) (*Database, error) {
    if _, exists := dm.databases[name]; exists {
        return nil, fmt.Errorf("database %s already exists", name)
    }
    
    database := &Database{
        Name:      name,
        Type:      dbType,
        Size:      0,
        Owner:     owner,
        CreatedAt: time.Now(),
        Status:    "active",
        Config:    make(map[string]string),
    }
    
    dm.databases[name] = database
    return database, nil
}

// GetDatabase retrieves a database
func (dm *DatabaseManager) GetDatabase(name string) (*Database, bool) {
    database, exists := dm.databases[name]
    return database, exists
}

// DeleteDatabase deletes a database
func (dm *DatabaseManager) DeleteDatabase(name string) error {
    if _, exists := dm.databases[name]; !exists {
        return fmt.Errorf("database %s not found", name)
    }
    
    // TODO: Implement actual database deletion
    delete(dm.databases, name)
    return nil
}

// CreateUser creates a new database user
func (dm *DatabaseManager) CreateUser(username, userType string) (*DatabaseUser, error) {
    if _, exists := dm.users[username]; exists {
        return nil, fmt.Errorf("user %s already exists", username)
    }
    
    user := &DatabaseUser{
        Username:    username,
        Type:        userType,
        Databases:   make([]string, 0),
        Permissions: make(map[string]string),
        CreatedAt:   time.Now(),
        Status:      "active",
    }
    
    dm.users[username] = user
    return user, nil
}

// GetUser retrieves a database user
func (dm *DatabaseManager) GetUser(username string) (*DatabaseUser, bool) {
    user, exists := dm.users[username]
    return user, exists
}

// GrantAccess grants user access to a database
func (dm *DatabaseManager) GrantAccess(username, dbName, permission string) error {
    user, exists := dm.users[username]
    if !exists {
        return fmt.Errorf("user %s not found", username)
    }
    
    database, exists := dm.databases[dbName]
    if !exists {
        return fmt.Errorf("database %s not found", dbName)
    }
    
    // Add database to user's list if not already present
    found := false
    for _, db := range user.Databases {
        if db == dbName {
            found = true
            break
        }
    }
    if !found {
        user.Databases = append(user.Databases, dbName)
    }
    
    user.Permissions[dbName] = permission
    _ = database // Used to verify database exists
    
    return nil
}

// ListDatabases returns all databases
func (dm *DatabaseManager) ListDatabases() map[string]*Database {
    return dm.databases
}

// ListUsers returns all database users
func (dm *DatabaseManager) ListUsers() map[string]*DatabaseUser {
    return dm.users
}
