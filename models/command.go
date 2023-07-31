package models

type Command struct {
	Cmd     string
	Args    []string `yaml:"-"`
	EnvVars map[string]string
	WorkDir string
	Timeout Timeout
}

type Timeout struct {
	Graceful int
	Forced   int
}
