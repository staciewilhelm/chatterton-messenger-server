package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"chatterton-messenger-server/application"
	"chatterton-messenger-server/domain"
	"chatterton-messenger-server/middleware"
	"chatterton-messenger-server/models"
)

type ApplicationHandlers struct {
	*application.MessageApplication
}

func InitializeHTTP() *mux.Router {
	app := ApplicationHandlers{}
	return app.router()
}

func (app *ApplicationHandlers) createAPIRoutes(r *mux.Router) *mux.Router {
	routers := r.PathPrefix("/api").Subrouter()
	routers.Handle("/", middleware.JSONHandlerFunc(app.handleBaseGet)).Methods(http.MethodGet)

	messageRouter := routers.PathPrefix("/messages").Subrouter()
	messageRouter.Handle("", middleware.JSONHandlerFunc(app.handleMessagesPost)).Methods(http.MethodPost)
	messageRouter.Handle("", middleware.JSONHandlerFunc(app.handleMessagesGet)).Methods(http.MethodGet)

	return routers
}

func (app *ApplicationHandlers) handleBaseGet(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	fmt.Println("Welcome to Chatterton!")
	return nil, nil
}

func (app *ApplicationHandlers) handleMessagesGet(w http.ResponseWriter, req *http.Request) (interface{}, error) {
	params := domain.GetQueryParams(req)

	resp, err := app.GetMessages(params)
	if err != nil {
		log.Println("Error returning messages from app", err)
		appErr := fmt.Sprintf("Error returning messages from app: %s", err)
		return nil, errors.New(appErr)
	}

	return resp, nil
}

func (app *ApplicationHandlers) handleMessagesPost(w http.ResponseWriter, req *http.Request) (interface{}, error) {
	message := &models.Message{}
	err := json.NewDecoder(req.Body).Decode(message)
	if err != nil {
		log.Println("Error parsing request body data", err)
		appErr := fmt.Sprintf("Error parsing request body data: %s", err)
		return nil, errors.New(appErr)
	}

	resp, err := app.CreateMessage(message)
	if err != nil {
		log.Println("Error creating message", err)
		appErr := fmt.Sprintf("Error creating message: %s", err)
		return nil, errors.New(appErr)
	}

	return resp, nil
}

func (app *ApplicationHandlers) router() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	routers := app.createAPIRoutes(r)

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
