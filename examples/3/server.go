package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type User struct {
	ID    int    `json:"id", xml:"id"`
	Name  string `json:"name", xml:"name"`
	Email string `json:"email", xml:"email"`
}

func json2xml(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = xml.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
}

func xml2json(w http.ResponseWriter, r *http.Request) {
	var user User
	err := xml.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func main() {
	http.HandleFunc("/json2xml", json2xml)
	http.HandleFunc("/xml2json", xml2json)
	http.ListenAndServe(":80", nil)
}
