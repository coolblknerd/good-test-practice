package main

import (
	"encoding/json"
	"net/http"
)

// Routes handles and routes that are passed into it
func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

// SendJSON handles sending json
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Reginald Davis",
		Email: "fakereggie@test.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
