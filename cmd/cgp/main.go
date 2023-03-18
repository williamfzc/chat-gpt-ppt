package main

import (
	"flag"
	"io/fs"
	"os"
	"strings"

	"github.com/abiosoft/ishell/v2"
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

	// prepare
	tokenBytes, err := os.ReadFile(*tokenFile)
	panicIfErr(err)
	topicContents, err := os.ReadFile(*topicFile)
	panicIfErr(err)
	questions := strings.Split(string(topicContents), "\n")

	config := cgp.ApiConfig{
		Token:        string(tokenBytes),
		Topics:       questions,
		OutputFile:   *outputFile,
		RendererType: *rendererType,
		RendererBin:  *rendererBin,
		ClientType:   *clientType,
	}

	shell := ishell.New()
	cmd := &ishell.Cmd{
		Name: "gen",
		Help: "gen",
		Func: func(c *ishell.Context) {
			content, err := cgp.GenAndRenderString(c, config)
			panicIfErr(err)
			err = os.WriteFile(config.OutputFile, []byte(content), fs.ModePerm)
			panicIfErr(err)
		},
	}
	shell.AddCmd(cmd)
	err = shell.Process("gen")
	panicIfErr(err)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
