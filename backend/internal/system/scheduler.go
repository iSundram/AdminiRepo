
package system

import (
    "log"
    "time"
)

type ScheduledTask struct {
    ID       string        `json:"id"`
    Name     string        `json:"name"`
    Schedule string        `json:"schedule"`
    Command  string        `json:"command"`
    NextRun  time.Time     `json:"next_run"`
    Enabled  bool          `json:"enabled"`
}

type Scheduler struct {
    tasks []ScheduledTask
}

func NewScheduler() *Scheduler {
    return &Scheduler{
        tasks: make([]ScheduledTask, 0),
    }
}

func (s *Scheduler) AddTask(task ScheduledTask) {
    s.tasks = append(s.tasks, task)
    log.Printf("Added scheduled task: %s", task.Name)
}

func (s *Scheduler) RunTasks() {
    for _, task := range s.tasks {
        if task.Enabled && time.Now().After(task.NextRun) {
            log.Printf("Running scheduled task: %s", task.Name)
            // TODO: Execute task command
        }
    }
}
