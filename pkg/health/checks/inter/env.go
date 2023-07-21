package checks

import (
	"fmt"
	"go-clean-architecture-example/pkg/health"
	"os"
	"regexp"
	"sync"
	"time"
)

type Env struct {
	EnvVariable string
	Regex       string
}

func NewEnvChecker(envVariable, regex string) *Env {
	return &Env{
		EnvVariable: envVariable,
		Regex:       regex,
	}
}

func (e Env) Check(name string, result *healthcheck.ApplicationHealthDetailed, wg *sync.WaitGroup, checklist chan healthcheck.Integration) {
	defer (*wg).Done()
	var (
		start        = time.Now()
		myStatus     = true
		errorMessage = ""
	)

	envValue := os.Getenv(e.EnvVariable)
	if envValue == "" {
		myStatus = false
		result.Status = false
		errorMessage = fmt.Sprintf("%s not found", envValue)
	}
	if e.Regex != "" {
		matched, _ := regexp.MatchString(e.Regex, envValue)
		if !matched {
			myStatus = false
			result.Status = false
			errorMessage = fmt.Sprintf("%s pattern doesn't match any environment variable", e.Regex)
		}
	}

	checklist <- healthcheck.Integration{
		Name:         name,
		Kind:         "Environmental variable",
		Status:       myStatus,
		ResponseTime: time.Since(start).Seconds(),
		Error:        errorMessage,
	}
}
