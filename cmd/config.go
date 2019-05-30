package main

import (
	"fmt"
	"io/ioutil"

	"github.com/duyanghao/GWorkerPools/worker"
	"gopkg.in/yaml.v2"
)

type Config struct {
	PrintWorker *worker.PrintWorkerConfig `yaml:"print_worker,omitempty"`
	// TODO: other configuration ...
}

// validate the configuration
func (c *Config) validate() error {
	if err := c.PrintWorker.Validate(); err != nil {
		return err
	}
	// TODO: other configuration validate ...
	return nil
}

// LoadConfig parses configuration file and returns
// an initialized Settings object and an error object if any. For instance if it
// cannot find the configuration file it will set the returned error appropriately.
func LoadConfig(path string) (*Config, error) {
	c := &Config{}
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read configuration file: %s,error: %s", path, err)
	}
	if err = yaml.Unmarshal(contents, c); err != nil {
		return nil, fmt.Errorf("Failed to parse configuration,error: %s", err)
	}
	if err = c.validate(); err != nil {
		return nil, fmt.Errorf("Invalid configuration,error: %s", err)
	}
	return c, nil
}
