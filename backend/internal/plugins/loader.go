
package plugins

import (
    "log"
    "plugin"
)

type PluginLoader struct {
    loadedPlugins map[string]*plugin.Plugin
}

func NewPluginLoader() *PluginLoader {
    return &PluginLoader{
        loadedPlugins: make(map[string]*plugin.Plugin),
    }
}

func (pl *PluginLoader) LoadPlugin(path string) error {
    log.Printf("Loading plugin from: %s", path)
    
    p, err := plugin.Open(path)
    if err != nil {
        return err
    }
    
    pl.loadedPlugins[path] = p
    return nil
}

func (pl *PluginLoader) UnloadPlugin(path string) {
    delete(pl.loadedPlugins, path)
}
