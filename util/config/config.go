package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Yaml(path string, cfg any) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(f, &cfg); err != nil {
		return err
	}

	return nil
}
