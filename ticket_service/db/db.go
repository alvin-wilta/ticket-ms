package db

import (
	"fmt"
	"log"

	"github.com/alvin-wilta/ticket-ms/ticket_service/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(c *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v port=%v sslmode=disable", c.PostgresAddr, c.PostgresUser, c.PostgresPass, c.PostgresPort)
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

func InitCache(c *config.Config) *redis.Client {
	addr := fmt.Sprintf("%s:%s", c.RedisAddr, c.RedisPort)
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return rdb
}
