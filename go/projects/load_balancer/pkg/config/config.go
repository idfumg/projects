package config

import (
	"io"

	"gopkg.in/yaml.v2"
)

type Service struct {
	Name     string   `yaml:"name"`
	Matcher  string   `yaml:"matcher"`
	Strategy string   `yaml:"strategy"`
	Replicas []string `yaml:"replicas"`
}

type Config struct {
	Services []*Service `yaml:"services"`
}

func New(reader io.Reader) (*Config, error) {
	if buf, err := io.ReadAll(reader); err != nil {
		return nil, err
	} else {
		config := &Config{}
		err = yaml.Unmarshal(buf, config)
		return config, err
	}
}
