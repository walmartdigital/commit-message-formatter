package cmd

import (
	"embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/walmartdigital/commit-message-formatter/cmf"
	"github.com/walmartdigital/commit-message-formatter/fs"
	git "github.com/walmartdigital/commit-message-formatter/git"
	promptxwrapper "github.com/walmartdigital/commit-message-formatter/promptxWrapper"
	"github.com/walmartdigital/commit-message-formatter/templaterunner"
)

var cmfInstance cmf.CMF

// Root Root cli command
var root = &cobra.Command{
	Use:   "cmf",
	Short: "Commit Message Formatter",
	Long:  "Generate custom commit message for your repo and standarize your commits log",
	Run: func(cmd *cobra.Command, args []string) {
		cmfInstance.CommitChanges()
	},
}

var version = &cobra.Command{
	Use:   "version",
	Short: "Version cmf",
	Long:  "Display version of commit message formatter",
	Run: func(cmd *cobra.Command, args []string) {
		cmfInstance.GetVersion()
	},
}

var amend = &cobra.Command{
	Use:   "amend",
	Short: "Amend commit message",
	Long:  "Amend last commit message",
	Run: func(cmd *cobra.Command, args []string) {
		cmfInstance.CommitAmend()
	},
}

var boilerplateCMD = &cobra.Command{
	Use:   "init",
	Short: "Create configuration file",
	Long:  "Create .cmf.yaml configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		cmfInstance.InitializeProject()
	},
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Build(vfs embed.FS) {
	fsManager := fs.NewFs(vfs)
	promptManager := promptxwrapper.NewPromptxWrapper()
	templateManager := templaterunner.NewTemplateRunner(promptManager)
	cmfInstance = cmf.NewCMF(git.NewGitWrapper(), templateManager, fsManager)
	root.AddCommand(version)
	root.AddCommand(boilerplateCMD)
	root.AddCommand(amend)
}
