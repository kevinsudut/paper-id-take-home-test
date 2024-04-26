package config

import (
	"os"

	"gopkg.in/yaml.v3"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/database"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/http"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/log"
)

type Config struct {
	Log      log.Config      `yaml:"log"`
	Http     http.Config     `yaml:"http"`
	Database database.Config `yaml:"database"`
}

func ReadConfig(file string) (*Config, error) {
	var conf Config

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
