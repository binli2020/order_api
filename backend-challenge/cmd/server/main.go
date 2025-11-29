package main

import (
	"log"
	"net/http"

	"github.com/binli2020/order_api/backend-challenge/internal/router"
)

func main() {
	r := router.NewRouter()
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
