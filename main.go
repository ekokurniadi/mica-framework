package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	env := godotenv.Load()
	if env != nil {
		log.Fatal(env.Error())
	}

	user := os.Getenv("USER_HOST")
	password := os.Getenv("USER_PASS")
	host := os.Getenv("HOST")
	database := os.Getenv("DATABASE_NAME")
	databasePort := os.Getenv("DATABASE_PORT")

	dsn := "" + user + ":" + password + "@tcp(" + host + ":" + databasePort + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	router := gin.Default()
	router.Run()
}
