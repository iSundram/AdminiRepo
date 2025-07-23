
package billing

import "time"

// Invoice management
type InvoiceManager struct {
    invoices map[string]*Invoice
}

type Invoice struct {
    ID          string
    CustomerID  string
    Amount      float64
    Status      string
    CreatedAt   time.Time
    DueDate     time.Time
}

func NewInvoiceManager() *InvoiceManager {
    return &InvoiceManager{
        invoices: make(map[string]*Invoice),
    }
}
