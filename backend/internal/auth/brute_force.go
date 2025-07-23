
package auth

import "time"

// Brute force protection
type BruteForceProtection struct {
    attempts map[string][]time.Time
    maxAttempts int
    timeWindow time.Duration
}

func NewBruteForceProtection() *BruteForceProtection {
    return &BruteForceProtection{
        attempts: make(map[string][]time.Time),
        maxAttempts: 5,
        timeWindow: 15 * time.Minute,
    }
}
