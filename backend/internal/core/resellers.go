
package core

// Reseller management
type ResellerManager struct {
    resellers map[string]*Reseller
}

type Reseller struct {
    ID          string  `json:"id"`
    UserID      string  `json:"user_id"`
    CompanyName string  `json:"company_name"`
    Commission  float64 `json:"commission"`
    Active      bool    `json:"active"`
}

func NewResellerManager() *ResellerManager {
    return &ResellerManager{
        resellers: make(map[string]*Reseller),
    }
}
