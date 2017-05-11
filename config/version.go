package config

import (
	"fmt"
	"os"
)

// App version
const version = "0.0.1"

// Version implements flag.Value interface to show the version
type Version bool

func (Version) String() string {
	return version
}

// IsBoolFlag defines a boolean flag
func (v *Version) IsBoolFlag() bool {
	return true
}

// Set shows version on triggering of the flag
func (*Version) Set(string) error {
	fmt.Println(version)
	os.Exit(0)

	return nil
}
