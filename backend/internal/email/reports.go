
package email

import "time"

// Email reporting
type ReportManager struct {
    reports map[string]*EmailReport
}

type EmailReport struct {
    ID        string    `json:"id"`
    Type      string    `json:"type"`
    Period    string    `json:"period"`
    Data      map[string]interface{} `json:"data"`
    CreatedAt time.Time `json:"created_at"`
}

func NewReportManager() *ReportManager {
    return &ReportManager{
        reports: make(map[string]*EmailReport),
    }
}
