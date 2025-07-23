
package monitoring

// Heatmap data generation
type HeatmapManager struct {
    heatmaps map[string]*Heatmap
}

type Heatmap struct {
    ID       string      `json:"id"`
    Type     string      `json:"type"`
    Data     [][]float64 `json:"data"`
    Labels   []string    `json:"labels"`
    MinValue float64     `json:"min_value"`
    MaxValue float64     `json:"max_value"`
}

func NewHeatmapManager() *HeatmapManager {
    return &HeatmapManager{
        heatmaps: make(map[string]*Heatmap),
    }
}
