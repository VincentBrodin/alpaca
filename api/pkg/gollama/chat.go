package gollama

import (
	"api/graph/models"
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
)

type Handler func(*models.Response) error

func Send(url string, chat *models.Chat, handler Handler) (*models.Message, error) {
	body, err := json.Marshal(chat)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(url+"/api/chat",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := new(models.Response)
	message := new(models.Message)

	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		line := scanner.Bytes()
		if err := json.Unmarshal(line, response); err != nil {
			return nil, err
		}
		if err := handler(response); err != nil {
			return nil, err
		}
		message.Role = response.Message.Role
		message.Content += response.Message.Content
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return message, nil
}
