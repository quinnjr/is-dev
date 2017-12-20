// Copyright (c) 2017 Joseph R. Quinn
// SPDX-License-Identifier: MIT

package isdev

import (
	"os"
	"testing"
)

func TestIsDev(t *testing.T) {
	if err := os.Setenv("ENV", "development"); err != nil {
		t.Error(err)
	}
	if ok := IsDev(); !ok { // Should be true
		t.Errorf("Failed to properly detect $ENV=\"development\", os.Setenv = %s, IsDev() = %v", os.Getenv("ENV"), ok)
	}
	// testing strings.ToLower
	if err := os.Setenv("ENV", "DEVELOPMENT"); err != nil {
		t.Error(err)
	}
	if ok := IsDev(); !ok { // Should be true
		t.Errorf("Failed to properly detect $ENV=\"development\", os.Setenv = %s, IsDev() = %v", os.Getenv("ENV"), ok)
	}
	if err := os.Setenv("ENV", "DEV"); err != nil {
		t.Error(err)
	}
	if ok := IsDev(); !ok { // Should be true
		t.Errorf("Failed to properly detect $ENV=\"dev\", os.Setenv = %s, IsDev() = %v", os.Getenv("ENV"), ok)
	}
}

func TestIsNotDev(t *testing.T) {
	// Testing non-development $ENV
	if err := os.Setenv("ENV", "production"); err != nil {
		t.Error(err)
	}
	if ok := IsDev(); ok { // Should not be true
		t.Errorf("Failed to properly detect $ENV=\"production\", os.Setenv = %s, IsDev() = %v", os.Getenv("ENV"), ok)
	}
	if err := os.Unsetenv("ENV"); err != nil { // Should unset
		t.Error(err)
	}
	if ok := IsDev(); ok { // Should not be true
		t.Errorf("Failed to properly detect unset $ENV, os.Setenv = %s, IsDev() = %v", os.Getenv("ENV"), ok)
	}
}

func TestExtendedIsDev(t *testing.T) {
	dev := New(map[string]string{
		"devdev":    "1",
		"DEV":       "true",
		"ISDEV_DEV": "yes",
	})

	if err := os.Setenv("ISDEV_DEV", "yes"); err != nil {
		t.Error(err)
	}
	if ok := dev.IsDev(); !ok {
		t.Error("Expected env `ISDEV_DEV` to be `yes`")
	}
	if err := os.Unsetenv("ISDEV_DEV"); err != nil {
		t.Error(err)
	}

	if err := os.Setenv("devdev", "1"); err != nil {
		t.Error(err)
	}
	if ok := dev.IsDev(); !ok {
		t.Error("Expected env `devdev` to be `1`")
	}
	if err := os.Unsetenv("devdev"); err != nil {
		t.Error(err)
	}

	if err := os.Setenv("DEV", "false"); err != nil {
		t.Error(err)
	}
	if ok := dev.IsDev(); ok {
		t.Error("Expected env `DEV` to be `false`")
	}
	if err := os.Unsetenv("DEV"); err != nil {
		t.Error(err)
	}
}
