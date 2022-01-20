package main

import (
	"github.com/haidityara/mtix/server"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Fatalln("Failed to start server:", err)
	}
}
