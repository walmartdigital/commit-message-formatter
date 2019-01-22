package cmd

import (
	"boilerplate"
	"template"

	"github.com/commit-message-formatter/pkg/git"
	"github.com/spf13/cobra"
)

// Root Root cli command
var Root = &cobra.Command{
	Use:    "cmf",
	Short:  "Commit Message Formatter",
	Long:   "Generate custom commit message for your repo and standarize your commits log",
	PreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		go git.CheckTree()
		message := template.Run()
		git.Commit(message)
	},
}

var amend = &cobra.Command{
	Use:    "amend",
	Short:  "Amend commit message",
	Long:   "Amend last commit message",
	PreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		message := template.Run()
		git.Amend(message)
	},
}

var boilerplateCMD = &cobra.Command{
	Use:    "init",
	Short:  "Create configuration file",
	Long:   "Create .cmf.yaml configuration file",
	PreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		boilerplate.Create()
	},
}

func init() {
	Root.AddCommand(boilerplateCMD)
	Root.AddCommand(amend)
}
