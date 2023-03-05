package chat_gpt_ppt

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

type ApiConfig struct {
	Token        string       `json:"token"`
	Topics       []string     `json:"topics"`
	OutputFile   string       `json:"outputFile"`
	RendererType RendererType `json:"rendererType"`
	RendererBin  string       `json:"rendererBin"`
	ClientType   ClientType   `json:"clientType"`
}

var logger = log.Default()

func GenAndRender(config ApiConfig) error {
	content, err := GenAndRenderString(config)
	if err != nil {
		return err
	}
	err = os.WriteFile(config.OutputFile, []byte(content), fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func GenAndRenderString(config ApiConfig) (string, error) {
	// init client
	c := GetClient(config.ClientType)
	if c == nil {
		return "", fmt.Errorf("no client named: %v", config.ClientType)
	}
	c.SetToken(config.Token)
	// init renderer
	renderer := GetRenderer(config.RendererType)
	if renderer == nil {
		return "", fmt.Errorf("no renderer named: %v", config.RendererType)
	}
	if config.RendererBin != "" {
		logger.Printf("set renderer bin: %v\n", config.RendererBin)
		renderer.SetBinPath(config.RendererBin)
	}

	// fill topics
	topics := make([]*Topic, 0)
	for _, eachTopic := range config.Topics {
		resp, err := c.FillTopic(eachTopic)
		if err != nil {
			return "", err
		}
		topics = append(topics, resp)
	}

	// renderer
	logger.Println("start rendering")
	for _, eachTopic := range topics {
		renderer.AddTopic(eachTopic)
	}
	str, err := renderer.RenderString()
	if err != nil {
		return "", err
	}
	return str, nil
}
