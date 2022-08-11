package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"

	"chatterton-messenger-server/handlers"
)

type Server struct {
	DB     *sql.DB
	Router *mux.Router
}

func (s *Server) initDB() error {
	var (
		dbPort   = os.Getenv("DB_PORT")
		dbName   = os.Getenv("DB_DATABASE")
		host     = os.Getenv("DB_HOST")
		username = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
	)

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, dbPort, username, password, dbName)

	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		log.Println("Error pinging database", err)
		return err
	}

	fmt.Println("Database successfully connected!")

	s.DB = conn
	return nil
}

func (s *Server) StartServer() {
	fmt.Println("Starting Chatterton server...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := s.initDB(); err != nil {
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
