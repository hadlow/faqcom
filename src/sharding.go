package main

import (
	"gopkg.in/yaml.v2"
}

type Shard struct {
	id int
	name string
}

type ShardConfig struct {
	shards []Shard
}

func loadConfig() error  {
	config := ShardConfig{}

	data, err := ioutil.ReadFile("./config.yaml")

	err := yaml.Unmarshal(data, &config)

	if err != nil {
		return err
	}

	return nil
}
