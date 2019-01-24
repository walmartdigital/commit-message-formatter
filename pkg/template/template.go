package template

import (
	"fmt"
	"os"
	"prompt"
	"strings"

	color "github.com/logrusorgru/aurora"
	"github.com/spf13/viper"
)

type option struct {
	Value string
	Desc  string
}

type keyValue struct {
	Key   string
	Value string
}

func checkErr(err error) {
	if err != nil {
		fmt.Print(color.Red("An error occurred, maybe 'cmf init' could fix it"))
		os.Exit(0)
	}
}

func parseTemplate(template string, variables []keyValue) string {
	for _, v := range variables {
		template = strings.Replace(template, "{{"+v.Key+"}}", v.Value, -1)
	}

	return template
}

func executeFlow(settings map[string]interface{}) (variables []keyValue) {
	prompts := settings["prompt"].([]interface{})
	for _, v := range prompts {
		result := ""
		var label = ""
		var errorMessage = "Field empty"
		var defaultValue = ""

		pr := v.(map[interface{}]interface{})

		if value, ok := pr["LABEL"]; ok {
			label = value.(string)
		}

		if value, ok := pr["ERROR_LABEL"]; ok {
			errorMessage = value.(string)
		}

		if value, ok := pr["DEFAULT_VALUE"]; ok {
			defaultValue = value.(string)
		}

		if options := pr["OPTIONS"]; options == nil {
			result = prompt.Input(label, errorMessage, defaultValue)
		} else {
			configurationSelectOption := &prompt.SelectConfiguration{
				ActiveTpl:    "\U0001F449  {{ .Title | cyan | bold }}",
				InactiveTpl:  "    {{ .Title | white }}",
				SelectPrompt: label,
				SelectedTpl:  "\U0001F44D {{ \"" + label + "\" | cyan }} {{ .Title | cyan | bold }}",
				DetailsTpl: `-------------------------------
{{ .Title | white | bold }} {{ .Description | white | bold }}`,
			}

			optList := pr["OPTIONS"].([]interface{})
			var opts []prompt.SelectItem
			for _, o := range optList {
				op := o.(map[interface{}]interface{})

				opts = append(opts, prompt.SelectItem{Title: op["VALUE"].(string), Value: op["VALUE"].(string), Description: op["DESC"].(string)})
			}

			selector := prompt.Select(opts, *configurationSelectOption)
			result = selector.Value
		}

		variables = append(variables, keyValue{
			Key:   pr["KEY"].(string),
			Value: result,
		})
	}

	return
}

func setFlowFileAsVariables() map[string]interface{} {
	projectDir, err := os.Getwd()
	checkErr(err)
	projectConfigFile := projectDir + "/.cmf.yaml"

	_, err = os.Stat(projectConfigFile)
	checkErr(err)

	viper.SetConfigFile(projectConfigFile)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	checkErr(err)

	return viper.AllSettings()
}

// Run initialize prompt flow
func Run() string {
	settings := setFlowFileAsVariables()
	variables := executeFlow(settings)
	template := settings["template"].(string)
	return parseTemplate(template, variables)
}
