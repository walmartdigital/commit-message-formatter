# CMF Commit Message Formatter

CMF is a simple to use utility to standarize commit messages on projects.

## Getting started

Install via npm, just do `$ npm install -g go-cmf`

Or from Go `go install github.com/walmartdigital/commit-message-formatter`

Once installed, go to your project an run `$ cmf` after add your files to your stage area.

## Flows

CMF have two flows (for now) default and jira

### Default

Running `$ cmf` you get the default flow, you will be prompt for:

    - Type of change you made to your code
    - Module affected by this change
    - Commit message or description of your change

### Jira

Running `$ cmf jira` you get jira flow, this time you will be prompt for:

    - Jira task ID
    - Type of change you made to your code
    - Commit message or description of your change

> If you want to do jira as a defualt flow just create a `.cmf.yaml`in the root of your project:

```
DEFAULT: jira
```

---

## Configurations

It is possible to config CMF as you like, you can change **custom flows, templates or assign default flows**. You can do this using a local file on the root of your porject or setting as global preferences with a file on your Home directory called `.cmf.yaml`.

### DEFAULT

Set a flow as default:

```
DEFAULT: jira
```

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

- JIRA*TASK \_jira task id*
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

Contributions
Use GitHub issues for requests.

I actively welcome pull requests; learn how to contribute.

---

License
CMF is available under the MIT License.
