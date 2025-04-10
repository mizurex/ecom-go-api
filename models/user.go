package models

type User struct {
	ID    int    `json:"id"`
	NAME  string `json:"name"`
	GMAIL string `json:"gmail"`
	PASS  string `json:"password"`
}
