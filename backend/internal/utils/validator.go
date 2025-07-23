
package utils

import (
    "regexp"
    "strings"
)

type Validator struct{}

func NewValidator() *Validator {
    return &Validator{}
}

func (v *Validator) ValidateEmail(email string) bool {
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return emailRegex.MatchString(email)
}

func (v *Validator) ValidateDomain(domain string) bool {
    domainRegex := regexp.MustCompile(`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$`)
    return domainRegex.MatchString(domain)
}

func (v *Validator) ValidatePassword(password string) bool {
    return len(password) >= 8 && 
           strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") &&
           strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") &&
           strings.ContainsAny(password, "0123456789") &&
           strings.ContainsAny(password, "!@#$%^&*()_+-=[]{}|;:,.<>?")
}

func (v *Validator) ValidateIP(ip string) bool {
    ipRegex := regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
    return ipRegex.MatchString(ip)
}
