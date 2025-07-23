
package automation

import "time"

// Task Scheduler
type Scheduler struct {
    tasks map[string]*ScheduledTask
}

type ScheduledTask struct {
    ID       string
    Name     string
    Schedule string
    NextRun  time.Time
    Action   func() error
}

func NewScheduler() *Scheduler {
    return &Scheduler{
        tasks: make(map[string]*ScheduledTask),
    }
}
