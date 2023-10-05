package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/k0msak007/go-todo-app/database"
	"github.com/k0msak007/go-todo-app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db_name = ""
var db_port = "3306"
var db_user = "root"
var db_password = ""
var db_host = "127.0.0.1"

func bootDatabase() {
	if dbNaemEnv := os.Getenv("DB_NAME"); dbNaemEnv != "" {
		db_name = dbNaemEnv
	}

	if dbPortEnv := os.Getenv("DB_PORT"); dbPortEnv != "" {
		db_port = dbPortEnv
	}

	if dbUserEnv := os.Getenv("DB_USER"); dbUserEnv != "" {
		db_user = dbUserEnv
	}

	if dbPasswordEnv := os.Getenv("DB_PASSWORD"); dbPasswordEnv != "" {
		db_password = dbPasswordEnv
	}

	if dbHostEnv := os.Getenv("DB_HOST"); dbHostEnv != "" {
		db_host = dbHostEnv
	}
}

func connectDatabase() {
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db_user,
		db_password,
		db_host,
		db_port,
		db_name,
	)
	database.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Can't connect to database")
	} else {
		log.Println("Connected to database")
	}
}

func runMigration() {
	err := database.DB.AutoMigrate(&models.Todo{})
	if err != nil {
		fmt.Println(err)
		log.Println("Failed to migrate schema")
	} else {
		log.Println("schema migrated")
	}
}
