# CMF Commit Message Formatter - V2

CMF is a simple to use utility to standarize commit messages on projects.

## Getting started

Install via npm, just do `$ npm install -g go-cmf`

Or from Go `go install github.com/walmartdigital/commit-message-formatter`

Once installed, go to your project an run `$ cmf init` and select one of the flows, it will create a `.cmf.yaml`file on your project with your selected flow.

## Flows

CMF have three flows (for now) default, Jira and custom

### Default

Running `$ cmf init` and select default, you will get the default flow and prompted for:

    - Type of change you made to your code
    - Module affected by this change
    - Commit message or description of your change

### Jira

Running `$ cmf init` and select Jira, you will get the jira flow, this time you will be prompt for:

    - Jira task ID
    - Type of change you made to your code
    - Commit message or description of your change

### Custom

Running `$ cmf init`and select custom, you will get the custom flow, this time it will create a `.cmf.yaml`file with default flow but with annotations of how change it.

---

## Variables

CMF have inner variables and you can access it throw templates using `{{}}`:

- {{BRANCH_NAME}} it will print the current branch name of your repository

## Configurations

It is possible to config CMF as you like, you can change **custom flows, templates or assign default flows**. You can do this using a local file on the root of your porject or setting as global preferences with a file on your Home directory called `.cmf.yaml`.

### TEMPLATE

Set a template string for commit messages.

#### Default flow

Default template `{{CHANGE}}({{MODULE}}): {{MESSAGE}}`. You can use this variables:

- CHANGE _type of change: feature, fix, update_
- MODULE _module affected_
- MESSAGE _commit message_

```
TEMPLATE: "{{CHANGE}}({{MODULE}}): {{MESSAGE}}"
```

#### Jira flow

Default template `{{JIRA-TASK}} ({{CHANGE}}): {{MESSAGE}}`. You can use this variables:

- JIRA*TASK \_jira task id, by default {{BRANCH_NAME}}*
- CHANGE _type of change: feature, fix, update_
- MESSAGE _commit message_

```
TEMPLATE: "{{JIRA-TASK}} ({{CHANGE}}): {{MESSAGE}}"
```

### PROMPT

You can create your custom flows using this configuration attribute.

Prompt accept two types of prompts:

- Single question:
  - KEY _variable name_
  - LABEL _prompt title_
- Select:
  - KEY _variable name_
  - LABEL _prompt title_
  - OPTIONS _list of options_
    - VALUE _variable value_
    - DESC _variable description_

_default .cmf.yaml sample file_

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

---

## Contributions

Use GitHub issues for requests.

I actively welcome pull requests; learn how to contribute.

---

## License

CMF is available under the MIT License.
