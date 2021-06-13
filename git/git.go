package git

import "fmt"

type git struct{}

type Git interface {
	CheckWorkspaceChanges()
	Commit(message string)
	Amend(message string)
}

func NewGitWrapper() Git {
	return &git{}
}

func (gitInstance *git) CheckWorkspaceChanges() {
	fmt.Println("check wrokspace")
}

func (gitInstance *git) Commit(message string) {
	fmt.Println(message)
}

func (gitInstance *git) Amend(message string) {
	fmt.Println("amend")
}

// func checkErr(err error) {
// 	if err != nil {
// 		fmt.Println(color.Red(err))
// 		os.Exit(1)
// 	}
// }

// // CheckTree perform a git tree check
// func CheckTree() {
// 	cmdGit := exec.Command("git", "diff", "--cached", "--exit-code")
// 	_, err := cmdGit.Output()
// 	if err == nil {
// 		checkErr(errors.New("No changes added to commit"))
// 	}
// }

// func commit(cmdGit *exec.Cmd, message ...interface{}) {
// 	fmt.Println("")
// 	fmt.Println(message...)
// 	// fmt.Println(color.Gray("-------------------------------"))
// 	fmt.Println("")
// 	cmdGit.Stdout = os.Stdout
// 	cmdGit.Stderr = os.Stderr
// 	err := cmdGit.Run()
// 	fmt.Println("")
// 	// fmt.Println(color.Gray("-------------------------------"))
// 	if err == nil {
// 		fmt.Println(color.Green("Done \U0001F604"))
// 	} else {
// 		fmt.Println(color.Red("Something went grong \U0001F92F"))
// 	}

// 	return
// }

// // Commit execute commit
// func Commit(message string) (err error) {
// 	ctx := context.Background()
// 	cmdGit := exec.CommandContext(ctx, "git", "commit", "-m", message)
// 	commit(cmdGit, "Committing: ", color.Blue(message))

// 	return
// }

// // Amend execute commit amend
// func Amend(message string) (err error) {
// 	ctx := context.Background()
// 	cmdGit := exec.CommandContext(ctx, "git", "commit", "--amend", "-m", message)
// 	commit(cmdGit, "Amending: ", color.Blue(message))

// 	return
// }

// // BranchName return current branch name
// func BranchName() string {
// 	cmdGit := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
// 	branchName, err := cmdGit.Output()
// 	if err != nil {
// 		return ""
// 	}

// 	return strings.TrimSuffix(string(branchName), "\n")
// }
