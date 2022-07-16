package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Addr     string `yaml:"addr"`
}

type Server struct {
	Port string `yaml:"port"`
}

func New(filepath string) (*Config, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.Wrapf(err, "read file %s", filepath)
	}
	var cfg Config
	if err = yaml.Unmarshal(bytes, &cfg); err != nil {
		return nil, errors.Wrap(err, "unmarshal yaml")
	}
	return &cfg, nil
}
