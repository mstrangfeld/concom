package main

// CommitType is the type of a conventional commit
type CommitType struct {
	Name        string
	Description string
}

// KnownTypes are all known conventional commit types
var KnownTypes = []CommitType{
	{Name: "feat", Description: "A new feature"},
	{Name: "fix", Description: "A bug fix"},
	{Name: "docs", Description: "Documentation only changes"},
	{Name: "build", Description: "Changes that affect the build-system or external dependencies"},
	{Name: "ci", Description: "Changes to CI configuration files and scripts"},
	{Name: "perf", Description: "A code change that improves performance"},
	{Name: "refactor", Description: "A code change that neither fixes a bug nor adds a new feature"},
	{Name: "style", Description: "Changes that do not affect the meaning of the code"},
	{Name: "test", Description: "Adding missing tests or fixing existing ones"},
	{Name: "chore", Description: "Etc. no production code changes"},
}
