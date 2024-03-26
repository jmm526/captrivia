package app

import (
	"fmt"
	"log"
	"os"
	
	"github.com/jmm526/captrivia/internal/api"
)

func Run() {
	fmt.Println("Running app")
	// Setup the server
	router, err := api.SetupServer()
	if err != nil {
		log.Fatalf("Server setup failed: %v", err)
	}

	// set port to PORT or 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Println("Server starting on port " + port)
	log.Fatal(router.Run(":" + port))
}
