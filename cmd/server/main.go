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

func main() {
    // Cargar las variables de entorno desde el archivo .env, si existe
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using default environment variables")
    }

    // Obtener las variables de entorno con valores por defecto
    port := getEnv("PORT", "8080")
    dbUser := getEnv("DB_USER", "defaultuser")
    dbPassword := getEnv("DB_PASSWORD", "defaultpassword")
    dbName := getEnv("DB_NAME", "defaultdb")
    dbPort := getEnv("DB_PORT", "5432")

    err := db.InitDB(dbUser, dbPassword, dbName, dbPort)
    if err != nil {
        log.Fatalf("Could not initialize database: %v", err)
    }
    r := mux.NewRouter()
    r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
    r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    r.HandleFunc("/", handlers.HelloAPI).Methods("GET")
    
    fmt.Printf("Server listening on http://localhost:%s\n", port)
    http.ListenAndServe(":"+port, r)
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}