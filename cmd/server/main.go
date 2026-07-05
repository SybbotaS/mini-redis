package main

import (
	"log"

	"github.com/SybbotaS/mini-redis/internal/server"
)

func main() {
	srv := server.New(":6379")

	log.Println("Mini Redis is starting...")

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
