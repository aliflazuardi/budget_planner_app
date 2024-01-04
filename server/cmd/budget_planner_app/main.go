package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aliflazuardi/budget_planner_app/server/internal/routes"
)

// main app for backend web server and api
func main() {
	router := routes.NewRouter()

	port := 8080
	addr := fmt.Sprintf(":%d", port)

	log.Printf("Server listening on port %d", port)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal(err)
	}
}
