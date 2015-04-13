package client

import "fmt"

// Config.
type Config struct {
	Connection   string `yaml:"connection"`
	Table        string `yaml:"table"`
	Column       string `yaml:"column"`
	MaxOpenConns int    `yaml:"max_open_connections"`
	Verbose      bool   `yaml:"verbose"`
}

// Validate the configuration and apply defaults.
func (c *Config) Validate() error {
	if c.Connection == "" {
		return fmt.Errorf(`"connection" required`)
	}

	if c.Table == "" {
		return fmt.Errorf(`"table" required`)
	}

	if c.Column == "" {
		return fmt.Errorf(`"column" required`)
	}

	if c.MaxOpenConns == 0 {
		c.MaxOpenConns = 10
	}

	return nil
}
