package data

import "html/template"

const (
	TRUNCATED_TEXT_SIZE = 150
)

type Page struct {
	Id       int64
	Title    string
	Content  template.HTML
	Date     string
	GUID     string
	Comments []*Comment
	Session  Session
}

func (p *Page) TruncatedContent() string {
	if len(p.Content) > TRUNCATED_TEXT_SIZE {
		return string(p.Content[:TRUNCATED_TEXT_SIZE]) + "..."
	}
	return string(p.Content)
}
