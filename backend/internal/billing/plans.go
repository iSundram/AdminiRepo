
package billing

// Subscription plans management
type PlanManager struct {
    plans map[string]*Plan
}

type Plan struct {
    ID          string
    Name        string
    Price       float64
    Interval    string
    Features    []string
    Limits      map[string]int
}

func NewPlanManager() *PlanManager {
    return &PlanManager{
        plans: make(map[string]*Plan),
    }
}
