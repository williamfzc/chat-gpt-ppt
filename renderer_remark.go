package chat_gpt_ppt

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

const RemarkTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <title>Title</title>
    <meta charset="utf-8">
    <style>
      @import url(https://fonts.googleapis.com/css?family=Yanone+Kaffeesatz);
      @import url(https://fonts.googleapis.com/css?family=Droid+Serif:400,700,400italic);
      @import url(https://fonts.googleapis.com/css?family=Ubuntu+Mono:400,700,400italic);

      body { font-family: 'Droid Serif'; }
      h1, h2, h3 {
        font-family: 'Yanone Kaffeesatz';
        font-weight: normal;
      }
      .remark-code, .remark-inline-code { font-family: 'Ubuntu Mono'; }
    </style>
  </head>
  <body>
    <textarea id="source">

class: center, middle

%s

    </textarea>
    <script src="https://remarkjs.com/downloads/remark-latest.min.js">
    </script>
    <script>
      var slideshow = remark.create();
    </script>
  </body>
</html>
`

type RemarkRenderer struct {
	Topics []*Topic
}

func NewRemarkRenderer() Renderer {
	return &RemarkRenderer{
		Topics: make([]*Topic, 0),
	}
}

func (r *RemarkRenderer) AddTopic(topic *Topic) {
	r.Topics = append(r.Topics, topic)
}

func (r *RemarkRenderer) RenderFile(outputPath string) error {
	output, err := r.RenderString()
	if err != nil {
		return err
	}
	err = os.WriteFile(outputPath, []byte(output), fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (r *RemarkRenderer) SetBinPath(_ string) {
}

func (r *RemarkRenderer) RenderString() (string, error) {
	parts := make([]string, 0)
	for _, each := range r.Topics {
		parts = append(parts, each.ToMarkdown())
	}
	final := strings.Join(parts, "\n---\n")
	outputContent := fmt.Sprintf(RemarkTemplate, final)
	return outputContent, nil
}
