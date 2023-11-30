package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres port=54320 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("CREATE DATABASE tickets;")
	if db == nil {
		log.Fatalf("DB not initalized, %v", err)
	}
	db.AutoMigrate(&Ticket{})
	return db
}
