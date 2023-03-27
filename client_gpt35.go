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
Here is a prompt:

I would like to prepare a speech in pure markdown format, 
with language and topic consistent with the outline provided. 
You do not need to add any additional replies or explanations.

Here is the global outline:

%s

Then, I will begin to provide you with titles, 
and you will generate content for each page accordingly. 
Each title corresponds to one page of content, with a maximum of 100 words per page. 
Conciseness is key.
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
