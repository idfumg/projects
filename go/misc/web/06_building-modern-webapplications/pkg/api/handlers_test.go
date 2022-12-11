package api

import (
	"log"
	"myapp/pkg/config"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var tests = []struct {
	name       string
	url        string
	method     string
	params     []postData
	statusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"room1", "/room1", "GET", []postData{}, http.StatusOK},
	{"room2", "/room2", "GET", []postData{}, http.StatusOK},
	{"reservation", "/reservation", "GET", []postData{}, http.StatusOK},
	{"availability", "/availability", "GET", []postData{}, http.StatusOK},
	{"post-availability", "/availability", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"post-availability-json", "/availability-json", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"post-reservation", "/reservation", "POST", []postData{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "Smith"},
		{key: "email", value: "me@example.com"},
		{key: "phone", value: "+7(999)1234567"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	app, err := config.New(config.Test)
	if err != nil {
		log.Fatalln("Couldn't load a config", err)
	}
	routes := Routes(app)

	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, test := range tests {
		if test.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + test.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != test.statusCode {
				t.Errorf("for %s, expected %d, but got %d", test.name, test.statusCode, resp.StatusCode)
			}
		} else if test.method == "POST" {
			values := url.Values{}
			for _, params := range test.params {
				values.Add(params.key, params.value)
			}

			resp, err := ts.Client().PostForm(ts.URL + test.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != test.statusCode {
				t.Errorf("for %s, expected %d, but got %d", test.name, test.statusCode, resp.StatusCode)
			}
		}
	}
}
