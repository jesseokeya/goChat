package main

import (
	"github.com/jesseokeya/goChat/database"
)

var (
	api API
)

// API creates an api struct which connects logics with the database
type API struct {
	db database.DB
}

func main() {
	db := api.db
	s := db.ConnectToDatabase()
}
