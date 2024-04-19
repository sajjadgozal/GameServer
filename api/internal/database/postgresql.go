package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Options struct {
	URI               string
	Database          string
	UserName          string
	Password          string
	Port              string
	SSLMode           string
	ConnectTimeout    int
	PingTimeout       time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	DisconnectTimeout time.Duration
}

func NewDb(opts Options) (*gorm.DB, error) {

	retries := 5
	delay := time.Second * 5
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%v", opts.URI, opts.Port, opts.UserName, opts.Password, opts.Database, opts.SSLMode, opts.ConnectTimeout)

	var db *gorm.DB
	var err error

	for retries > 0 {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		// Check if the connection is successful
		if err == nil {
			log.Println("Connected to database.")
			break // Connection successful
		}
		log.Printf("Failed to connect to database: %v. Retrying...", err)
		time.Sleep(delay)
		retries--
	}
	if retries == 0 {
		return nil, fmt.Errorf("failed to connect to database after 5 retries : %s", err)
	}
	// Automatically create the necessary table structure
	// err = db.AutoMigrate(&Account{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return db, nil
}
