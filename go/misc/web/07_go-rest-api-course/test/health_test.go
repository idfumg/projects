package test

import (
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/health")
	if err != nil {
		t.Fatal(err)
		t.Fail()
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode())
}
