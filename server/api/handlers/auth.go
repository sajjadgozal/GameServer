package handlers

import (
	"net/http"

	"sajjadgozal/gameserver/internal/services/auth"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, authService *auth.AuthService) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()

	email := r.FormValue("email")

	password := r.FormValue("password")

	account, err := authService.GetAccountByEmail(email)
	if err != nil {
		http.Error(w, "Failed to get account", http.StatusInternalServerError)
		return
	}

	if account.ID == 0 {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	err = authService.VerifyPassword(account.Password, password)
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	jwt, err := authService.GenerateJWT(account)
	if err != nil {
		http.Error(w, "Failed to generate JWT", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(`{"message": "Login handler" , "jwt": "` + jwt + `"}`))
}

func RegisterHandler(w http.ResponseWriter, r *http.Request, authService *auth.AuthService) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	account := authService.CreateAccount(name, email, password)

	if account.ID == 0 {
		http.Error(w, "Failed to create account", http.StatusInternalServerError)
		return
	}

	jwt, err := authService.GenerateJWT(account)
	if err != nil {
		http.Error(w, "Failed to generate JWT", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(`{"message": "Register handler" , "jwt": "` + jwt + `"}`))
}
