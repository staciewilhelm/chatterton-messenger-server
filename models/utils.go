package models

import (
	"fmt"

	"github.com/staciewilhelm/chatterton-messenger-server/domain"
)

func GetMessagesSQLWithQueryParams(sql string, params *domain.QueryParams) string {
	const (
		whereSQL   = " created_at > now() - interval '30 day' AND message_type = 'message'"
		orderBySQL = " ORDER BY created_at DESC"
	)
	var sqlWithParams = fmt.Sprint(sql)

	if params.RecipientID != "" && params.SenderID != "" {
		sqlWithParams += fmt.Sprintf(" WHERE recipient_id = %s AND sender_id = %s", params.RecipientID, params.SenderID)
		sqlWithParams += " AND" + whereSQL + orderBySQL
	} else {
		sqlWithParams += " WHERE " + whereSQL + orderBySQL
	}

	sqlLimit := "100"
	if params.Limit != "" {
		sqlLimit = params.Limit
	}
	sqlWithParams += fmt.Sprintf(" LIMIT %s", sqlLimit)

	sqlWithParams += ";"
	return sqlWithParams
}
