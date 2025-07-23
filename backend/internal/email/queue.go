
package email

// Email queue management
type EmailQueue struct {
    queue []EmailMessage
}

type EmailMessage struct {
    ID        string `json:"id"`
    To        string `json:"to"`
    From      string `json:"from"`
    Subject   string `json:"subject"`
    Body      string `json:"body"`
    Status    string `json:"status"`
    Priority  int    `json:"priority"`
}

func NewEmailQueue() *EmailQueue {
    return &EmailQueue{
        queue: make([]EmailMessage, 0),
    }
}
