
package system

import (
    "os/exec"
    "sync"
)

type Terminal struct {
    ID      string `json:"id"`
    UserID  int    `json:"user_id"`
    Active  bool   `json:"active"`
    Command *exec.Cmd
    mutex   sync.Mutex
}

func NewTerminal(userID int) *Terminal {
    return &Terminal{
        ID:     generateID(),
        UserID: userID,
        Active: false,
    }
}

func (t *Terminal) ExecuteCommand(command string) ([]byte, error) {
    t.mutex.Lock()
    defer t.mutex.Unlock()
    
    cmd := exec.Command("sh", "-c", command)
    return cmd.Output()
}

func generateID() string {
    return "term_" + string(rune(time.Now().Unix()))
}
