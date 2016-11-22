package archon

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type test_post struct {
	Service string `json:"service"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	dec := json.NewDecoder(r.Body)
	var t test_post
	for {
		if err := dec.Decode(&t); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	res, _ := json.Marshal(t)
	io.WriteString(w, string(res))
}
