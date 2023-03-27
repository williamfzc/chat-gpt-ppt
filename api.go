package chat_gpt_ppt

import (
	"fmt"
	"log"
	"strings"

	"github.com/abiosoft/ishell/v2"
)

type ApiConfig struct {
	Token        string       `json:"token"`
	Topics       []string     `json:"topics"`
	OutputFile   string       `json:"outputFile"`
	RendererType RendererType `json:"rendererType"`
	RendererBin  string       `json:"rendererBin"`
	ClientType   ClientType   `json:"clientType"`
	Interactive  bool         `json:"interactive"`
}

var logger = log.Default()

func GenAndRenderString(shellContext *ishell.Context, config ApiConfig) (string, error) {
	// init client
	c := GetClient(config.ClientType)
	if c == nil {
		return "", fmt.Errorf("no client named: %v", config.ClientType)
	}

	// https://github.com/williamfzc/chat-gpt-ppt/issues/12
	safeToken := strings.TrimSpace(config.Token)
	c.SetToken(safeToken)

	// init renderer
	renderer := GetRenderer(config.RendererType)
	if renderer == nil {
		return "", fmt.Errorf("no renderer named: %v", config.RendererType)
	}
	if config.RendererBin != "" {
		shellContext.Printf("set renderer bin: %v\n", config.RendererBin)
		renderer.SetBinPath(config.RendererBin)
	}

	// prepare
	shellContext.Println("start preparing ...")
	err := c.Prepare(config.Topics)
	if err != nil {
		return "", err
	}

	// fill
	shellContext.Println("start generating ...")
	topics := make([]*Topic, 0)
	for _, eachTopic := range config.Topics {
		finalTopic, err := getFinalTopic(shellContext, c, eachTopic, config.Interactive)
		if err != nil {
			return "", err
		}
		topics = append(topics, finalTopic)
	}

	// renderer
	shellContext.Println("start rendering ...")
	shellContext.Stop()
	for _, eachTopic := range topics {
		renderer.AddTopic(eachTopic)
	}
	str, err := renderer.RenderString()
	if err != nil {
		return "", err
	}
	return str, nil
}

func getFinalTopic(shellContext *ishell.Context, c Client, eachTopic string, needConfirm bool) (*Topic, error) {
	resp, err := c.FillTopic(eachTopic)
	if err != nil {
		return nil, err
	}

	if !needConfirm {
		return resp, nil
	}

	shellContext.Println("Here is your response, type any key to continue, type 'n' to edit", resp.ToMarkdown())
	ok := shellContext.ReadLine()
	if ok != "n" {
		return resp, nil
	} else {
		shellContext.Println("You can enter a new topic to regenerate this page.")
		newTopic := shellContext.ReadLine()
		return getFinalTopic(shellContext, c, newTopic, needConfirm)
	}
}
