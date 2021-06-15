package cmf

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	color "github.com/logrusorgru/aurora/v3"
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

func askForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

func NewCMF(repository Repository, templateManager TemplateManager, fsManager FS) CMF {
	return &cmf{
		repository:      repository,
		templateManager: templateManager,
		fs:              fsManager,
	}
}

func (cmfInstance *cmf) getInnerVariables() map[string]string {
	extra := map[string]string{
		"BRANCH_NAME": cmfInstance.repository.BranchName(),
	}

	return extra
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

	message, _ := cmfInstance.templateManager.Run(cmfFile, cmfInstance.getInnerVariables())
	cmfInstance.repository.Commit(message)
}

// CommitAmend perform a commit amend over current repository
func (cmfInstance *cmf) CommitAmend() {
	cmfInstance.repository.CheckWorkspaceChanges()
	currentDirectory, _ := cmfInstance.fs.GetCurrentDirectory()
	cmfFile, err := cmfInstance.fs.GetFileFromFS(currentDirectory + "/" + defaultCMFFile)
	if err != nil {
		cmfFile, _ = cmfInstance.fs.GetFileFromVirtualFS(defaultYamlFile)
	}

	message, _ := cmfInstance.templateManager.Run(cmfFile, cmfInstance.getInnerVariables())
	cmfInstance.repository.Amend(message)
}

// InitializeProject initialize current directory with a inner cmf template
func (cmfInstance *cmf) InitializeProject() {
	if askForConfirmation("This action will create a new .cmf.yaml file on your working directory. Do you want to continue?") {
		currentDirectory, _ := cmfInstance.fs.GetCurrentDirectory()
		cmfFilePath := currentDirectory + "/" + defaultCMFFile
		cmfFile, _ := cmfInstance.fs.GetFileFromVirtualFS(defaultYamlFile)
		err := ioutil.WriteFile(cmfFilePath, []byte(cmfFile), 0644)
		if err != nil {
			fmt.Println(color.Red("Cannot create .cmf.yaml file"))
			os.Exit(2)
		}

		fmt.Println(color.Green("You can customize your flow, just visit: https://github.com/walmartdigital/commit-message-formatter. Enjoy!"))
	}
}
