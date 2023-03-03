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
	// cmd parse
	tokenFile := flag.String("token", "./token.txt", "token file path")
	topicFile := flag.String("topic", "./topic.txt", "topic file path")
	outputFile := flag.String("output", "./output.html", "out path")
	rendererType := flag.String("renderer", cgp.RendererRemark, "renderer type")
	rendererBin := flag.String("rendererBin", "", "binary file for renderer")
	clientType := flag.String("client", cgp.ClientGpt35, "gpt client type")
	flag.Parse()

	logger := log.Default()

	// prepare
	tokenBytes, err := os.ReadFile(*tokenFile)
	panicIfErr(err)
	topicContents, err := os.ReadFile(*topicFile)
	panicIfErr(err)
	questions := strings.Split(string(topicContents), "\n")

	// init client
	c := cgp.GetClient(*clientType)
	if c == nil {
		panic(fmt.Errorf("no client named: %v", *clientType))
	}
	c.SetToken(string(tokenBytes))
	// init renderer
	renderer := cgp.GetRenderer(*rendererType)
	if renderer == nil {
		panic(fmt.Errorf("no renderer named: %v", *rendererType))
	}
	if *rendererBin != "" {
		logger.Printf("set renderer bin: %v\n", *rendererBin)
		renderer.SetBinPath(*rendererBin)
	}

	// fill topics
	topics := make([]*cgp.Topic, 0)
	for _, eachTopic := range questions {
		resp, err := c.FillTopic(eachTopic)
		panicIfErr(err)
		topics = append(topics, resp)
	}

	// renderer
	logger.Println("start rendering")
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
