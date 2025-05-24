package main

import (
	"context"
	"log"
	"net/http"

	"github.com/sancheschris/user-auth-system/internal/database"
)
func main() {

	db, err := database.ConnectMongoDB("mongodb://root:secret@localhost:27017/?authSource=admin")

	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer db.Client().Disconnect(context.Background())

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}