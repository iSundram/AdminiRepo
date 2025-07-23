
package plugins

import (
    "encoding/json"
    "io/ioutil"
)

type PluginRegistry struct {
    plugins map[string]PluginInfo
}

type PluginInfo struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Version     string `json:"version"`
    Description string `json:"description"`
    Author      string `json:"author"`
    Path        string `json:"path"`
    Enabled     bool   `json:"enabled"`
}

func NewPluginRegistry() *PluginRegistry {
    return &PluginRegistry{
        plugins: make(map[string]PluginInfo),
    }
}

func (pr *PluginRegistry) LoadFromFile(filename string) error {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return err
    }
    
    return json.Unmarshal(data, &pr.plugins)
}

func (pr *PluginRegistry) SaveToFile(filename string) error {
    data, err := json.Marshal(pr.plugins)
    if err != nil {
        return err
    }
    
    return ioutil.WriteFile(filename, data, 0644)
}
