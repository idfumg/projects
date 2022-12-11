package config

import (
	"fmt"
	"path/filepath"
	"text/template"
)

var (
	TemplatesPath     = "./templates"
	TemplatesTestPath = "./../../templates"
)

type Cache struct {
	Values map[string]*template.Template
	Path   string
}

func getTemplatePath(basePath, name string) string {
	return fmt.Sprintf("%s/%s", basePath, name)
}

func NewCache(path string) (*Cache, error) {
	pages, err := filepath.Glob(getTemplatePath(path, "*page.gohtml"))
	if err != nil {
		return nil, err
	}

	cache := Cache{
		Values: map[string]*template.Template{},
		Path: path,
	}

	for _, page := range pages {
		baseName := filepath.Base(page)
		ts, err := template.New(baseName).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		layouts, err := filepath.Glob(getTemplatePath(path, "*layout.gohtml"))
		if err != nil {
			return nil, err
		}

		if len(layouts) > 0 {
			ts, err = ts.ParseGlob(getTemplatePath(path, "*layout.gohtml"))
			if err != nil {
				return nil, err
			}
		}

		cache.Values[baseName] = ts
	}

	return &cache, nil
}
