// Copyright © 2018 Rodrigo Navarro <rodrigonavarro23@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	. "github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	gitconfig "walmart.com/cfm/src/github.com/tcnksm/go-gitconfig"
)

func checkErr(err error) {
	if err != nil {
		color.Set(color.FgMagenta)
		defer color.Unset()
		fmt.Print(err)
		os.Exit(1)
	}
}

func parseTemplate(template string) string {
	for _, v := range variables {
		template = strings.Replace(template, "{{"+v.Key+"}}", v.Value, -1)
	}

	return template
}

func promptList() {
	settings := viper.AllSettings()
	prompts := settings["prompt"].([]interface{})
	for _, v := range prompts {
		result := ""
		pr := v.(map[interface{}]interface{})
		if options := pr["OPTIONS"]; options == nil {
			validate := func(input string) error {
				if input == "" {
					return errors.New("Empty value")
				}
				return nil
			}
			p := promptui.Prompt{
				Label: pr["LABEL"],
				// Templates: templates,
				Validate: validate,
			}
			r, err := p.Run()
			checkErr(err)
			result = r
		} else {

			templates := &promptui.SelectTemplates{
				Label:    "{{ . | bold }}",
				Active:   "\U0001F449 {{ .Value | cyan }}  {{ .Desc | faint }}",
				Inactive: "   {{ .Value }}  {{ .Desc | faint }}",
				Selected: "\U0001F44D {{ .Value |  bold }}",
			}
			optList := pr["OPTIONS"].([]interface{})
			var opts []*option
			for _, o := range optList {
				op := o.(map[interface{}]interface{})
				opts = append(opts, &option{Value: op["VALUE"].(string), Desc: op["DESC"].(string)})
			}
			p := promptui.Select{
				Label:     pr["LABEL"],
				Items:     opts,
				Templates: templates,
			}
			i, _, err := p.Run()
			checkErr(err)
			result = opts[i].Value
		}

		variables = append(variables, keyValue{
			Key:   pr["KEY"].(string),
			Value: result,
		})
	}
}

func commit(message string) (err error) {
	directory, err := os.Getwd()
	checkErr(err)

	// Opens an already existent repository.
	r, err := git.PlainOpen(directory)
	checkErr(err)

	w, err := r.Worktree()
	checkErr(err)
	fmt.Println(Gray("Checking working tree..."))

	s, err := w.Status()
	checkErr(err)

	hasStagingFiles := false
	for _, status := range s {
		if status.Staging != 32 && status.Staging != 63 {
			hasStagingFiles = true
		}
	}

	if hasStagingFiles {
		username, err := gitconfig.Username()
		email, err := gitconfig.Email()
		fmt.Println(Gray("Changes added, Preparing commit..."))

		_, err = w.Commit(message, &git.CommitOptions{
			Author: &object.Signature{
				Name:  username,
				Email: email,
				When:  time.Now(),
			},
		})
		checkErr(err)
		lastCommit := "Last commit: " + message
		fmt.Println(Gray(lastCommit))
		fmt.Println(Green("Done"))
	} else {
		checkErr(errors.New("No changes added to commit"))
	}

	return
}
