
package models

import (
    "time"
    "gorm.io/gorm"
)

// Workflow represents an automation workflow
type Workflow struct {
    ID          uint           `json:"id" gorm:"primarykey"`
    Name        string         `json:"name" gorm:"not null"`
    Description string         `json:"description"`
    UserID      uint           `json:"user_id"`
    User        User           `json:"user"`
    Type        string         `json:"type"` // backup, deployment, monitoring
    Trigger     string         `json:"trigger"` // schedule, event, manual
    Schedule    string         `json:"schedule"` // cron expression
    Status      string         `json:"status"` // active, paused, failed
    Steps       string         `json:"steps" gorm:"type:json"`
    Config      string         `json:"config" gorm:"type:json"`
    LastRun     *time.Time     `json:"last_run"`
    NextRun     *time.Time     `json:"next_run"`
    RunCount    int            `json:"run_count"`
    SuccessRate float64        `json:"success_rate"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// WorkflowExecution represents a workflow execution
type WorkflowExecution struct {
    ID         uint           `json:"id" gorm:"primarykey"`
    WorkflowID uint           `json:"workflow_id"`
    Workflow   Workflow       `json:"workflow"`
    Status     string         `json:"status"` // running, completed, failed
    StartedAt  time.Time      `json:"started_at"`
    CompletedAt *time.Time    `json:"completed_at"`
    Duration   int            `json:"duration"` // seconds
    Steps      string         `json:"steps" gorm:"type:json"`
    Logs       string         `json:"logs" gorm:"type:text"`
    Error      string         `json:"error"`
    CreatedAt  time.Time      `json:"created_at"`
    UpdatedAt  time.Time      `json:"updated_at"`
    DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
