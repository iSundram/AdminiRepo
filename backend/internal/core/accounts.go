
package core

// Account management
type AccountManager struct {
    accounts map[string]*Account
}

type Account struct {
    ID       string `json:"id"`
    UserID   string `json:"user_id"`
    Domain   string `json:"domain"`
    Package  string `json:"package"`
    Status   string `json:"status"`
}

func NewAccountManager() *AccountManager {
    return &AccountManager{
        accounts: make(map[string]*Account),
    }
}
