package config

import (
	"fmt"
	"time"

	"github.com/faridlan/auth-go/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase() *gorm.DB {

	// host := os.Getenv("HOST")
	// if host == "" {
	// 	host = "localhost"
	// }
	// dbname := os.Getenv("DB_NAME")
	// user := os.Getenv("USER")
	// password := os.Getenv("PASSWORD")
	// port := os.Getenv("PORT")

	config, err := helper.GetEnv()
	if err != nil {
		panic(err)
	}

	host := config.GetString("HOST")
	if host == "" {
		host = "localhost"
	}

	dbname := config.GetString("DB_NAME")
	user := config.GetString("USER")
	password := config.GetString("PASSWORD")
	port := config.GetString("PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	return db

}
