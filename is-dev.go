// Copyright (c) 2017 Joseph R. Quinn
// SPDX-License-Identifier: MIT

package isdev

import (
	"os"
	"strings"
)

// IsDev reads the local environment for the key "ENV" and
// returns whether the value of the key is set to "development"
// or "dev"
func IsDev() bool {
	env, ok := os.LookupEnv("ENV")
	if !ok {
		return false
	}
	env = strings.ToLower(env)
	if env == "development" || env == "dev" {
		return true
	}
	return false
}

// ExtendedVars creates a struct with specified variables to look
// for in the os' environment to determine if the program is running // in development environment.
type ExtendedVars struct {
	vars map[string]string
}

// New creates a new is-dev struct with specified custom vars set
func New(vars map[string]string) *ExtendedVars {
	return &ExtendedVars{
		vars: vars,
	}
}

// IsDev reads the local environment for the key "ENV" and
// returns whether the value of the key is set in the environment
// to indicate that the environment is a development environment
func (e *ExtendedVars) IsDev() bool {
	for key, val := range e.vars {
		env, ok := os.LookupEnv(key)
		if !ok {
			continue
		}
		if env == val {
			return true
		}
	}
	return false
}
