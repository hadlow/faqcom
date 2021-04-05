package main

import (
	"io/ioutil"
)

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	Database string `yaml:"database"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Shards []struct {
		Id int `yaml:"id"`
		Name string `yaml:"name"`
	} `yaml:"shards"`
}

func loadConfig() (Config, error) {
	var config Config

	data, errRead := ioutil.ReadFile("./config.yml")

	if errRead != nil {
		return config, errRead
	}

	errParse := yaml.Unmarshal(data, &config)

	if errParse != nil {
		return config, errParse
	}

	return config, nil
}
