package chat_gpt_ppt

import (
	"context"

	gogpt "github.com/sashabaranov/go-gpt3"
)

type ChatGPTClient struct {
	token string
}

func NewGpt35Client() Client {
	return &ChatGPTClient{}
}

func (c *ChatGPTClient) FillTopic(topic string) (*Topic, error) {
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

	return &Topic{
		Title:   topic,
		Content: resp.Choices[0].Message.Content,
	}, nil
}

func (c *ChatGPTClient) SetToken(token string) {
	c.token = token
}
