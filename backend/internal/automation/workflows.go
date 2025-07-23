
package automation

// Workflows automation module
type WorkflowManager struct {
    workflows map[string]*Workflow
}

type Workflow struct {
    ID    string
    Name  string
    Steps []WorkflowStep
}

type WorkflowStep struct {
    Action string
    Config map[string]interface{}
}

func NewWorkflowManager() *WorkflowManager {
    return &WorkflowManager{
        workflows: make(map[string]*Workflow),
    }
}
