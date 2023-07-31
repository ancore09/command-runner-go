package models

type Recipe struct {
	Name         string `yaml:"-"`
	Commands     []Command
	Parameters   []Parameter
	WorkDir      string
	Shell        string
	ExportParams bool `yaml:"export_params"`
}

type Parameter struct {
	Name  string
	Value string
}
