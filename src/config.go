package main

import (
	"io/ioutil"
	"fmt"
)

import (
	"gopkg.in/yaml.v2"
)

func loadConfig() error  {
	config := ShardConfig{}

	data, fileErr := ioutil.ReadFile("./config.yml")

	yamlErr := yaml.Unmarshal(data, &config)

	if fileErr != nil || yamlErr != nil {
		return yamlErr
	}

	fmt.Printf("--- t dump:\n%s\n\n", string(config.shards[0].id))

	return nil
}
