package main

import (
	"log"
	"net/http"

	"sajjadgozal/gameserver/api/routes"
	"sajjadgozal/gameserver/internal/database"
	"sajjadgozal/gameserver/internal/services/auth"
	"sajjadgozal/gameserver/internal/services/wallet"
)

func main() {

	db, err := database.NewDb(database.Options{
		URI: "postgresdb_test",
		// URI:      "localhost",
		Database: "project_test",
		UserName: "root",
		Password: "root183729",
		Port:     "5432",
		// Port:              "1234",
		SSLMode:           "disable",
		ConnectTimeout:    5,
		PingTimeout:       5,
		ReadTimeout:       5,
		WriteTimeout:      5,
		DisconnectTimeout: 5,
	})
	if err != nil {
		log.Fatal(err)
	}

	// new auth service using the db
	authService := auth.NewAuthService(db)

	// wallet
	walletService := wallet.NewWalletService()

	// Register routes - pass the authService
	router := routes.NewRouter(authService)

	router.RegisterAuthRoutes()
	router.RegisterWalletRoutes(walletService)

	// Start the server
	log.Println("Server started successfully")
	log.Fatal(http.ListenAndServe(":3000", router))

}
