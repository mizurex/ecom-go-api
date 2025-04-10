package database

import(
	"ecom/models"
	
	"database/sql"
)

func InsertUser(db *sql.DB, s models.User) error {
	query := "INSERT INTO users (name,gmail,pass) VALUES (?,?,?)"
	_, err := db.Exec(query,s.NAME,s.GMAIL,s.PASS)
	return err
}