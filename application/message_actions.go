package application

import (
	"chatterton-messenger-server/domain"
	"chatterton-messenger-server/models"
)

func (app *MessageApplication) GetMessages(params *domain.QueryParams) ([]models.Message, error) {
	messages, err := models.FindAll(params)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

type MessageApplication struct{}
