package chat_gpt_ppt

import (
	"os"
	"testing"

	"github.com/williamfzc/chat-gpt-ppt/assets"
)

func TestRender(t *testing.T) {
	renderer := NewMarpRenderer()
	tmpF, err := os.CreateTemp(os.TempDir(), "tmpMarp*")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpF.Name())
	_ = os.WriteFile(tmpF.Name(), assets.Static, 0755)
	renderer.SetBinPath(tmpF.Name())

	renderer.AddTopic(&Topic{
		Title:   "title",
		Content: "content",
	})
	err = renderer.RenderFile("ok.html")
}
