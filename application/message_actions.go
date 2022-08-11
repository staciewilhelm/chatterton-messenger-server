package application

import (
	"time"

	"github.com/google/uuid"

	"chatterton-messenger-server/domain"
)

type MessageApplication struct{}

func (app *MessageApplication) GetMessages() ([]domain.Message, error) {
	return []domain.Message{
		{
			ID:          uuid.New(),
			RecipientID: "recipient-id",
			SenderID:    "sender-id",
			Text:        "First message",
			CreatedAt:   time.Now(),
		},
		{
			ID:          uuid.New(),
			RecipientID: "recipient-id",
			SenderID:    "sender-id",
			Text:        "Second message",
			CreatedAt:   time.Now(),
		},
		{
			ID:          uuid.New(),
			RecipientID: "sender-id",
			SenderID:    "recipient-id",
			Text:        "A reply",
			CreatedAt:   time.Now(),
		},
		{
			ID:          uuid.New(),
			RecipientID: "recipient-two-id",
			SenderID:    "sender-id",
			Text:        "A message to someone else.",
			CreatedAt:   time.Now(),
		},
	}, nil
}
