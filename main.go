package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")

	router := http.NewServeMux()

	router.HandleFunc("/", ServeHome)
	fmt.Println("Welcome to the backend jungle")
	fmt.Printf("Secret port: %v\n", port)
	log.Fatal(http.ListenAndServe(":6969", router))
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./src/homepage_tml.html")
}
