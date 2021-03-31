package endpoints

// Local imports
import (
	"net/http"
)

import (
	"github.com/hadlow/genomdb/src/database"
)

type Endpoint struct {
	DB *database.Database
}

func New(DB *database.Database) *Endpoint {
	return &Endpoint {
		DB: DB,
	}
}

func (e *Endpoint) Serve(address string, port string) error {
	http.HandleFunc("/g", e.Get)
	http.HandleFunc("/s", e.Set)

	return http.ListenAndServe(address + ":" + port, nil)
}
