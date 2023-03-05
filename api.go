package chat_gpt_ppt

import (
	"fmt"
	"log"
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
	// init client
	c := GetClient(config.ClientType)
	if c == nil {
		return fmt.Errorf("no client named: %v", config.ClientType)
	}
	c.SetToken(config.Token)
	// init renderer
	renderer := GetRenderer(config.RendererType)
	if renderer == nil {
		return fmt.Errorf("no renderer named: %v", config.RendererType)
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
			return err
		}
		topics = append(topics, resp)
	}

	// renderer
	logger.Println("start rendering")
	for _, eachTopic := range topics {
		renderer.AddTopic(eachTopic)
	}
	err := renderer.RenderFile(config.OutputFile)
	if err != nil {
		return err
	}
	logger.Println("everything done, output saved to " + config.OutputFile)
	return nil
}
