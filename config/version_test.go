package config_test

import (
	"regexp"
	"testing"

	"github.com/VodkaBears/yaspell/config"
)

func TestVersionString(t *testing.T) {
	var v config.Version
	re := regexp.MustCompile(`^\d+\.\d+\.\d+$`)

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
