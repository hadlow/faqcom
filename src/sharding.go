package main

import (
	"io/ioutil"
	"fmt"
)

type Shard struct {
	id int
	name string
}

type ShardConfig struct {
	shards []Shard
}

