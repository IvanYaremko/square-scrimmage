package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	mux.HandleFunc("/", serveHome)

	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	server.ListenAndServe()
}

func serveHome(w http.ResponseWriter, r *http.Request) {

}
