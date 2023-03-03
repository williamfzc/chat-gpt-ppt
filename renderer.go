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

type Renderer interface {
	AddTopic(*Topic)
	RenderFile(outputPath string) error
	SetBinPath(path string)
}
