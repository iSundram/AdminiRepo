
package core

// Permission delegation management
type DelegationManager struct {
    delegations map[string]*Delegation
}

type Delegation struct {
    ID          string   `json:"id"`
    GranterID   string   `json:"granter_id"`
    GranteeID   string   `json:"grantee_id"`
    Permissions []string `json:"permissions"`
    Active      bool     `json:"active"`
}

func NewDelegationManager() *DelegationManager {
    return &DelegationManager{
        delegations: make(map[string]*Delegation),
    }
}
