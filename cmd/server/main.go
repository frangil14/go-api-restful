package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/frangil14/go-api-restful/internal/db"
	"github.com/frangil14/go-api-restful/internal/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func main() {
    // Load environment variables from .env
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using default environment variables")
    }

    dbUser := getEnv("DB_USER", "defaultuser")
    dbPassword := getEnv("DB_PASSWORD", "defaultpassword")
    dbName := getEnv("DB_NAME", "defaultdb")
    dbPort := getEnv("DB_PORT", "5432")
    port := getEnv("PORT", "8080")

    // Connect with database
    err := db.InitDB(dbUser, dbPassword, dbName, dbPort)
    if err != nil {
        log.Fatalf("Could not initialize database: %v", err)
    }

    // Add endpoints
    r := mux.NewRouter()
    r.HandleFunc("/", handlers.HelloAPI).Methods("GET")
    r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
    r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", handlers.GetUserById).Methods("GET")
    

    
    fmt.Printf("Server listening on http://localhost:%s\n", port)
    http.ListenAndServe(":"+port, r)
}