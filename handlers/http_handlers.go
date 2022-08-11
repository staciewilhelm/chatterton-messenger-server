package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"chatterton-messenger-server/application"
	"chatterton-messenger-server/middleware"
)

type ApplicationHandlers struct {
	*application.MessageApplication
}

func InitializeHTTP() *mux.Router {
	app := ApplicationHandlers{}
	return app.router()
}

func (app *ApplicationHandlers) router() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	routers := r.PathPrefix("/api").Subrouter()
	routers.Handle("/", middleware.JSONHandlerFunc(app.handleBaseGet)).Methods(http.MethodGet)

	messageRouter := routers.PathPrefix("/messages").Subrouter()
	messageRouter.Handle("/", middleware.JSONHandlerFunc(app.handleMessagesGet)).Methods(http.MethodGet)

	var port = os.Getenv("PORT")
	if port == "" {
		port = ":80"
	}
	host := fmt.Sprintf("127.0.0.1%v", port)
	srv := &http.Server{
		Addr:         host,
		Handler:      routers,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println(fmt.Sprintf("Chatterton is up and running at: http://127.0.0.1%v", port))
	log.Fatal(srv.ListenAndServe())

	return r
}

func (app *ApplicationHandlers) handleBaseGet(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	fmt.Println("Welcome to Chatterton!")
	return nil, nil
}

func (app *ApplicationHandlers) handleMessagesGet(w http.ResponseWriter, req *http.Request) (interface{}, error) {
	resp, err := app.GetMessages()
	if err != nil {
		log.Println("Error returning messages from app", err)
	}

	return resp, nil
}
