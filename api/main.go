package main

import (
	"fmt"
	"go-screenshot/internal/handlers"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HandleRoot)

	fmt.Println("server starting on port ", port)
	http.ListenAndServe(":"+port, mux)
}
