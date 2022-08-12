package config

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"

	"github.com/staciewilhelm/chatterton-messenger-server/handlers"
	"github.com/staciewilhelm/chatterton-messenger-server/models"
)

func StartServer() {
	fmt.Println("Starting Chatterton server...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := models.InitializeDB(); err != nil {
		log.Fatalln("Error initializing database", err)
	}

	r := handlers.InitializeHTTP()

	var port = os.Getenv("PORT")
	if port == "" {
		port = ":80"
	}

	log.Println(fmt.Sprintf("Chatterton is up and running at: http://127.0.0.1%v", port))
	log.Fatalln(http.ListenAndServe(port, r))
}
