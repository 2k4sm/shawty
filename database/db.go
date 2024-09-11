package database

import (
	"context"
	"fmt"
	"os"

	"github.com/2k4sm/shawty/models"
	"github.com/gofiber/fiber/v2/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Ctx = context.Background()

type DBConfig struct {
	DB_HOST   string
	DB_USER   string
	DB_PG_PWD string
	DB_NAME   string
	DB_PORT   string
	SSLMODE   string
}

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PWD"),
		DB:       dbNo,
	})

	return rdb
}

func InitPGdb(conf *DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		conf.DB_HOST, conf.DB_USER, conf.DB_PG_PWD, conf.DB_NAME, conf.DB_PORT, conf.SSLMODE)

	var db *gorm.DB
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL.")
	}

	if err := autoMigrate(db); err != nil {
		log.Fatalf("Error automigrating database: %s", err)
	}

	return db
}

func autoMigrate(database *gorm.DB) error {
	return database.AutoMigrate(&models.User{}, &models.Store{})
}
