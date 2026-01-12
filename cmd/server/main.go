package main

import (
	"fmt"
	"os"

	"github.com/aykutterzi/intellilog/internal/server"
)

func main() {
	server := server.NewServer()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := server.Start(":" + port)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
