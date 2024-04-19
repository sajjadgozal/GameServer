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

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(`{"message": "Login handler"}`))

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
