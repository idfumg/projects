package render

import (
	"bytes"
	"myapp/pkg/config"
	"net/http"

	"github.com/justinas/nosurf"
)

func RenderTemplate(appConfig *config.AppConfig, w http.ResponseWriter, r *http.Request, name string, td *TemplateData) error {
	t, err := appConfig.GetTemplate(name)
	if err != nil {
		return err
	}

	td = addDefaultData(appConfig, td, r)

	buff := new(bytes.Buffer)

	err = t.Execute(buff, td)
	if err != nil {
		return err
	}

	_, err = buff.WriteTo(w)
	return err
}

func addDefaultData(appConfig *config.AppConfig, td *TemplateData, r *http.Request) *TemplateData {
	td.Flash = appConfig.Session.PopString(r.Context(), "flash")
	td.Error = appConfig.Session.PopString(r.Context(), "error")
	td.Warning = appConfig.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}
