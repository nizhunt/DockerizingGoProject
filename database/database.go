package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


var DB *gorm.DB

func ConnectDb(){
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),
)
fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil{
		log.Fatal("Faild to connect to the db\n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	// Running Migration for the models
	log.Println("running migration")
	db.AutoMigrate(&Fact{})

	DB = db
}