package main

import (
	"io/ioutil"
)

import (
	"github.com/hadlow/genomdb/src/types"
)

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	Database string `yaml:"database"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Shards []types.Shard `yaml:"shards"`
}

func loadConfig(path string) (Config, error) {
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
