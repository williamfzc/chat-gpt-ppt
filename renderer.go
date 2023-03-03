package chat_gpt_ppt

import (
	"fmt"
)

type Topic struct {
	Title   string
	Content string
}

func (t *Topic) ToMarkdown() string {
	return fmt.Sprintf(`
# %s

%s
`, t.Title, t.Content)
}

type RendererType = string

const (
	RendererMarp   = "MARP"
	RendererRemark = "REMARK"
)

type Renderer interface {
	AddTopic(*Topic)
	RenderFile(outputPath string) error
	SetBinPath(path string)
}

func GetRenderer(rendererType RendererType) Renderer {
	switch rendererType {
	case RendererMarp:
		return NewMarpRenderer()
	case RendererRemark:
		return NewRemarkRenderer()
	}
	return nil
}
