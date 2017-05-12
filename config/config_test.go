package config_test

import (
	"reflect"
	"testing"

	"github.com/yaspell/config"
)

func TestNewConfig(t *testing.T) {
	cfgType := reflect.TypeOf(config.NewConfig()).String()
	expected := "*config.Config"

	if cfgType != expected {
		t.Errorf("Expected to get %s, got %s", expected, cfgType)
	}
}
