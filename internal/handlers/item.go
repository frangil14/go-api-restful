package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/frangil14/go-api-restful/internal/db"
	"github.com/frangil14/go-api-restful/internal/models"
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

func CreateUser(w http.ResponseWriter, r *http.Request) {
    // Decodifica el cuerpo de la solicitud
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Ejecuta la consulta de inserción
    query := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id`
    err := db.DB.QueryRow(query, user.Name, user.Age).Scan(&user.ID)
    if err != nil {
        http.Error(w, "Error inserting user into database", http.StatusInternalServerError)
        return
    }

    // Establece el encabezado de respuesta y codifica el usuario en JSON
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func HelloAPI(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "¡Hola, API!")
}