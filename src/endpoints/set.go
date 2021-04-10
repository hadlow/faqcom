package endpoints

// Local imports
import (
	"fmt"
	"log"
	"net/http"
)

func (e *Endpoint) Set(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	key := r.Form.Get("key")
	value := r.Form.Get("data")

	shard := e.getShard(key)

	err := e.DB.Set(key, []byte(value))

	if shard != e.ShardId {
		e.Route(w, r, shard)

		return
	}

	if err != nil {
		log.Fatal("Error setting value")
	}

	fmt.Println("Key value set")
}
