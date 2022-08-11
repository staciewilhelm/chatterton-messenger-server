package models

import (
	"fmt"
	"strconv"

	"chatterton-messenger-server/domain"
)

func GetSQLWithQueryParams(sql string, params *domain.QueryParams) string {
	var sqlWithParams = sql
	sqlLimit := "100"
	if params.Limit != "" {
		intLimit, err := strconv.ParseInt(params.Limit, 6, 12)
		if err == nil && intLimit < 100 {
			sqlLimit = params.Limit
		}
	}
	sqlWithParams += fmt.Sprintf(" LIMIT %s", sqlLimit)
	sqlWithParams += ";"
	return sqlWithParams
}
