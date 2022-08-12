package application

import (
	"github.com/staciewilhelm/chatterton-messenger-server/domain"
	"github.com/staciewilhelm/chatterton-messenger-server/models"
)

func (app *MessageApplication) CreateMessage(message *models.Message) (string, error) {
	messageID, err := models.Add(message)
	if err != nil {
		return "", err
	}

	return messageID, nil
}

func (app *MessageApplication) GetMessages(params *domain.QueryParams) ([]models.Message, error) {
	messages, err := models.FindAll(params)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

type MessageApplication struct{}
