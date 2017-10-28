// Copyright (c) 2017 Joseph R. Quinn
// SPDX-License-Identifier: MIT

package isdev

import (
  "os"
  "strings"
)

// IsDev reads the local environment for the key "ENV" and returns whether
// the value of the key is set to "development" or "dev"
func IsDev() bool {
  env, ok := os.LookupEnv("ENV")
  if !ok {
    return false
  }
  if strings.ToLower(env) != "development" || strings.ToLower(env) != "dev" {
    return false
  }
  return true
}
