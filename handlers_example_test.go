package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func ExampleSendJSON() {
	req, _ := http.NewRequest("GET", "/sendjson", nil)
	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, req)

	var u struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
		log.Println(":ERROR", err)
	}

	fmt.Println(u)
	// Output:
	// {Reginald Davis fakereggie@test.com}

}
