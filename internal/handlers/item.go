package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"database/sql"

	"github.com/frangil14/go-api-restful/internal/db"
	"github.com/frangil14/go-api-restful/internal/models"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
    rows, err := db.DB.Query("SELECT id, name, age FROM users")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var users []models.User 
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        users = append(users, user)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var user models.User

    query := `SELECT id, name, age FROM users WHERE id = $1`
    err := db.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Age)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "User not found", http.StatusNotFound)
            return
        }
        http.Error(w, "Error retrieving user from database", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    // Decode request body
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Execute in SQL
    query := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id`
    err := db.DB.QueryRow(query, user.Name, user.Age).Scan(&user.ID)
    if err != nil {
        http.Error(w, "Error inserting user into database", http.StatusInternalServerError)
        return
    }

    // Return JSON format
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func HelloAPI(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "¡Hola, API!")
}