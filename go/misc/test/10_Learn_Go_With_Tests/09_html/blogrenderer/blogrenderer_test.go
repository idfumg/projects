package blogrenderer_test

import (
	"bytes"
	"testing"

	"myapp/blogrenderer"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	var post = blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		renderer, err := blogrenderer.NewPostRenderer()
		if err != nil {
			t.Fatal(err)
		}

		buf := bytes.Buffer{}
		err = renderer.Render(&buf, post)
		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
