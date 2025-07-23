
package core

// Feature set management
type FeatureSetManager struct {
    featureSets map[string]*FeatureSet
}

type FeatureSet struct {
    ID          string   `json:"id"`
    Name        string   `json:"name"`
    Features    []string `json:"features"`
    Enabled     bool     `json:"enabled"`
}

func NewFeatureSetManager() *FeatureSetManager {
    return &FeatureSetManager{
        featureSets: make(map[string]*FeatureSet),
    }
}
