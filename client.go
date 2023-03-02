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

func (c *Client) AskTopic(topic string) (*gogpt.ChatCompletionResponse, error) {
	gptClient := gogpt.NewClient(c.token)
	ctx := context.Background()

	resp, err := gptClient.CreateChatCompletion(ctx, gogpt.ChatCompletionRequest{
		Model: gogpt.GPT3Dot5Turbo,
		Messages: []gogpt.ChatCompletionMessage{
			{
				Role:    "user",
				Content: topic,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
