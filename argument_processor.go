package main

import (
	"command-runner/models"
	"regexp"
	"strings"
)

type ArgumentProcessor struct {
}

func NewArgumentProcessor() *ArgumentProcessor {
	return &ArgumentProcessor{}
}

func (a *ArgumentProcessor) ProcessRecipe(recipe *models.Recipe, args []string) {
	for i := range recipe.Parameters {
		if args[i] != "-" {
			recipe.Parameters[i].Value = args[i]
		}
	}
}

func (a *ArgumentProcessor) ProcessCommand(command *models.Command, recipe models.Recipe) {
	tokens := strings.Split(command.Cmd, " ")

	for i, token := range tokens {
		var regex = regexp.MustCompile(`{{(.*)}}`)
		var matches = regex.FindStringSubmatch(token)

		if len(matches) > 0 {
			var parameterName = matches[1]

			for i2 := range recipe.Parameters {
				if recipe.Parameters[i2].Name == parameterName {
					parameter := recipe.Parameters[i2]
					tokens[i] = strings.ReplaceAll(tokens[i], "{{"+parameterName+"}}", parameter.Value)
				}
			}
		}
	}

	command.Cmd = strings.Join(tokens, " ")
}
