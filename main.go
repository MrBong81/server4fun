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

	fs := http.FileServer(http.Dir("src"))
	router.Handle("/src/", http.StripPrefix("/src/", fs))

	router.HandleFunc("/test", ServeHome)

	fmt.Println("Welcome to the backend jungle")
	fmt.Printf("Secret port: %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./src/templ/homepage_tml.html")
}
