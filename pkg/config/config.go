package config

// Env returns the value of the environment variable named by the key.
type Env string

// Config is the configuration struct
type Config struct {
	Port *string `env:"PORT"`
	Env  string  `env:"ENV"`
}

// GetEnv returns the current environment
func (c *Config) GetEnv() Env {
	return Env(c.Env)
}

// Instance is the global configuration
var Instance *Config
