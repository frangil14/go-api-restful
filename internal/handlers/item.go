package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
    idStr := vars["id"]

    // Validate that the id is a number
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid user ID format", http.StatusBadRequest)
        return
    }

    var user models.User

    query := `SELECT id, name, age FROM users WHERE id = $1`
    err = db.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Age)
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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"]

    // Validate that the id is a number
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid user ID format", http.StatusBadRequest)
        return
    }

    var user models.User
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields() 

    if err := decoder.Decode(&user); err != nil {
        http.Error(w, fmt.Sprintf("Invalid request payload or unknown fields: %v", err), http.StatusBadRequest)
        return
    }

    if user.Name == nil || user.Age == nil {
        http.Error(w, "Name and age are required fields", http.StatusBadRequest)
        return
    }

    query := `UPDATE users SET name = $1, age = $2 WHERE id = $3 RETURNING id`
    err = db.DB.QueryRow(query, user.Name, user.Age, id).Scan(&user.ID)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "User not found", http.StatusNotFound)
            return
        }
        http.Error(w, "Error updating user in database", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func UpdateUserPartially(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"]

    // Validate that the id is a number
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid user ID format", http.StatusBadRequest)
        return
    }

    var user models.User
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields() 

    if err := decoder.Decode(&user); err != nil {
        http.Error(w, fmt.Sprintf("Invalid request payload or unknown fields: %v", err), http.StatusBadRequest)
        return
    }


    var updateFields []string
    var updateValues []interface{}

    if user.Name != nil {
        updateFields = append(updateFields, "name")
        updateValues = append(updateValues, user.Name)
    }

    if user.Age != nil {
        updateFields = append(updateFields, "age")
        updateValues = append(updateValues, user.Age)
    }

    // Check if there are valid fields to update
    if len(updateFields) == 0 {
        http.Error(w, "No fields to update", http.StatusBadRequest)
        return
    }

    // Dynamically build the SQL query
    setClause := ""
    for i, field := range updateFields {
        if i > 0 {
            setClause += ", "
        }
        setClause += field + "=$" + fmt.Sprintf("%d", i+1)
    }
    query := fmt.Sprintf("UPDATE users SET %s WHERE id=%d RETURNING id", setClause, id)
    

    err = db.DB.QueryRow(query, updateValues...).Scan(&user.ID)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "User not found", http.StatusNotFound)
            return
        }
        http.Error(w, "Error updating user in database", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}



func HelloAPI(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "¡Hola, API!")
}