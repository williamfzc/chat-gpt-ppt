package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	cgp "github.com/williamfzc/chat-gpt-ppt"
)

func main() {
	tokenFile := flag.String("token", "./token.txt", "token file path")
	topicFile := flag.String("topic", "./topic.txt", "topic file path")
	outputFile := flag.String("output", "./output.html", "out path")
	rendererType := flag.String("renderer", cgp.RendererRemark, "renderer type")
	rendererBin := flag.String("rendererBin", "", "binary file for renderer")
	flag.Parse()

	tokenBytes, err := os.ReadFile(*tokenFile)
	panicIfErr(err)
	topicContents, err := os.ReadFile(*topicFile)
	panicIfErr(err)
	questions := strings.Split(string(topicContents), "\n")

	logger := log.Default()

	topics := make([]*cgp.Topic, 0)
	c := cgp.NewClient(string(tokenBytes))
	for _, eachTopic := range questions {
		resp, err := c.AskTopic(eachTopic)
		panicIfErr(err)
		topics = append(topics, &cgp.Topic{
			Title:   eachTopic,
			Content: resp.Choices[0].Message.Content,
		})
		logger.Printf("query topic %v done\n", eachTopic)
	}

	// renderer
	logger.Println("start rendering")
	renderer := cgp.GetRenderer(*rendererType)
	if renderer == nil {
		panic(fmt.Errorf("no renderer named: %v", *rendererType))
	}
	if *rendererBin != "" {
		logger.Printf("set renderer bin: %v\n", *rendererBin)
		renderer.SetBinPath(*rendererBin)
	}
	for _, eachTopic := range topics {
		renderer.AddTopic(eachTopic)
	}
	err = renderer.RenderFile(*outputFile)
	panicIfErr(err)
	logger.Println("everything done, output saved to " + *outputFile)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
