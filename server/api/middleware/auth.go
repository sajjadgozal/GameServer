package middleware

import (
	"net/http"
	"sajjadgozal/gameserver/internal/services/auth"
	"strconv"
)

type AuthMiddleware struct {
	as *auth.AuthService
}

func NewAuthMiddleware(as *auth.AuthService) *AuthMiddleware {
	return &AuthMiddleware{as}
}

func (a AuthMiddleware) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !a.isAuthenticated(r) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func (a AuthMiddleware) isAuthenticated(r *http.Request) bool {

	token := r.Header.Get("Authorization")
	if token == "" {
		return false
	}

	token = token[7:]

	claims, err := a.as.VerifyJWT(token)
	if err != nil {
		return false
	}

	r.Header.Set("user_id", strconv.Itoa(claims.UserID))

	return true
}
