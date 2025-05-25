package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sancheschris/user-auth-system/internal/auth"
	"github.com/sancheschris/user-auth-system/internal/config"
	"github.com/sancheschris/user-auth-system/internal/database"
	jwtmiddleware "github.com/sancheschris/user-auth-system/internal/middleware"
)
func main() {

	cfg := config.LoadConfig()

	db, err := database.ConnectMongoDB(cfg.MongoURI)

	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer db.Client().Disconnect(context.Background())

	repo := auth.NewMongoUserRepository(db)
	service := auth.NewAuthService(repo)
	handler := &auth.AuthHandler{AuthService: service}
	
	r := mux.NewRouter()

	r.HandleFunc("/register", handler.Register).Methods("POST")
	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.Handle("/protected", jwtmiddleware.JWTAuth(http.HandlerFunc(handler.ProtectedEndpoint))).Methods("POST")

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}