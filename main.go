package main

import (
	"fmt"
	"golang-core/database"
	"golang-core/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	InitEnv()
	database.Init()
	port := os.Getenv("PORT")
	handler := router.Init()
	_ = handler.Run(port)
	fmt.Println(port)
}

func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("ERROR: .env file cannot load")
	}
}
