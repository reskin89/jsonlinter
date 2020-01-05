package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var dat interface{}

func linter(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&dat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	encoder := json.NewEncoder(w)

	encoder.SetIndent("", "  ")

	err = encoder.Encode(dat)
	log.Println(fmt.Sprintf("[ %q ] -- [ %d ] -- [ %q ]", r.Method, r.ContentLength, r.RemoteAddr))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func main() {
	http.HandleFunc("/", linter)

	log.Fatal(http.ListenAndServe(":8400", nil))
}
