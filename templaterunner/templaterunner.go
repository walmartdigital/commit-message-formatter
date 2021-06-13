package templaterunner

import (
	"errors"
	"strings"

	"gopkg.in/yaml.v2"
)

type templateRunner struct {
	promptManager PromptManager
}

// PromptManager ...
type PromptManager interface {
	ReadValue(title string, errorMessage string, defaultValue string) string
	ReadValueFromList(title string, options []Options) string
}

// TemplateRunner main template interface
type TemplateRunner interface {
	Run(yamlData string, injectedVariables map[string]string) (message string, err error)
}

// Template main template struct
type Template struct {
	Prompt   []PromptItem `yaml:"PROMPT"`
	Template string       `yaml:"TEMPLATE"`
}

//PromptItem ...
type PromptItem struct {
	Key          string    `yaml:"KEY"`
	Label        string    `yaml:"LABEL"`
	ErrorLabel   string    `yaml:"ERROR_LABEL"`
	DefaultValue string    `yaml:"DEFAULT_VALUE"`
	Options      []Options `yaml:"OPTIONS"`
}

// Options multiselect option struct
type Options struct {
	Value       string `yaml:"VALUE"`
	Description string `yaml:"DESC"`
}

type keyValue struct {
	Key   string
	Value string
}

// NewTemplateRunner return a  bluetnew instance of template
func NewTemplateRunner(promptManager PromptManager) TemplateRunner {
	return &templateRunner{
		promptManager: promptManager,
	}
}

func (tr *templateRunner) parseYaml(yamlData string) (Template, error) {
	template := Template{}
	err := yaml.Unmarshal([]byte(yamlData), &template)

	if err != nil {
		return Template{}, errors.New("parsing yaml error")
	}

	return template, nil
}

// Run return the result of run the template
func (tr *templateRunner) Run(yamlData string, injectedVariables map[string]string) (string, error) {
	template, err := tr.parseYaml(yamlData)
	if err != nil {
		return "", err
	}

	variables := tr.prompt(template)
	for k, v := range injectedVariables {
		variables = append(variables, keyValue{Key: k, Value: v})
	}

	message := tr.parseTemplate(template.Template, variables)

	return message, err
}

func (tr *templateRunner) parseTemplate(template string, variables []keyValue) string {
	for _, v := range variables {
		template = strings.Replace(template, "{{"+v.Key+"}}", v.Value, -1)
	}

	return template
}

func (tr *templateRunner) prompt(template Template) []keyValue {
	variables := []keyValue{}
	for _, step := range template.Prompt {
		result := ""
		var errorMessage = "empty value"

		if step.ErrorLabel != "" {
			errorMessage = step.ErrorLabel
		}

		if step.Options == nil {
			var labelMessage = step.Label

			if step.DefaultValue != "" {
				labelMessage += " (" + step.DefaultValue + ")"
			}

			labelMessage += ":"

			result = tr.promptManager.ReadValue(labelMessage, errorMessage, step.DefaultValue)
		} else {
			result = tr.promptManager.ReadValueFromList(step.Label, step.Options)
		}

		variables = append(variables, keyValue{
			Key:   step.Key,
			Value: result,
		})
	}

	return variables
}
