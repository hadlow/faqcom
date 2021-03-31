package endpoints

// Local imports
import (
	"fmt"
	"log"
	"net/http"
)

func (e *Endpoint) Get(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	key := r.Form.Get("key")

	value, err := e.DB.Get(key)

	if err != nil {
		log.Fatal("Error getting value")
	}

	fmt.Println(value)
}
