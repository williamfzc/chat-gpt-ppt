package main

import (
	"flag"
	"os"
	"strings"

	cgp "github.com/williamfzc/chat-gpt-ppt"
	_ "github.com/williamfzc/chat-gpt-ppt/assets"
)

func main() {
	tokenFile := flag.String("token", "./token.txt", "token file path")
	topicFile := flag.String("topic", "./topic.txt", "topic file path")
	outputFile := flag.String("output", "./output.ppt", "out path")
	flag.Parse()

	tokenBytes, err := os.ReadFile(*tokenFile)
	panicIfErr(err)
	topicContents, err := os.ReadFile(*topicFile)
	panicIfErr(err)
	questions := strings.Split(string(topicContents), "\n")

	topics := make([]*cgp.Topic, 0)
	c := cgp.NewClient(string(tokenBytes))
	for _, eachTopic := range questions {
		resp, err := c.AskTopic(eachTopic)
		panicIfErr(err)
		topics = append(topics, &cgp.Topic{
			Title:   eachTopic,
			Content: resp.Choices[0].Message.Content,
		})
	}

	// renderer
	renderer := cgp.NewMarpRenderer()
	for _, eachTopic := range topics {
		renderer.AddTopic(eachTopic)
	}
	err = renderer.RenderFile(*outputFile)
	panicIfErr(err)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
