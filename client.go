package chat_gpt_ppt

import (
	"context"

	gogpt "github.com/sashabaranov/go-gpt3"
)

type Client struct {
	token string
}

func NewClient(token string) *Client {
	return &Client{
		token,
	}
}

func (c *Client) AskTopic(topic string) (*gogpt.CompletionResponse, error) {
	gptClient := gogpt.NewClient(c.token)
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3Dot5Turbo,
		MaxTokens: 5,
		Prompt:    topic,
	}
	resp, err := gptClient.CreateCompletion(ctx, req)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
