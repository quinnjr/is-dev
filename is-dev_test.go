// Copyright (c) 2017 Joseph R. Quinn
// SPDX-License-Identifier: MIT

package isdev

import (
	"os"
	"testing"
)

func TestIsDev(t *testing.T) {
	err := os.Setenv("ENV", "development")
	if err != nil {
		t.Error(err)
	}
	ok := IsDev()
	if !ok { // Should be true
		t.Errorf("Failed to properly detect $ENV=\"development\", os.Setenv = %s, IsDev() = %v", os.Getenv("ENV"), ok)
	}
	err = os.Setenv("ENV", "DEVELOPMENT") // testing string.ToLower
	if err != nil {
		t.Error(err)
	}
	ok = IsDev()
	if !ok { // Should be true
		t.Errorf("Failed to properly detect $ENV=\"development\", os.Setenv = %s, IsDev() = %v", os.Getenv("ENV"), ok)
	}
	err = os.Setenv("ENV", "DEV") // testing strings.ToLower and shorthand development set as $ENV
	if err != nil {
		t.Error(err)
	}
	ok = IsDev()
	if !ok { // Should be true
		t.Errorf("Failed to properly detect $ENV=\"dev\", os.Setenv = %s, IsDev() = %v", os.Getenv("ENV"), ok)
	}
}

func TestIsNotDev(t *testing.T) {
	err := os.Setenv("ENV", "production") // Testing non-development $ENV
	if err != nil {
		t.Error(err)
	}
	ok := IsDev()
	if ok { // Should not be true
		t.Errorf("Failed to properly detect $ENV=\"production\", os.Setenv = %s, IsDev() = %v", os.Getenv("ENV"), ok)
	}
	err = os.Unsetenv("ENV") // Unsetting $ENV
	if err != nil {
		t.Error(err)
	}
	ok = IsDev()
	if ok { // Should not be true
		t.Errorf("Failed to properly detect unset $ENV, os.Setenv = %s, IsDev() = %v", os.Getenv("ENV"), ok)
	}
}
