package routes

import (
	"net/http"

	"sajjadgozal/gameserver/api/handlers"
	"sajjadgozal/gameserver/api/middleware"
	"sajjadgozal/gameserver/internal/services/auth"
	"sajjadgozal/gameserver/internal/services/wallet"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
	as *auth.AuthService
	m  *middleware.AuthMiddleware
}

func NewRouter(as *auth.AuthService) *Router {
	router := mux.NewRouter().StrictSlash(true)
	return &Router{router, as, middleware.NewAuthMiddleware(as)}
}

func (r *Router) RegisterAuthRoutes() {

	r.HandleFunc("/health", handlers.HealthHandler).Methods("GET")
	r.HandleFunc("/protected", r.m.Auth(handlers.ProtectedHandler)).Methods("GET")

	r.HandleFunc("/register", func(w http.ResponseWriter, req *http.Request) {
		handlers.RegisterHandler(w, req, r.as)
	}).Methods("POST")

	r.HandleFunc("/login", func(w http.ResponseWriter, req *http.Request) {
		handlers.LoginHandler(w, req, r.as)
	}).Methods("POST")

}

func (r *Router) RegisterWalletRoutes(walletService *wallet.WalletService) {
	// router := mux.NewRouter()
	// router.HandleFunc("/wallet", walletService.CreateWallet).Methods("POST")
	// router.HandleFunc("/wallet/{address}", walletService.GetBallance).Methods("GET")
	// return router
}
