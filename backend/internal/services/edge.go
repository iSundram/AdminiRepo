
package services

import (
    "fmt"
    "time"
)

// EdgeManager handles edge computing and CDN
type EdgeManager struct {
    nodes     map[string]*EdgeNode
    zones     map[string]*EdgeZone
    configs   map[string]*EdgeConfig
}

// EdgeNode represents an edge computing node
type EdgeNode struct {
    ID         string             `json:"id"`
    Name       string             `json:"name"`
    Location   EdgeLocation       `json:"location"`
    Status     string             `json:"status"`
    Capacity   EdgeCapacity       `json:"capacity"`
    Usage      EdgeUsage          `json:"usage"`
    Services   []string           `json:"services"`
    CreatedAt  time.Time          `json:"created_at"`
    LastSeen   time.Time          `json:"last_seen"`
    Health     EdgeHealth         `json:"health"`
}

// EdgeZone represents an edge zone
type EdgeZone struct {
    ID          string      `json:"id"`
    Name        string      `json:"name"`
    Region      string      `json:"region"`
    Nodes       []string    `json:"nodes"`
    Status      string      `json:"status"`
    Config      EdgeConfig  `json:"config"`
    CreatedAt   time.Time   `json:"created_at"`
}

// EdgeLocation represents geographical location
type EdgeLocation struct {
    Country   string  `json:"country"`
    City      string  `json:"city"`
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    Region    string  `json:"region"`
}

// EdgeCapacity represents node capacity
type EdgeCapacity struct {
    CPU    int `json:"cpu"`    // cores
    Memory int `json:"memory"` // MB
    Storage int `json:"storage"` // GB
    Bandwidth int `json:"bandwidth"` // Mbps
}

// EdgeUsage represents current usage
type EdgeUsage struct {
    CPU       float64 `json:"cpu"`       // percentage
    Memory    float64 `json:"memory"`    // percentage
    Storage   float64 `json:"storage"`   // percentage
    Bandwidth float64 `json:"bandwidth"` // percentage
}

// EdgeHealth represents node health
type EdgeHealth struct {
    Score    float64           `json:"score"`    // 0-100
    Issues   []string          `json:"issues"`
    Metrics  map[string]float64 `json:"metrics"`
    Status   string            `json:"status"`
}

// EdgeConfig represents edge configuration
type EdgeConfig struct {
    Caching    CacheConfig    `json:"caching"`
    Security   SecurityConfig `json:"security"`
    Routing    RoutingConfig  `json:"routing"`
    Analytics  bool           `json:"analytics"`
}

// CacheConfig represents caching configuration
type CacheConfig struct {
    Enabled    bool              `json:"enabled"`
    TTL        int               `json:"ttl"`
    Rules      []CacheRule       `json:"rules"`
    Purge      PurgeConfig       `json:"purge"`
}

// CacheRule represents a caching rule
type CacheRule struct {
    Pattern string `json:"pattern"`
    TTL     int    `json:"ttl"`
    Action  string `json:"action"`
}

// PurgeConfig represents cache purging configuration
type PurgeConfig struct {
    Auto     bool     `json:"auto"`
    Patterns []string `json:"patterns"`
}

// SecurityConfig represents security configuration
type SecurityConfig struct {
    WAF        bool     `json:"waf"`
    DDoS       bool     `json:"ddos"`
    RateLimit  bool     `json:"rate_limit"`
    GeoBlock   []string `json:"geo_block"`
}

// RoutingConfig represents routing configuration
type RoutingConfig struct {
    LoadBalancing string            `json:"load_balancing"`
    HealthCheck   bool              `json:"health_check"`
    Failover      bool              `json:"failover"`
    Headers       map[string]string `json:"headers"`
}

// NewEdgeManager creates a new edge manager
func NewEdgeManager() *EdgeManager {
    return &EdgeManager{
        nodes:   make(map[string]*EdgeNode),
        zones:   make(map[string]*EdgeZone),
        configs: make(map[string]*EdgeConfig),
    }
}

