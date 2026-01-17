package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pradeepbgs/internals/config"
	"github.com/pradeepbgs/internals/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	fmt.Println("port==", os.Getenv("PORT"))
	config := config.Load()

	app := http.NewServeMux()

	// setup router
	router.SetupRouter(app)

	// run the main server
	log.Println("Server running on port", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, app))
}
