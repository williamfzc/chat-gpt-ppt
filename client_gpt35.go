package chat_gpt_ppt

import (
	"context"
	"fmt"
	"strings"

	gogpt "github.com/sashabaranov/go-gpt3"
)

type ChatGPTClient struct {
	token  string
	client *gogpt.Client
	topics []string
}

func (c *ChatGPTClient) Prepare(topics []string) error {
	c.client = gogpt.NewClient(c.token)

	topicsStr := strings.Join(topics, "\n")

	// prompt
	_, err := c.client.CreateChatCompletion(context.Background(), gogpt.ChatCompletionRequest{
		Model: gogpt.GPT3Dot5Turbo,
		Messages: []gogpt.ChatCompletionMessage{
			{
				Role: `system`,
				Content: fmt.Sprintf(`
你现在是一个PPT生成工具。
我接下来会给你发一个主题列表，请你根据我的主题列表，为我输出若干页PPT内容；
ppt内容是纯markdown，语言与主题列表一致，你不需要额外添加任何回复与解释。

%s

在我下次回复时开始生成。
`, topicsStr),
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
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
				Role:    `system`,
				Content: "ok, go on",
			},
		},
	})
	if err != nil {
		return nil, err
	}

	logger.Printf("topic %s done \n", topic)
	return &Topic{
		Title:   topic,
		Content: resp.Choices[0].Message.Content,
	}, nil
}

func (c *ChatGPTClient) SetToken(token string) {
	c.token = token
}
