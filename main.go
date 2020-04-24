package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Select{
		Label: "What type is your commit?",
		Items: KnownTypes,
		Templates: &promptui.SelectTemplates{
			Active:   fmt.Sprintf("%s {{ .Name | underline}} - {{ .Description }}", promptui.IconSelect),
			Inactive: "  {{ .Name }}",
			Selected: fmt.Sprintf(`{{ "%s" | green }} {{ .Name }} - {{ .Description }}`, promptui.IconGood),
		},
		Size: len(KnownTypes),
	}

	index, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	scopePrompt := promptui.Prompt{
		Label:     "Scope",
		AllowEdit: true,
		Default:   "",
	}

	scope, err := scopePrompt.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	breakingPrompt := promptui.Prompt{
		Label:     "Breaking",
		Default:   "N",
		IsConfirm: true,
	}

	breakingString, err := breakingPrompt.Run()
	if err != nil && err != promptui.ErrAbort {
		fmt.Printf("%v\n", err)
		return
	}

	breaking := strings.ToLower(breakingString) == "y"
	commitType := KnownTypes[index].Name

	sb := strings.Builder{}
	sb.WriteString(commitType)
	if breaking {
		sb.WriteRune('!')
	}
	if scope != "" {
		sb.WriteString(fmt.Sprintf("(%s)", scope))
	}
	sb.WriteRune(':')
	if breaking {
		sb.WriteString("\n\nBREAKING CHANGE:")
	}

	message := sb.String()
	args := []string{"commit", "-m", message, "--edit"}
	cmd := exec.Command("git", args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Could not commit: %v\n", err)
	}
}
