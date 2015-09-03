package git

import (
	"fmt"

	"github.com/libgit2/git2go"
	"github.com/mason-ci/mason/tool"
)

type GitTool struct {
}

func (t *GitTool) Identifier() string {
	return "git"
}

func (t *GitTool) Use(job tool.Job) (bool, error) {
	fmt.Println("GitTool getting latest")
	return true, nil
}

func (t *GitTool) Checkout() {
	git.OpenRepository("lulz")
}

func (t *GitTool) CheckChanges() {
	git.OpenRepository("lulz")
}

func init() {
	git := &GitTool{}
	tool.RegisterVersionControlTool(git)
}
