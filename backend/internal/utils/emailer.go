
package utils

import (
    "log"
    "net/smtp"
)

type Emailer struct {
    smtpHost string
    smtpPort string
    username string
    password string
}

func NewEmailer(host, port, username, password string) *Emailer {
    return &Emailer{
        smtpHost: host,
        smtpPort: port,
        username: username,
        password: password,
    }
}

func (e *Emailer) SendEmail(to, subject, body string) error {
    log.Printf("Sending email to: %s", to)
    
    auth := smtp.PlainAuth("", e.username, e.password, e.smtpHost)
    
    msg := []byte("To: " + to + "\r\n" +
        "Subject: " + subject + "\r\n" +
        "\r\n" +
        body + "\r\n")
    
    return smtp.SendMail(e.smtpHost+":"+e.smtpPort, auth, e.username, []string{to}, msg)
}

func (e *Emailer) SendTemplateEmail(to, template string, data map[string]string) error {
    log.Printf("Sending template email (%s) to: %s", template, to)
    
    // TODO: Implement template rendering
    return e.SendEmail(to, "Template Email", "Template content")
}
