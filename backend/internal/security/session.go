
package security

import (
    "crypto/rand"
    "encoding/hex"
    "sync"
    "time"
)

// SessionManager handles secure session management
type SessionManager struct {
    sessions map[string]*Session
    mutex    sync.RWMutex
    timeout  time.Duration
}

// Session represents a user session
type Session struct {
    ID        string    `json:"id"`
    UserID    string    `json:"user_id"`
    IPAddress string    `json:"ip_address"`
    UserAgent string    `json:"user_agent"`
    CreatedAt time.Time `json:"created_at"`
    LastSeen  time.Time `json:"last_seen"`
    Active    bool      `json:"active"`
}

// NewSessionManager creates a new session manager
func NewSessionManager(timeout time.Duration) *SessionManager {
    return &SessionManager{
        sessions: make(map[string]*Session),
        timeout:  timeout,
    }
}

// CreateSession creates a new session
func (sm *SessionManager) CreateSession(userID, ipAddress, userAgent string) (*Session, error) {
    sessionID, err := generateSessionID()
    if err != nil {
        return nil, err
    }
    
    session := &Session{
        ID:        sessionID,
        UserID:    userID,
        IPAddress: ipAddress,
        UserAgent: userAgent,
        CreatedAt: time.Now(),
        LastSeen:  time.Now(),
        Active:    true,
    }
    
    sm.mutex.Lock()
    sm.sessions[sessionID] = session
    sm.mutex.Unlock()
    
    return session, nil
}

// GetSession retrieves a session by ID
func (sm *SessionManager) GetSession(sessionID string) (*Session, bool) {
    sm.mutex.RLock()
    session, exists := sm.sessions[sessionID]
    sm.mutex.RUnlock()
    
    if !exists || !session.Active {
        return nil, false
    }
    
    if time.Since(session.LastSeen) > sm.timeout {
        sm.InvalidateSession(sessionID)
        return nil, false
    }
    
    session.LastSeen = time.Now()
    return session, true
}

// InvalidateSession invalidates a session
func (sm *SessionManager) InvalidateSession(sessionID string) {
    sm.mutex.Lock()
    if session, exists := sm.sessions[sessionID]; exists {
        session.Active = false
    }
    sm.mutex.Unlock()
}

// CleanupExpiredSessions removes expired sessions
func (sm *SessionManager) CleanupExpiredSessions() {
    sm.mutex.Lock()
    defer sm.mutex.Unlock()
    
    for id, session := range sm.sessions {
        if time.Since(session.LastSeen) > sm.timeout {
            delete(sm.sessions, id)
        }
    }
}

func generateSessionID() (string, error) {
    bytes := make([]byte, 32)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}
