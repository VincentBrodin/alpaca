package gollama

import "api/graph/models"

type Output struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int      `json:"created"`
	Model             string   `json:"model"`
	ServiceTier       string   `json:"service_tier"`
	SystemFingerprint string   `json:"system_fingerprint"`
	Choices           []Choise `json:"choices"`
}

type Choise struct {
	Index        int     `json:"index"`
	Delta        Delta   `json:"delta"`
	Logprobs     *string `json:"logprobs"`
	FinishReason *string `json:"finish_reason"`
}

type Delta struct {
	Content string `json:"content"`
}

func (o *Output) ToResponse() *models.Response {
	response := new(models.Response)
	if len(o.Choices) == 0 {
		response.Content = ""
	} else if o.Choices[0].FinishReason != nil {
		response.Content = ""
		response.Done = true
	} else {
		response.Content = o.Choices[0].Delta.Content
	}

	return response
}
