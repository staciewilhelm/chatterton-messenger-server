package domain

import (
	"net/http"
	"strconv"
)

type QueryParams struct {
	Limit       string
	RecipientID string
	SenderID    string
}

func GetQueryParams(r *http.Request) *QueryParams {
	var params QueryParams
	query := r.URL.Query()

	if query.Has("limit") {
		queryLimit := query.Get("limit")
		intLimit, err := strconv.ParseInt(queryLimit, 6, 12)
		if err == nil && intLimit <= 100 {
			params.Limit = query.Get("limit")
		}
	}

	if query.Has("recipient_id") && query.Has("sender_id") {
		params.RecipientID = query.Get("recipient_id")
		params.SenderID = query.Get("sender_id")
	}

	return &params
}
