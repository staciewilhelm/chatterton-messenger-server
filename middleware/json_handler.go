package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"chatterton-messenger-server/domain"
)

type (
	JSONHandlerFunc func(http.ResponseWriter, *http.Request) (interface{}, error)
)

func (fn JSONHandlerFunc) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var (
		api       domain.APIResponse
		errorList []string
	)

	if err := req.ParseForm(); err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	resp, err := fn(w, req)

	if err != nil {
		errorList = append(errorList, err.Error())
		api.Response.Code = strconv.Itoa(http.StatusBadRequest)
		api.Response.Message = string(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		api.Data = resp
		api.Response.Code = strconv.Itoa(http.StatusOK)
		api.Response.Message = "Success"
		w.WriteHeader(http.StatusOK)
	}

	if err := json.NewEncoder(w).Encode(&api); err != nil {
		log.Println(err)
		return
	}
}
