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
我想准备一场演讲；
演讲内容是纯markdown，语言与主题列表一致，你不需要额外添加任何回复与解释；

先提供提纲给你：

%s

之后我会开始传递标题给你，你按标题给我生成内容即可；
每个标题对应一页内容；
词数不超过100词；
简明扼要为主；
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
				Role: `system`,
				Content: fmt.Sprintf(`
generate one slide,
- within 100 words
- with markdown format, such as 'ordered list'
- without title
- about '%s'
`, topic),
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
