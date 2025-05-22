package main

import (
	"context"

	"github.com/sancheschris/user-auth-system/internal/database"
)
func main() {

	db, err := database.ConnectMongoDB("mongodb://root:secret@localhost:27017/?authSource=admin")

	if err != nil {
		panic(err)
	}
	defer db.Client().Disconnect(context.Background())
	
	db.Client()
}