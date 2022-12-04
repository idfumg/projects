package cache

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"myapp/server/rest"
	"path/filepath"
)

type Cache struct {
	data     map[string]*template.Template
	useCache bool
}

func New(useCache bool) (*Cache, error) {
	c := &Cache{
		data:     map[string]*template.Template{},
		useCache: useCache,
	}

	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return c, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return c, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return c, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return c, err
			}
		}

		c.data[name] = ts
	}

	return c, nil
}

func (c *Cache) Render(w io.Writer, name string, td *rest.TemplateData) error {
	data := c.data
	if !c.useCache {
		cache2, err := New(c.useCache)
		if err != nil {
			return err
		}
		data = cache2.data
	}

	t, ok := data[name]

	if !ok {
		return fmt.Errorf("couldn't find template with the name: %s", name)
	}

	buf := &bytes.Buffer{}

	err := t.Execute(buf, td)
	if err != nil {
		return err
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		return err
	}

	return nil
}
