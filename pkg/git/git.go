package git

import (
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
	output, err := cmdGit.Output()
	checkErr(err)

	fmt.Println("")
	fmt.Println("Committing: ", color.Blue(message))
	fmt.Println(color.Gray("-------------------------------"))
	fmt.Println("")
	fmt.Println(color.Gray(string(output)))
	fmt.Println(color.Gray("-------------------------------"))
	fmt.Println(color.Green("Done \U0001F604"))

	return
}

// Commit execute commit
func Commit(message string) (err error) {
	cmdGit := exec.Command("git", "commit", "-m", message)
	commit(cmdGit, "Committing: ", color.Blue(message))

	return
}

// Amend execute commit amend
func Amend(message string) (err error) {
	cmdGit := exec.Command("git", "commit", "--amend", "-m", message)
	commit(cmdGit, "Amending: ", color.Blue(message))

	return
}
