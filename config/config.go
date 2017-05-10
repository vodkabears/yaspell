package config

// Config stores settings
type Config struct {
	Lang       string
	Format     string
	Options    Bitmask
	Dictionary Dictionary
}

// NewConfig creates new configuration with default settings
func NewConfig() *Config {
	return &Config{
		Lang:   "ru,en",
		Format: "plain",
	}
}
