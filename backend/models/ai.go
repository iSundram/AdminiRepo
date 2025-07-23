
package models

import (
    "time"
    "gorm.io/gorm"
)

// AIModel represents an AI model configuration
type AIModel struct {
    ID          uint           `json:"id" gorm:"primarykey"`
    Name        string         `json:"name" gorm:"not null"`
    Type        string         `json:"type"` // prediction, analysis, copilot
    Version     string         `json:"version"`
    Status      string         `json:"status"` // active, training, inactive
    Accuracy    float64        `json:"accuracy"`
    Config      string         `json:"config" gorm:"type:json"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// AIPrediction represents an AI prediction result
type AIPrediction struct {
    ID          uint           `json:"id" gorm:"primarykey"`
    ModelID     uint           `json:"model_id"`
    Model       AIModel        `json:"model"`
    Type        string         `json:"type"` // resource, failure, security
    Input       string         `json:"input" gorm:"type:json"`
    Output      string         `json:"output" gorm:"type:json"`
    Confidence  float64        `json:"confidence"`
    ActualValue string         `json:"actual_value" gorm:"type:json"`
    Accuracy    float64        `json:"accuracy"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
