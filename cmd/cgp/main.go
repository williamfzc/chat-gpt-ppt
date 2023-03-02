package main

import cgp "github.com/williamfzc/chat-gpt-ppt"
import _ "github.com/williamfzc/chat-gpt-ppt/assets"

func main() {
	token := ""
	questions := []string{
		"a",
		"b",
		"c",
	}

	topics := make([]*cgp.Topic, 0)
	c := cgp.NewClient(token)
	for _, eachTopic := range questions {
		resp, err := c.AskTopic(eachTopic)
		panicIfErr(err)
		topics = append(topics, &cgp.Topic{
			Title:   eachTopic,
			Content: resp.Choices[0].Text,
		})
	}

	// renderer
	renderer := cgp.NewMarpRenderer()
	for _, eachTopic := range topics {
		renderer.AddTopic(eachTopic)
	}
	err := renderer.RenderFile("./output.ppt")
	panicIfErr(err)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
