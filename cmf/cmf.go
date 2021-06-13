package cmf

import (
	"fmt"
)

var version = "3.0"

type Repository interface {
	CheckWorkspaceChanges()
	Commit(message string)
	Amend(message string)
}
type cmf struct {
	repository Repository
}

type CMF interface {
	GetVersion()
	CommitChanges()
	CommitAmend()
	InitializeProject()
}

func NewCMF(repository Repository) CMF {
	return &cmf{
		repository: repository,
	}
}

// GetVersion return current cmf version
func (cmfInstance *cmf) GetVersion() {
	fmt.Println("Git - Commit Message Formatter v", version)
}

// CommitChanges perform a commit changes over current repository
func (cmfInstance *cmf) CommitChanges() {
	cmfInstance.repository.CheckWorkspaceChanges()
	// message := template.Run()
	cmfInstance.repository.Commit("message")
}

// CommitAmend perform a commit amend over current repository
func (cmfInstance *cmf) CommitAmend() {
	fmt.Println("amend!!")
}

// InitializeProject initialize current directory with a inner cmf template
func (cmfInstance *cmf) InitializeProject() {
	fmt.Println("initialize!!")
}
