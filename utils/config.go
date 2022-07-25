package utils

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Template Template
	Keys     []Key `json:"keys" toml:"keys"`
}

type Template struct {
	Region string `json:"region" toml:"region"`
	Normal string `json:"normal" toml:"normal"`
}

// Key layout format
// format to use
// Prefix - Description
type Key struct {
	Prefix      string `json:"prefix" toml:"prefix"`
	Description string `json:"description" toml:"description"`
}

func (k Key) FilterValue() string {
	return k.Prefix
}

func GenerateConfigFromFile(path string) (c Config, err error) {
	_, err = toml.DecodeFile(path, &c)
	if err != nil {
		return Config{}, err
	}

	return c, err
}

func GenerateConfig(data string) (c Config, err error) {
	_, err = toml.Decode(data, &c)
	if err != nil {
		return Config{}, err
	}

	return c, err
}
