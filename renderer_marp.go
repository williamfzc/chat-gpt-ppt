package chat_gpt_ppt

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type MarpRenderer struct {
	BinPath string
	Topics  []*Topic
}

func NewMarpRenderer() Renderer {
	return &MarpRenderer{
		BinPath: "",
		Topics:  make([]*Topic, 0),
	}
}

func (m *MarpRenderer) SetBinPath(path string) {
	m.BinPath = path
}

func (m *MarpRenderer) AddTopic(topic *Topic) {
	m.Topics = append(m.Topics, topic)
}

func (m *MarpRenderer) RenderFile(outputPath string) error {
	parts := make([]string, 0)
	for _, each := range m.Topics {
		parts = append(parts, each.ToMarkdown())
	}
	final := strings.Join(parts, "\n---\n")

	// call marp
	file, err := os.CreateTemp(os.TempDir(), "cgp*.md")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())
	err = os.WriteFile(file.Name(), []byte(final), 0644)
	if err != nil {
		return err
	}

	command := exec.Command(m.BinPath, file.Name(), "-o", outputPath)
	err = command.Run()
	if err != nil {
		return err
	}
	return nil
}

func (m *MarpRenderer) RenderString() (string, error) {
	return "", fmt.Errorf("not supported")
}
