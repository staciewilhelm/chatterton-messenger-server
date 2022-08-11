package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID          uuid.UUID `json:"id"`
	RecipientID string    `json:"recipient_id"`
	SenderID    string    `json:"sender_id"`
	Text        string    `json:"text"`
	CreatedAt   time.Time `json:"created_at"`
}
