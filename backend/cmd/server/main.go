package main

import (
	"log"
	"os"

	httpadapter "regtech-backend/internal/adapters/http"
	_ "regtech-backend/internal/docs"
)

// @title           Regulatory Deadline Calendar API
// @version         1.0
// @description     Simple in-memory API to track regulatory deadlines for startups in Australia.
// @BasePath        /api/v1
func main() {
	handler := httpadapter.WireDeadlineHandler()
	router := httpadapter.NewRouter(handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on :%s ...", port)
	if err := router.Run("0.0.0.0:" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
