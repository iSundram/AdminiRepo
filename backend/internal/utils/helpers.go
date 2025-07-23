
package utils

import (
    "crypto/rand"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "time"
)

func GenerateRandomString(length int) (string, error) {
    bytes := make([]byte, length)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes)[:length], nil
}

func FormatBytes(bytes int64) string {
    const unit = 1024
    if bytes < unit {
        return fmt.Sprintf("%d B", bytes)
    }
    div, exp := int64(unit), 0
    for n := bytes / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func PrettyJSON(v interface{}) (string, error) {
    b, err := json.MarshalIndent(v, "", "  ")
    return string(b), err
}

func TimeAgo(t time.Time) string {
    now := time.Now()
    diff := now.Sub(t)
    
    if diff < time.Minute {
        return "just now"
    } else if diff < time.Hour {
        return fmt.Sprintf("%d minutes ago", int(diff.Minutes()))
    } else if diff < 24*time.Hour {
        return fmt.Sprintf("%d hours ago", int(diff.Hours()))
    } else {
        return fmt.Sprintf("%d days ago", int(diff.Hours()/24))
    }
}

func Contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
}
