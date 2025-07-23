
package billing

import "time"

// Promotional codes and discounts
type PromoManager struct {
    codes map[string]*PromoCode
}

type PromoCode struct {
    Code        string
    Discount    float64
    Type        string // percentage or fixed
    ExpiresAt   time.Time
    UsageLimit  int
    UsageCount  int
}

func NewPromoManager() *PromoManager {
    return &PromoManager{
        codes: make(map[string]*PromoCode),
    }
}
