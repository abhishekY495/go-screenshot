package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)

	fmt.Println("Server starting on port ", port)
	http.ListenAndServe(":"+port, mux)
}
