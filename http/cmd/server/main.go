package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pradeepbgs/internal/config"
	sqlc "github.com/pradeepbgs/internal/db"
	"github.com/pradeepbgs/internal/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	fmt.Println("port==", os.Getenv("PORT"))
	conf := config.Load()

	app := http.NewServeMux()
	dbPool := config.NewPostgres(conf.DB_URL)
	queries := sqlc.New(dbPool)
	// setup router
	router.SetupRouter(app,queries)

	// run the main server
	log.Println("Server running on port", conf.Port)
	log.Fatal(http.ListenAndServe(conf.Port, app))
}
