package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type SafeUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func respondJson(w http.ResponseWriter, data interface{}, status int) {

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)

}
func handleSignup(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only post method allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)

	if user.Username == "" {
		http.Error(w, "username required", http.StatusBadRequest)
		return
	}
	if user.Password == "" {
		http.Error(w, "password required", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO users(username,password) VALUES ($1,$2) RETURNING id`
	err := db.QueryRow(query, user.Username, user.Password).Scan(&user.ID)
	if err != nil {
		http.Error(w, "Username already taken or a DB error", http.StatusConflict)
		return
	}

	// this is for  prevent sending password to frontend , topil oru SafeUser struct ndd ath undakiyatah user structin  password matram ozhivakittan
	safeuser := SafeUser{
		ID:       user.ID,
		Username: user.Username,
	}
	respondJson(w, safeuser, http.StatusCreated)

}

func handleLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "only post method allowed ", http.StatusMethodNotAllowed)
		return
	}
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid json ", http.StatusBadRequest)
		return
	}

	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)

	if user.Username == "" || user.Password == "" {
		http.Error(w, "username and password cant be blank", http.StatusBadRequest)
		return
	}
	query := `SELECT id FROM users WHERE username=$1 AND password=$2`
	err := db.QueryRow(query, user.Username, user.Password).Scan(&user.ID)
	if err == sql.ErrNoRows {
		http.Error(w, "Invaild username OR Password", http.StatusUnauthorized)
		return

	} else if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	safeuser := SafeUser{
		ID:       user.ID,
		Username: user.Username,
	}
	respondJson(w, safeuser, http.StatusOK)

}
