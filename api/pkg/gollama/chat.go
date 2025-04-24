package gollama

import (
	"api/graph/models"
	"context"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type Handler func(*models.Response) error

func Send(chat *models.Chat, handler Handler) (*models.Message, error) {

	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPEN_AI_API_KEY")),
	)

	messages := make([]openai.ChatCompletionMessageParamUnion, len(chat.Messages))

	for i := range messages {
		if chat.Messages[i].Role == "user" {
			messages[i] = openai.UserMessage(chat.Messages[i].Content)

		} else if chat.Messages[i].Role == "assistant" {
			messages[i] = openai.AssistantMessage(chat.Messages[i].Content)

		} else {
			messages[i] = openai.DeveloperMessage(chat.Messages[i].Content)
		}
	}

	stream := client.Chat.Completions.NewStreaming(context.TODO(), openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    chat.Model,
		Seed:     openai.Int(0),
	})

	acc := openai.ChatCompletionAccumulator{}

	for stream.Next() {
		response := new(models.Response)
		chunk := stream.Current()
		acc.AddChunk(chunk)

		if _, ok := acc.JustFinishedContent(); ok {
			response.Done = true
		}

		if _, ok := acc.JustFinishedToolCall(); ok {
			response.Done = true
		}

		if _, ok := acc.JustFinishedRefusal(); ok {
			response.Done = true
		}

		if len(chunk.Choices) > 0 {
			content := chunk.Choices[0].Delta.Content
			response.Content = content

			if err := handler(response); err != nil {
				return nil, err
			}
		}
	}

	if stream.Err() != nil {
		return nil, stream.Err()
	}

	message := new(models.Message)
	message.Role = "assistant"
	message.Content = acc.Choices[0].Message.Content
	return message, nil
}
