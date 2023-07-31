package main

import (
	"command-runner/models"
	"os"
	"path"
)

type Finder struct {
	Options models.FinderOptions
}

func NewFinder() *Finder {
	return &Finder{}
}

func (f *Finder) WithOptions(options models.FinderOptions) *Finder {
	f.Options = options
	return f
}

func (f *Finder) FindRecipeFile(name string) (string, error) {
	startPath := f.Options.StartSearchDir

	if startPath == "" {
		startPath = "."
	}

	currentPath := startPath

	for {
		recipesFile := path.Join(currentPath, name)

		if _, err := os.Stat(recipesFile); err == nil {
			return recipesFile, nil
		}

		parentPath := path.Join(currentPath, "..")

		if _, err := os.Stat(parentPath); err == nil {
			currentPath = parentPath
		} else {
			return "", err
		}
	}
}
