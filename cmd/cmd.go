package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/walmartdigital/commit-message-formatter/cmf"
	git "github.com/walmartdigital/commit-message-formatter/git"
)

var cmfInstance cmf.CMF

// Root Root cli command
var root = &cobra.Command{
	Use:   "cmf",
	Short: "Commit Message Formatter",
	Long:  "Generate custom commit message for your repo and standarize your commits log",
	Run: func(cmd *cobra.Command, args []string) {
		cmfInstance.CommitChanges()
		// go git.CheckTree()
		// message := template.Run()
		// git.Commit(message)
	},
}

var version = &cobra.Command{
	Use:   "version",
	Short: "Version cmf",
	Long:  "Display version of commit message formatter",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("CMF - Commit Message Formatter v2.0")
		cmfInstance.GetVersion()
	},
}

var amend = &cobra.Command{
	Use:   "amend",
	Short: "Amend commit message",
	Long:  "Amend last commit message",
	Run: func(cmd *cobra.Command, args []string) {
		cmfInstance.CommitAmend()
		// message := template.Run()
		// git.Amend(message)
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

func init() {
	cmfInstance = cmf.NewCMF(git.NewGitWrapper())
	root.AddCommand(version)
	root.AddCommand(boilerplateCMD)
	root.AddCommand(amend)
}
