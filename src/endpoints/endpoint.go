package endpoints

// Local imports
import (
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
