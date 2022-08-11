package domain

import (
	"net/http"
)

type QueryParams struct {
	Limit string
}

func GetQueryParams(r *http.Request) *QueryParams {
	var params QueryParams
	query := r.URL.Query()

	if query.Has("limit") {
		params.Limit = query.Get("limit")
	}

	return &params
}
