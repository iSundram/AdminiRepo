
package system

import (
    "runtime"
    "time"
)

type SystemInfo struct {
    OS           string    `json:"os"`
    Architecture string    `json:"architecture"`
    CPUs         int       `json:"cpus"`
    GoVersion    string    `json:"go_version"`
    Uptime       time.Time `json:"uptime"`
    Version      string    `json:"version"`
}

func GetSystemInfo() *SystemInfo {
    return &SystemInfo{
        OS:           runtime.GOOS,
        Architecture: runtime.GOARCH,
        CPUs:         runtime.NumCPU(),
        GoVersion:    runtime.Version(),
        Uptime:       time.Now(),
        Version:      "1.0.0",
    }
}