// RegisterNode registers a new edge node
func (em *EdgeManager) RegisterNode(name string, location EdgeLocation, capacity EdgeCapacity) (*EdgeNode, error) {
    nodeID := fmt.Sprintf("edge_%s_%d", name, time.Now().Unix())
    
    node := &EdgeNode{
        ID:       nodeID,
        Name:     name,
        Location: location,
        Status:   "online",
        Capacity: capacity,
        Usage: EdgeUsage{
            CPU:       0,
            Memory:    0,
            Storage:   0,
            Bandwidth: 0,
        },
        Services:  make([]string, 0),
        CreatedAt: time.Now(),
        LastSeen:  time.Now(),
        Health: EdgeHealth{
            Score:   100,
            Issues:  make([]string, 0),
            Metrics: make(map[string]float64),
            Status:  "healthy",
        },
    }
    
    em.nodes[nodeID] = node
    return node, nil
}

// GetNode retrieves an edge node
func (em *EdgeManager) GetNode(nodeID string) (*EdgeNode, bool) {
    node, exists := em.nodes[nodeID]
    return node, exists
}

// UpdateNodeHealth updates node health status
func (em *EdgeManager) UpdateNodeHealth(nodeID string, health EdgeHealth) error {
    node, exists := em.nodes[nodeID]
    if !exists {
        return fmt.Errorf("node not found: %s", nodeID)
    }
    
    node.Health = health
    node.LastSeen = time.Now()
    
    if health.Score < 50 {
        node.Status = "degraded"
    } else if health.Score < 20 {
        node.Status = "offline"
    } else {
        node.Status = "online"
    }
    
    return nil
}

// CreateZone creates a new edge zone
func (em *EdgeManager) CreateZone(name, region string) (*EdgeZone, error) {
    zoneID := fmt.Sprintf("zone_%s_%s", region, name)
    
    zone := &EdgeZone{
        ID:        zoneID,
        Name:      name,
        Region:    region,
        Nodes:     make([]string, 0),
        Status:    "active",
        Config:    EdgeConfig{},
        CreatedAt: time.Now(),
    }
    
    em.zones[zoneID] = zone
    return zone, nil
}

// AddNodeToZone adds a node to an edge zone
func (em *EdgeManager) AddNodeToZone(zoneID, nodeID string) error {
    zone, exists := em.zones[zoneID]
    if !exists {
        return fmt.Errorf("zone not found: %s", zoneID)
    }
    
    if _, exists := em.nodes[nodeID]; !exists {
        return fmt.Errorf("node not found: %s", nodeID)
    }
    
    // Check if node is already in zone
    for _, existingNode := range zone.Nodes {
        if existingNode == nodeID {
            return nil // Already in zone
        }
    }
    
    zone.Nodes = append(zone.Nodes, nodeID)
    return nil
}

// ListNodes returns all edge nodes
func (em *EdgeManager) ListNodes() map[string]*EdgeNode {
    return em.nodes
}

// ListZones returns all edge zones
func (em *EdgeManager) ListZones() map[string]*EdgeZone {
    return em.zones
}

// GetClosestNode finds the closest node to a location
func (em *EdgeManager) GetClosestNode(lat, lon float64) (*EdgeNode, error) {
    var closestNode *EdgeNode
    var minDistance float64 = -1
    
    for _, node := range em.nodes {
        if node.Status != "online" {
            continue
        }
        
        distance := calculateDistance(lat, lon, node.Location.Latitude, node.Location.Longitude)
        if minDistance == -1 || distance < minDistance {
            minDistance = distance
            closestNode = node
        }
    }
    
    if closestNode == nil {
        return nil, fmt.Errorf("no online nodes available")
    }
    
    return closestNode, nil
}

// Simple distance calculation (Haversine formula would be more accurate)
func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
    return ((lat1-lat2)*(lat1-lat2) + (lon1-lon2)*(lon1-lon2))
}
