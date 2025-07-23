
package integrations

import (
    "bytes"
    "encoding/json"
    "net/http"
)

type MailChannels struct {
    APIKey  string
    BaseURL string
}

type EmailMessage struct {
    From    string   `json:"from"`
    To      []string `json:"to"`
    Subject string   `json:"subject"`
    Body    string   `json:"body"`
    HTML    bool     `json:"html"`
}

type DeliveryStatus struct {
    MessageID string `json:"message_id"`
    Status    string `json:"status"`
    Delivered bool   `json:"delivered"`
}

func NewMailChannels(apiKey string) *MailChannels {
    return &MailChannels{
        APIKey:  apiKey,
        BaseURL: "https://api.mailchannels.net",
    }
}

func (mc *MailChannels) SendEmail(message EmailMessage) (*DeliveryStatus, error) {
    jsonData, _ := json.Marshal(message)
    
    req, err := http.NewRequest("POST", mc.BaseURL+"/tx/v1/send", bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, err
    }
    
    req.Header.Set("Authorization", "Bearer "+mc.APIKey)
    req.Header.Set("Content-Type", "application/json")
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    var status DeliveryStatus
    return &status, json.NewDecoder(resp.Body).Decode(&status)
}
