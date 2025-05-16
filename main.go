package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")

	fmt.Println("Welcome to the backend jungle")
	fmt.Printf("Secret port: %v\n", port)
}
