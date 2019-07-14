package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func init() {
	Routes()
}

func TestSendJSON(t *testing.T) {
	t.Log("Given the need to test SendJSON endpoint.")
	{
		req, err := http.NewRequest("GET", "/sendjson", nil)
		if err != nil {
			t.Fatal("\tShould be able to create a request.", ballotX, err)
		}
		t.Log("\tShould be able to create a reqeust.", checkMark)

		rw := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rw, req)

		if rw.Code != 200 {
			t.Fatal("\tShould receive \"200\"", ballotX, rw.Code)
		}
		t.Log("\tShould receive \"200\"", checkMark)

		u := struct {
			Name  string
			Email string
		}{}

		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatal("\tShould decode the response.", ballotX)
		}
		t.Log("\tShould decode the response.", checkMark)

		if u.Name == "Reginald Davis" {
			t.Log("\tShould have a name.", checkMark)
		} else {
			t.Error("\tShould have a name.", ballotX, u.Name)
		}

		if u.Email == "fakereggie@test.com" {
			t.Log("\tShould have a email.", checkMark)
		} else {
			t.Error("\tShould have a email.", ballotX, u.Email)
		}
	}
}
