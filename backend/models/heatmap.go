
package models

import (
    "time"
)

// Heatmap represents a heatmap data model
type Heatmap struct {
    ID          uint              `gorm:"primaryKey" json:"id"`
    Name        string            `gorm:"not null" json:"name"`
    Type        string            `gorm:"not null" json:"type"`
    Description string            `json:"description"`
    Data        string            `gorm:"type:text" json:"data"` // JSON data
    Config      string            `gorm:"type:text" json:"config"` // JSON config
    UserID      uint              `json:"user_id"`
    IsPublic    bool              `gorm:"default:false" json:"is_public"`
    ViewCount   int               `gorm:"default:0" json:"view_count"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
    
    // Relationships
    User        User              `gorm:"foreignKey:UserID" json:"user,omitempty"`
    Tags        []HeatmapTag      `gorm:"many2many:heatmap_tags" json:"tags,omitempty"`
}

// HeatmapTag represents tags for heatmaps
type HeatmapTag struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Name      string    `gorm:"unique;not null" json:"name"`
    Color     string    `json:"color"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// HeatmapView represents heatmap view tracking
type HeatmapView struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    HeatmapID uint      `json:"heatmap_id"`
    UserID    uint      `json:"user_id"`
    IPAddress string    `json:"ip_address"`
    UserAgent string    `json:"user_agent"`
    ViewedAt  time.Time `json:"viewed_at"`
    
    // Relationships
    Heatmap   Heatmap   `gorm:"foreignKey:HeatmapID" json:"heatmap,omitempty"`
    User      User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// TableName sets the table name for Heatmap
func (Heatmap) TableName() string {
    return "heatmaps"
}

// TableName sets the table name for HeatmapTag
func (HeatmapTag) TableName() string {
    return "heatmap_tags"
}

// TableName sets the table name for HeatmapView
func (HeatmapView) TableName() string {
    return "heatmap_views"
}
