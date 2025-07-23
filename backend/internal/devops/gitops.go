
package devops

import (
    "log"
    "os/exec"
)

type GitOps struct {
    repoURL    string
    branch     string
    deployPath string
}

func NewGitOps(repoURL, branch, deployPath string) *GitOps {
    return &GitOps{
        repoURL:    repoURL,
        branch:     branch,
        deployPath: deployPath,
    }
}

func (g *GitOps) SyncRepository() error {
    log.Printf("Syncing repository: %s (branch: %s)", g.repoURL, g.branch)
    
    cmd := exec.Command("git", "clone", "-b", g.branch, g.repoURL, g.deployPath)
    return cmd.Run()
}

func (g *GitOps) Deploy() error {
    log.Println("Deploying via GitOps...")
    
    // Pull latest changes
    cmd := exec.Command("git", "pull", "origin", g.branch)
    cmd.Dir = g.deployPath
    
    return cmd.Run()
}
