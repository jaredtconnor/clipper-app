package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jaredtconnor/clipper-app/config"
	"github.com/jaredtconnor/clipper-app/internals/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	var err error

	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Port failed to parse")
	}

	// Connection URL setup

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	fmt.Println(dsn)

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to connect to database")
	}
	// DB.AutoMigrate()
	fmt.Println("Connection Opened to Database")

	// Migration of database
	DB.AutoMigrate(&model.Clipping{})
	fmt.Println("Database Migrated")

}
