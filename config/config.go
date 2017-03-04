// Package config provides config parameter access to the other packages
package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Config models a config file
type Config struct {
	Sites []siteConfig
}

// siteConfig models an individual Site entry
type siteConfig struct {
	URL string
}

// Read reads the config.yaml file
func Read(path string) *Config {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	c := new(Config)
	err = yaml.Unmarshal(file, &c)
	if err != nil {
		panic(err)
	}
	return c
}
