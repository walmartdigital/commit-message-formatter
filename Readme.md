# CMF Commit Message Formatter

CMF is a simple to use utility to standarize commit messages on projects.

## Getting started

Install via npm, just do `$ npm install -g go-cmf`

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
