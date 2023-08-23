package models

type Command struct {
	Cmd          string
	Args         []string `yaml:"-"`
	EnvVars      map[string]string
	WorkDir      string
	Timeout      Timeout
	AllowFailure bool   `yaml:"allow_failure"`
	OutputMode   string `yaml:"output_mode"`
}

type Timeout struct {
	Graceful int
	Forced   int
}
