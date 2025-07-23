
package core

// User management
type UserManager struct {
    users map[string]*User
}

type User struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Role     string `json:"role"`
    Active   bool   `json:"active"`
}

func NewUserManager() *UserManager {
    return &UserManager{
        users: make(map[string]*User),
    }
}
