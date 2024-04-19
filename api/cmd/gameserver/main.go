package main

import (
	"log"
	"net/http"

	"sajjadgozal/gameserver/internal/database"
	"sajjadgozal/gameserver/internal/handlers"
	"sajjadgozal/gameserver/internal/services/auth"

	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	db, err := database.NewDb(database.Options{
		URI:               "localhost",
		Database:          "project_test",
		UserName:          "root",
		Password:          "root183729",
		Port:              "1234",
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

	// if err == nil {
	// 	log.Println("Connected to database.")
	// 	break // Connection successful
	// }

	// 	log.Printf("Failed to connect to database: %v. Retrying...", err)
	// 	time.Sleep(delay)
	// 	retries--
	// }
	// if retries == 0 {
	// 	log.Fatalf("Failed to connect to database after 5 retries : %s", err)
	// }

	// server := NewApiServer(db, ":3000")

	// err = server.Start()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, authService)
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterHandler(w, r, authService)
	})

	http.HandleFunc("/health", handlers.HealthHandler)

	print("Server started successfully")

	_ = http.ListenAndServe(":3000", nil)

}
