package main

import (
	"io/ioutil"

	"github.com/hadlow/genomdb/src/types"
	
	"gopkg.in/yaml.v2"
)

type Config struct {
	Database string `yaml:"database"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Shards []types.Shard `yaml:"shards"`
}

func loadConfig(path string) (Config, error) {
	// Load the config file and return as object
	var config Config

	data, errRead := ioutil.ReadFile(path)

	if errRead != nil {
		return config, errRead
	}

	errParse := yaml.Unmarshal(data, &config)

	if errParse != nil {
		return config, errParse
	}

	return config, nil
}
