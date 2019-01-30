package git

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"

	color "github.com/logrusorgru/aurora"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(color.Red(err))
		os.Exit(1)
	}
}

// CheckTree perform a git tree check
func CheckTree() {
	cmdGit := exec.Command("git", "diff", "--cached", "--exit-code")
	_, err := cmdGit.Output()
	if err == nil {
		checkErr(errors.New("No changes added to commit"))
	}
}

func commit(cmdGit *exec.Cmd, message ...interface{}) {
	fmt.Println("")
	fmt.Println(message...)
	fmt.Println(color.Gray("-------------------------------"))
	fmt.Println("")
	cmdGit.Stdout = os.Stdout
	cmdGit.Run()
	fmt.Println("")
	fmt.Println(color.Gray("-------------------------------"))
	fmt.Println(color.Green("Done \U0001F604"))

	return
}

// Commit execute commit
func Commit(message string) (err error) {
	ctx := context.Background()
	cmdGit := exec.CommandContext(ctx, "git", "commit", "-m", message)
	commit(cmdGit, "Committing: ", color.Blue(message))

	return
}

// Amend execute commit amend
func Amend(message string) (err error) {
	ctx := context.Background()
	cmdGit := exec.CommandContext(ctx, "git", "commit", "--amend", "-m", message)
	commit(cmdGit, "Amending: ", color.Blue(message))

	return
}
