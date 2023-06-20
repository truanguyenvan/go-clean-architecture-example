package checks

import (
	"os"
	"regexp"
)

type EnvCheck struct {
	EnvVariable string
	Regex       string
}

func NewEnvCheck(EnvVariable string) EnvCheck {
	return EnvCheck{
		EnvVariable: EnvVariable,
		Regex:       "",
	}
}

func (e EnvCheck) SetRegexValidator(Regex string) {
	e.Regex = Regex
}

func (e EnvCheck) Pass() bool {
	envValue := os.Getenv(e.EnvVariable)
	if envValue == "" {
		return false
	}
	if e.Regex != "" {
		matched, _ := regexp.MatchString(e.Regex, envValue)
		return matched
	}
	return true
}

func (e EnvCheck) Name() string {
	return "Environmental variable \"" + e.EnvVariable + "\""
}
