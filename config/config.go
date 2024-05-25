package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Users []User `yaml:"users"`
}

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func Parse(path *string) (*Config, error) {
	data, err := os.ReadFile(*path)
	if err != nil {
		return nil, err
	}
	config := Config{}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
