package boilerplate

import (
	"fmt"
	"os"
	"prompt"

	packr "github.com/gobuffalo/packr/v2"
	color "github.com/logrusorgru/aurora"
)

func checkErr(err error) {
	if err != nil {
		fmt.Print(color.Red(err))
		os.Exit(0)
	}
}

var flows = map[string][]byte{
	"default": []byte(""),
	"jira":    []byte(""),
	"custom":  []byte(""),
}

var flowTypes = []prompt.SelectItem{
	{Title: "Default", Description: "Flow with module affected and message", Value: "default"},
	{Title: "Jira", Description: "Flow based on Jira task", Value: "jira"},
	{Title: "Custom", Description: "Flow with your own rules", Value: "custom"},
}

var configurationSelectFlow = &prompt.SelectConfiguration{
	ActiveTpl:    "\U0001F449  {{ .Title | cyan }}",
	InactiveTpl:  "   {{ .Title | white }}",
	SelectPrompt: "Flow Type",
	SelectedTpl:  "{{ \"Flow:\" | bold }} {{ .Title | green | bold }}",
	DetailsTpl: `-------------------------------
{{ .Title | white | bold }} {{ .Description | white | bold }}`,
}

func createConfigurationFile(flowType string) error {
	projectFolder, err := os.Getwd()
	if err != nil {
		return err
	}

	configFile, err := os.Create(projectFolder + "/.cmf.yaml")
	defer configFile.Close()

	if err != nil {
		return err
	}

	_, err = configFile.Write(flows[flowType])

	return err
}

// Create a configuration file
func Create() {
	selector :=
		prompt.Select(flowTypes, *configurationSelectFlow)

	err := createConfigurationFile(selector.Value)
	checkErr(err)

	fmt.Println(color.Gray("File .cmf.yaml for " + selector.Value + " flow generated "))
	fmt.Println("Happy coding. \U0001F604")
}

func init() {
	box := packr.New("config", "../../configs")
	flows["default"], _ = box.Find("default.yaml")
	flows["jira"], _ = box.Find("jira.yaml")
	flows["custom"], _ = box.Find("custom.yaml")
}
