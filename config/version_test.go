package config_test

import (
	"regexp"
	"testing"

	"github.com/yaspell/config"
)

func TestVersionString(t *testing.T) {
	re := regexp.MustCompile(`^\d+\.\d+\.\d+$`)
	var v config.Version

	if !re.MatchString(v.String()) {
		t.Error("Expected semver, got", v)
	}
}

func TestVersionIsBoolFlag(t *testing.T) {
	var v config.Version

	if !v.IsBoolFlag() {
		t.Error("Expected to be a bool flag")
	}
}
