package prompt

import (
	"errors"
	"strings"

	"github.com/mritd/promptx"
)

// SelectItem a default item struct
type SelectItem struct {
	Title       string
	Description string
	Value       string
}

// SelectConfiguration a complete configuration select options
type SelectConfiguration struct {
	ActiveTpl    string
	InactiveTpl  string
	SelectPrompt string
	SelectedTpl  string
	DetailsTpl   string
}

// Select Returns a selected item
func Select(items []SelectItem, config SelectConfiguration) SelectItem {
	configuration := &promptx.SelectConfig{
		ActiveTpl:    config.ActiveTpl,
		InactiveTpl:  config.InactiveTpl,
		SelectPrompt: config.SelectPrompt,
		SelectedTpl:  config.SelectedTpl,
		DisPlaySize:  9,
		DetailsTpl:   config.DetailsTpl,
	}

	selector := &promptx.Select{
		Items:  items,
		Config: configuration,
	}

	return items[selector.Run()]
}

// Input return a simple user input
func Input(title string, errorMessage string, defaultValue string) string {
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
