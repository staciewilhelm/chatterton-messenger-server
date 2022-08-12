package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"chatterton-messenger-server/domain"
)

type Message struct {
	ID          uuid.UUID `json:"id"`
	SenderID    string    `json:"sender_id"`
	RecipientID string    `json:"recipient_id"`
	Text        string    `json:"message_text"`
	Type        string    `json:"message_type"`
	CreatedAt   time.Time `json:"created_at"`
}

func FindAll(params *domain.QueryParams) ([]Message, error) {
	initialSQL := "SELECT * FROM messages"
	sql := GetMessagesSQLWithQueryParams(initialSQL, params)

	data, err := DB.Query(sql)
	if err != nil {
		queryErr := fmt.Sprintf("error running query: %s", err)
		return nil, errors.New(queryErr)
	}

	var messages []Message
	for data.Next() {
		var oneMessage Message
		err = data.Scan(&oneMessage.ID, &oneMessage.SenderID, &oneMessage.RecipientID, &oneMessage.Text, &oneMessage.Type, &oneMessage.CreatedAt)
		if err != nil {
			scanErr := fmt.Sprintf("error scanning data: %s", err)
			return nil, errors.New(scanErr)
		}
		messages = append(messages, oneMessage)
	}

	return messages, nil
}
