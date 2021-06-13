package cmf

import (
	"fmt"
)

const version = "3.0"
const defaultYamlFile = "resources/default.yaml"
const defaultCMFFile = ".cmf.yaml"

type Repository interface {
	CheckWorkspaceChanges()
	Commit(message string)
	Amend(message string)
	BranchName() string
}

type TemplateManager interface {
	Run(yamlData string, injectedVariables map[string]string) (string, error)
}

type FS interface {
	GetFileFromVirtualFS(path string) (string, error)
	GetFileFromFS(path string) (string, error)
	GetCurrentDirectory() (string, error)
}

type cmf struct {
	repository      Repository
	templateManager TemplateManager
	fs              FS
}

type CMF interface {
	GetVersion()
	CommitChanges()
	CommitAmend()
	InitializeProject()
}

func NewCMF(repository Repository, templateManager TemplateManager, fsManager FS) CMF {
	return &cmf{
		repository:      repository,
		templateManager: templateManager,
		fs:              fsManager,
	}
}

// GetVersion return current cmf version
func (cmfInstance *cmf) GetVersion() {
	fmt.Println("Git - Commit Message Formatter v", version)
}

// CommitChanges perform a commit changes over current repository
func (cmfInstance *cmf) CommitChanges() {
	cmfInstance.repository.CheckWorkspaceChanges()
	currentDirectory, _ := cmfInstance.fs.GetCurrentDirectory()
	cmfFile, err := cmfInstance.fs.GetFileFromFS(currentDirectory + "/" + defaultCMFFile)
	if err != nil {
		cmfFile, _ = cmfInstance.fs.GetFileFromVirtualFS(defaultYamlFile)
	}

	extra := map[string]string{
		"BRANCH_NAME": cmfInstance.repository.BranchName(),
	}
	message, _ := cmfInstance.templateManager.Run(cmfFile, extra)
	cmfInstance.repository.Commit(message)
}

// CommitAmend perform a commit amend over current repository
func (cmfInstance *cmf) CommitAmend() {
	fmt.Println("amend!!")
}

// InitializeProject initialize current directory with a inner cmf template
func (cmfInstance *cmf) InitializeProject() {
	fmt.Println("initialize!!")
}
