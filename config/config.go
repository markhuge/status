// Package config provides config parameter access to the other packages
package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Config models a config file
type Config struct {
	Sites []siteConfig
	Path  string
}

// siteConfig models an individual Site entry
type siteConfig struct {
	URL string
}

// Read reads the config.yaml file
// this acts like a constructor
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
	c.Path = path
	c.validate()
	return c
}

// TODO make a better validator once I'm sure how config should be laid out
func (c *Config) validate() {
	if len(c.Sites) < 1 {
		panic("No valid config found at " + c.Path)
	}
}
