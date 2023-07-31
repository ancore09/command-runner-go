package main

import (
	"command-runner/models"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"path/filepath"
)

func main() {
	finder := NewFinder().WithOptions(models.FinderOptions{
		StartSearchDir:   ".",
		FallbackToParent: true,
	})

	file, err := finder.FindRecipeFile("recipes.yaml")
	if err != nil {
		return
	}

	file = path.Clean(file)
	file, _ = filepath.Abs(file)

	color.New(color.FgHiMagenta).Add(color.Bold).Print("[Recipes] ")
	color.New(color.FgHiCyan).Println(file)

	yamlString, err := os.ReadFile(file)

	if err != nil {
		println(err)
	}

	recipes := make(map[string]models.Recipe)

	err = yaml.Unmarshal(yamlString, &recipes)

	if err != nil {
		println(err)
	}

	runner := NewRunner().WithRecipes(recipes)

	runner.RunRecipe(os.Args[1], os.Args[2:])
}
