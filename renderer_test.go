package chat_gpt_ppt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemark(t *testing.T) {
	r := GetRenderer(RendererRemark)
	assert.NotNil(t, r)
	r.AddTopic(&Topic{
		Title:   "title",
		Content: "content",
	})
	err := r.RenderFile("ok.html")
	assert.Nil(t, err)
}
