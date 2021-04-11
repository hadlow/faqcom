package endpoints

// Local imports
import (
	"fmt"
	"io"
	"strconv"
	"net/http"
	"hash/fnv"
)

import (
	"github.com/hadlow/genomdb/src/database"
	"github.com/hadlow/genomdb/src/types"
)

type Endpoint struct {
	DB *database.Database
	ShardId int
	Shards []types.Shard
}

func New(DB *database.Database, shardId int, shards []types.Shard) *Endpoint {
	return &Endpoint {
		DB: DB,
		ShardId: shardId,
		Shards: shards,
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

	shardId := int(hash.Sum64() % uint64(len(e.Shards)))

	return shardId
}

func (e *Endpoint) Route(w http.ResponseWriter, r *http.Request, shard int) {
	response, err := http.Get("http://" + e.Shards[shard].Host + ":" + strconv.Itoa(e.Shards[shard].Port) + r.RequestURI)

	if err != nil {
		w.WriteHeader(500)
		fmt.Println("Error at: " + err.Error())

		return
	}

	defer response.Body.Close()

	io.Copy(w, response.Body)
}
