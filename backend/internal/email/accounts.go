
package email

// Email account management
type EmailAccountManager struct {
    accounts map[string]*EmailAccount
}

type EmailAccount struct {
    ID       string `json:"id"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Quota    int    `json:"quota"`
    Used     int    `json:"used"`
}

func NewEmailAccountManager() *EmailAccountManager {
    return &EmailAccountManager{
        accounts: make(map[string]*EmailAccount),
    }
}
