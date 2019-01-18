package cmd

import (
	"boilerplate"
	"fmt"

	"github.com/spf13/cobra"
)

// Root Root cli command
var Root = &cobra.Command{
	Use:    "cmf",
	Short:  "Commit Message Formatter",
	Long:   "Generate custom commit message for your repo and standarize your commits log",
	PreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("RUN...")
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
}
