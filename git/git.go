package git

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	color "github.com/logrusorgru/aurora/v3"
)

type git struct{}

type Git interface {
	CheckWorkspaceChanges()
	Commit(message string)
	Amend(message string)
	BranchName() string
}

func commit(cmdGit *exec.Cmd, message ...interface{}) {
	fmt.Println("")
	fmt.Println(message...)
	fmt.Println(color.Gray(16-1, "-------------------------------"))
	fmt.Println("")
	cmdGit.Stdout = os.Stdout
	cmdGit.Stderr = os.Stderr
	err := cmdGit.Run()
	fmt.Println("")
	fmt.Println(color.Gray(16-1, "-------------------------------"))
	if err == nil {
		fmt.Println(color.Green("Done \U0001F604"))
	} else {
		fmt.Println(color.Red("Something went grong \U0001F92F"))
		os.Exit(2)
	}
}

func NewGitWrapper() Git {
	return &git{}
}

func (gitInstance *git) CheckWorkspaceChanges() {
	cmdGit := exec.Command("git", "diff", "--cached", "--exit-code")
	_, err := cmdGit.Output()
	if err == nil {
		fmt.Println(color.Red(errors.New("no tracked changes")))
		fmt.Println(color.Blue("run git add <file> or . 'to track changes'"))
		os.Exit(1)
	}
}

func (gitInstance *git) Commit(message string) {
	ctx := context.Background()
	cmdGit := exec.CommandContext(ctx, "git", "commit", "-m", message)
	commit(cmdGit, "Committing: ", color.Blue(message))
}

func (gitInstance *git) Amend(message string) {
	ctx := context.Background()
	cmdGit := exec.CommandContext(ctx, "git", "commit", "--amend", "-m", message)
	commit(cmdGit, "Amending: ", color.Blue(message))
}

func (gitInstance *git) BranchName() string {
	cmdGit := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	branchName, err := cmdGit.Output()
	if err != nil {
		return ""
	}

	return strings.TrimSuffix(string(branchName), "\n")
}
