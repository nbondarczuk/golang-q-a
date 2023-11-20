package main

import (
	"fmt"
	"net/http"
)

func json(w http.ResponseWriter, req *http.Request) {
}

func json(w http.ResponseWriter, req *http.Request) {
}

func main() {
    http.HandleFunc("/json", json)
    http.HandleFunc("/xml", xml)
    http.ListenAndServe(":80", nil)
}
