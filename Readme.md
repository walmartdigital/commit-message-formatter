# CMF Commit Message Formatter - V3

CMF is a simple-to-use utility to standarized commit messages on projects.

## Migrating from V2

- Installation via npm command is now deprecated and no longer maintained on **V3**.
- You can still be using your `.cmf.yaml` file. Now you can extend this file with new attributes.

## Major changes

- .cmf.yaml file is no longer required if you want to use a simple flow.
- The default flow of v3 is now strongly forced to use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).
- CMF binary is renamed to git-cmf, so now you can use it directly from git as `git cmf` command.
- CMF is now available from Homebrew.
- `cmf init` is cleaner now, it just creates a .cmf.yaml file with 
a simple flow to let you customize as you want.

## Getting Started

### Install CMF

- Using go install: `go install github.com/walmartdigital/commit-message-formatter`
- Using Brew:
  - `brew tap walmartdigital/homebrew-git-cmf`
  - `brew install homebrew-git-cmf`
- Download binaries from Github

### Initialize project (optional)

If you want to customize your flow, you can run `git cmf init`,
this command will create a .cmf.yaml file with the default flow.

Then you can change the flow as you want.

### Variables

CMF has inner variables and you can access it throw templates using `{{}}`:

- {{BRANCH_NAME}} it will print the current branch name of your repository

Additionally, you can include external environment variables using the **ENV** 
block described on the template.

### Extending

It is possible to config CMF as you like, you can change
 **custom flows, templates, or assign default flows**. You can do this using 
 a local file on the root of your project or set as global 
 preferences with a file on your Home directory called `.cmf.yaml`.

### Template Structure

A `.cmf.yaml` file is composite by 3 main blocks:
  - ENV
  - PROMPT
  - TEMPLATE

#### ENV

It is a list of environment variables names, that later are mappings to be 
accessible from other blocks.

```
ENV:
  - ENVIRONMENT_1
  - ENVIRONMENT_2
  ...
  - ENVIRONMENT_10
```

#### PROMPT

Describe an input flow.

You can create your custom flows using this configuration attribute.` 
Every `KEY` attribute is mapping as  a variable within the flow.

Prompt accepts two kinds of questions:

- Single question:
  - KEY _variable name_
  - LABEL _prompt title_

- Select question:
  - KEY _variable name_
  - LABEL _prompt title_
  - OPTIONS _list of options_
    - VALUE _variable value_
    - DESC _variable description_

```
PROMPT:
  - KEY: "CHANGE"
    LABEL: "Select the type of change:"
    OPTIONS:
      - VALUE: "feature"
        DESC: "A new feature"
      - VALUE: "fix"
        DESC: "A Bug fix"
      - VALUE: "update"
        DESC: "An update code change (moving or split code)"
      - VALUE: "docs"
        DESC: "Documentation only changes"
      - VALUE: "style"
        DESC: "Small changes of code style"
      - VALUE: "test"
        DESC: "Add, change or update test code"
  - KEY: "MODULE"
    LABEL: "Affected module:"
  - KEY: "MESSAGE"
    LABEL: "Commit message:"
TEMPLATE: "{{CHANGE}}({{MODULE}}): {{MESSAGE}}"
```

#### TEMPLATE

Defines the way the commit message will be formatted using variables described on .cmf.yaml file.

```
TEMPLATE: "{{CHANGE}}({{MODULE}}): {{MESSAGE}}"
```
---

## Contributions

Use GitHub issues for requests.

I actively welcome pull requests; learn how to [contribute](CONTRIBUTING.md).

---

## License

CMF is available under the MIT License.
