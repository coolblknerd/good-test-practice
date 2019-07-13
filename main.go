package main

import "net/http"

func main() {
	Routes()

	http.ListenAndServe(":4000", nil)
}
