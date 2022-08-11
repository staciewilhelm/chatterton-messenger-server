package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID          uuid.UUID `json:"id"`
	SenderID    string    `json:"sender_id"`
	RecipientID string    `json:"recipient_id"`
	Text        string    `json:"message_text"`
	Type        string    `json:"message_type"`
	CreatedAt   time.Time `json:"created_at"`
}
