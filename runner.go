package main

import (
	"command-runner/models"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Runner struct {
	Recipes map[string]models.Recipe
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) WithRecipes(recipes map[string]models.Recipe) *Runner {
	r.Recipes = recipes
	return r
}

func (r *Runner) RunRecipe(name string, args []string) {
	recipe := r.Recipes[name]

	argumentProcessor := NewArgumentProcessor()
	argumentProcessor.ProcessRecipe(&recipe, args)

	for _, command := range recipe.Commands {
		argumentProcessor.ProcessCommand(&command, recipe)

		color.New(color.FgHiMagenta).Add(color.Bold).Print("[Command] ")
		color.New(color.FgHiCyan).Println(command.Cmd)

		var cmd *exec.Cmd
		if recipe.Shell == "" {
			cmdArgs := strings.Fields(command.Cmd)
			cmd = exec.Command(cmdArgs[0], cmdArgs[1:]...)
		} else {
			cmd = exec.Command(recipe.Shell, "-c", command.Cmd)
		}

		if recipe.ExportParams {
			for _, parameter := range recipe.Parameters {
				cmd.Env = append(cmd.Env, parameter.Name+"="+parameter.Value)
			}
		}

		cmd.Dir = recipe.WorkDir

		if command.WorkDir != "" {
			cmd.Dir = command.WorkDir
		}

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()

		color.New(color.FgHiMagenta).Add(color.Bold).Print("[Pid] ")
		color.New(color.FgHiCyan).Println(cmd.Process.Pid)

		err := cmd.Wait()

		if exiterr, ok := err.(*exec.ExitError); ok {
			color.New(color.FgRed).Add(color.Bold).Println("[Error] " + strconv.Itoa(exiterr.ExitCode()))
		} else {
			color.New(color.FgHiMagenta).Add(color.Bold).Print("[Exit] ")
			color.New(color.FgHiCyan).Println(0)
		}

		/*go func() {
			scanner := bufio.NewScanner(stderr)
			scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				m := scanner.Text()
				fmt.Println("ERR| " + m)
			}
			cmd.Wait()
		}()

		scanner := bufio.NewScanner(stdout)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			m := scanner.Text()
			fmt.Println("OUT| " + m)
		}
		err := cmd.Wait()
		if err != nil {
			println(err)
		}*/
	}
}
