
package billing

// Payment processing
type PaymentProcessor struct {
    gateways map[string]PaymentGateway
}

type PaymentGateway interface {
    ProcessPayment(amount float64, currency string) (*PaymentResult, error)
}

type PaymentResult struct {
    TransactionID string
    Status        string
    Amount        float64
}

func NewPaymentProcessor() *PaymentProcessor {
    return &PaymentProcessor{
        gateways: make(map[string]PaymentGateway),
    }
}
