package main

type User struct {
	ID    int    `json:"id", xml:"id"`
	Name  string `json:"name", xml:"name"`
	Email string `json:"email", xml:"email"`
}
