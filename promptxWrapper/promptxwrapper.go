package promptxwrapper

import (
	"errors"
	"strings"

	"github.com/mritd/promptx"
	"github.com/walmartdigital/commit-message-formatter/prompt"
	"github.com/walmartdigital/commit-message-formatter/templaterunner"
)

// PromptxWrapper promptx wrapper object
type PromptxWrapper struct{}

// NewPromptxWrapper return a new instance of promptxWrapper
func NewPromptxWrapper() *PromptxWrapper {
	return &PromptxWrapper{}
}

// ReadValue return a value from user single input
func (pw *PromptxWrapper) ReadValue(title string, errorMessage string, defaultValue string) string {
	input := promptx.NewDefaultPrompt(func(line []rune) error {
		if strings.TrimSpace(string(line)) == "" && defaultValue == "" {
			return errors.New(errorMessage)
		}

		return nil
	}, title)

	value := input.Run()

	if value == "" && defaultValue != "" {
		return defaultValue
	}

	return value
}

// ReadValueFromList return a value from user multi select input
func (pw *PromptxWrapper) ReadValueFromList(title string, options []templaterunner.Options) string {
	configuration := &promptx.SelectConfig{
		ActiveTpl:    "\U0001F449  {{ .Title | cyan | bold }}",
		InactiveTpl:  "    {{ .Title | white }}",
		SelectPrompt: title,
		SelectedTpl:  "\U0001F44D {{ \"" + title + "\" | cyan }} {{ .Title | cyan | bold }}",
		DetailsTpl: `-------------------------------
		{{ .Title | white | bold }} {{ .Description | white | bold }}`,
	}

	var items []prompt.SelectItem
	for _, option := range options {
		items = append(items, prompt.SelectItem{Title: option.Value, Value: option.Value, Description: option.Description})
	}

	selector := &promptx.Select{
		Items:  items,
		Config: configuration,
	}

	return items[selector.Run()].Value
}
