package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func main() {
	http.HandleFunc("/json2xml", json2xml)
	http.HandleFunc("/xml2json", xml2json)
	http.ListenAndServe(":80", nil)
}
