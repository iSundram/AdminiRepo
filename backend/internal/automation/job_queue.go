
package automation

// Job Queue management
type JobQueue struct {
    jobs []Job
}

type Job struct {
    ID       string
    Type     string
    Payload  map[string]interface{}
    Status   string
    Priority int
}

func NewJobQueue() *JobQueue {
    return &JobQueue{
        jobs: make([]Job, 0),
    }
}
