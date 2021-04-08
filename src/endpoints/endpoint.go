package endpoints

// Local imports
import (
	"strconv"
	"net/http"
	"hash/fnv"
)

import (
	"github.com/hadlow/genomdb/src/database"
)

type ShardInfo struct {
	Id int
	Count int
}

type Endpoint struct {
	DB *database.Database
	Shard ShardInfo
}

func New(DB *database.Database, shardId int, shardCount int) *Endpoint {
	shard := ShardInfo{
		Id: shardId,
		Count: shardCount,
	}

	return &Endpoint {
		DB: DB,
		Shard: shard,
	}
}

func (e *Endpoint) Serve(address string, port int) error {
	http.HandleFunc("/g", e.Get)
	http.HandleFunc("/s", e.Set)

	return http.ListenAndServe(address + ":" + strconv.Itoa(port), nil)
}

func (e *Endpoint) getShard(key string) int {
	hash := fnv.New64()
	hash.Write([]byte(key))

	shardId := int(hash.Sum64() % uint64(e.Shard.Count))

	return shardId
}
