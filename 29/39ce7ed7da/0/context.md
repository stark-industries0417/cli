# Session Context

## User Prompts

### Prompt 1

Implement the following plan:

# Detect External Hook Managers and Warn Users

## Context

PR #355 fixed hook installation to use `git rev-parse --git-path hooks`, which correctly resolves `core.hooksPath` (e.g., Husky's `.husky/_/`). However, Husky regenerates `.husky/_/` on every `npm install` (via the `prepare` script), so Entire's hooks get overwritten.

**Goal:** Detect when an external hook manager is in use, warn the user during `entire enable`, and provide copy-paste instructions for a d...

### Prompt 2

commit this

### Prompt 3

[Request interrupted by user for tool use]

### Prompt 4

move to another branch before committing

### Prompt 5

lets run the manual tests

### Prompt 6

Lefthook looks for a configuration file using the pattern {.,}lefthook.{yml,yaml,json,toml} DeepWiki. So all valid main config files are:

lefthook.yml
lefthook.yaml
lefthook.json
lefthook.toml
.lefthook.yml
.lefthook.yaml
.lefthook.json
.lefthook.toml

Plus the local override variants (same pattern with -local):

lefthook-local.yml / .lefthook-local.yml (and .yaml, .json, .toml)

### Prompt 7

why is HookManager and DetectHookManagers exported (capitalized)?

