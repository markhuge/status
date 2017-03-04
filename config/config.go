// Package config provides config parameter access to the other packages
package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

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
		log.Fatal(err)
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
		fmt.Printf("No valid config found at %s", c.Path)
		os.Exit(1)
	}
}
