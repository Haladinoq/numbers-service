package config

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config parses all env vars. If not set uses default values.
type Config struct {
	ServerPort      int    `yaml:"SERVER_PORT"`
	DBHost          string `yaml:"DB_HOST"`
	DBPort          int    `yaml:"DB_PORT"`
	DBName          string `yaml:"DB_NAME"`
	DBUser          string `yaml:"DB_USER"`
	DBPassword      string `yaml:"DB_PASSWORD"`
	DBSSLMode       string `yaml:"DB_SSLMODE"`
	SwaggerUser     string `yaml:"SWAGGER_USER"`
	SwaggerPassword string `yaml:"SWAGGER_PASSWORD"`
}

// LoadConfig load and parses yaml file.
func LoadConfig(file string) (*Config, error) {
	if file == "" {
		return nil, errors.New("empty config filename")
	}
	// #nosec
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	if err := f.Close(); err != nil {
		return nil, err
	}
	var c *Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		return nil, err
	}
	return c, nil
}
