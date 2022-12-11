package config

import (
	"encoding/gob"
	"fmt"
	"log"
	"myapp/pkg/models"
	"os"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

var (
	Production = true
	Test       = false
)

type AppConfig struct {
	UseCache      bool
	TemplateCache *Cache
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}

func New(inProduction bool) (*AppConfig, error) {
	templatesPath := TemplatesTestPath
	if inProduction {
		templatesPath = TemplatesPath
	}

	cache, err := NewCache(templatesPath)
	if err != nil {
		return nil, err
	}

	session, err := NewSession(inProduction)
	if err != nil {
		return nil, err
	}

	gob.Register(models.Reservation{})

	return &AppConfig{
		UseCache:      inProduction,
		TemplateCache: cache,
		InfoLog:       log.New(os.Stdout, "myapp", log.Lshortfile|log.Lmicroseconds),
		InProduction:  inProduction,
		Session:       session,
	}, nil
}

func (c *AppConfig) GetTemplate(name string) (*template.Template, error) {
	t, ok := c.getCache().Values[name]
	if !ok {
		return nil, fmt.Errorf("page %s was not found in the cache", name)
	}
	return t, nil
}

func (c *AppConfig) getCache() *Cache {
	if c.UseCache {
		return c.TemplateCache
	} else if cache, err := NewCache(c.TemplateCache.Path); err == nil {
		return cache
	} else {
		c.InfoLog.Fatalln(err)
	}
	return nil
}
