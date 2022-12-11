package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type Post struct {
	Title       string
	Description string
	Body        string
	Tags        []string
}

type PostRenderer struct {
	t *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	t, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	
	return &PostRenderer{t: t}, nil
}

func (p *PostRenderer) Render(w io.Writer, post Post) error {
	if err := p.t.ExecuteTemplate(w, "blog.gohtml", post); err != nil {
		return err
	}

	return nil
}
