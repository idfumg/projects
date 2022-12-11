package render

import (
	"myapp/pkg/config"
	"net/http"
	"testing"
)

type myWriter struct{}

func (w myWriter) Header() http.Header         { return http.Header{} }
func (w myWriter) WriteHeader(status int)      {}
func (w myWriter) Write(b []byte) (int, error) { return len(b), nil }

func TestAddDefaultData(t *testing.T) {
	app, err := config.New(config.Test)
	if err != nil {
		t.Error(err)
	}

	r, err := getSession(app)
	if err != nil {
		t.Error(err)
	}

	app.Session.Put(r.Context(), "flash", "123")

	td := addDefaultData(app, &TemplateData{}, r)
	if td == nil {
		t.Error("failed to add default data to a template data")
		return
	}
	if td.Flash != "123" {
		t.Error("flash value 123 was not found in a session")
		return
	}
}

func TestRenderTemplate(t *testing.T) {
	app, err := config.New(config.Test)
	if err != nil {
		t.Error(err)
	}

	r, err := getSession(app)
	if err != nil {
		t.Error(err)
	}

	writer := myWriter{}

	err = RenderTemplate(app, writer, r, "home.page.gohtml", &TemplateData{})
	if err != nil {
		t.Error(err)
	}

	err = RenderTemplate(app, writer, r, "?.page.gohtml", &TemplateData{})
	if err == nil {
		t.Error("rendered a template that do not exist:", err)
	}
}

func getSession(app *config.AppConfig) (*http.Request, error) {
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		return nil, err
	}

	ctx, err := app.Session.Load(r.Context(), r.Header.Get("X-Session"))
	if err != nil {
		return nil, err
	}
	r = r.WithContext(ctx)

	return r, nil
}
