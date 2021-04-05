package main

type Shard struct {
	id int
	name string
}

type ShardConfig struct {
	shards []Shard
}

