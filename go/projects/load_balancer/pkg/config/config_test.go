package config

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	data := `
services:
    - name: "Test service"
      strategy: "RoundRobin"
      matcher: "/api/v1"
      replicas:
        - "localhost:8081"
        - "localhost:8082"
`

	config, err := New(strings.NewReader(data))
	assert.NoError(t, err)
	assert.Equal(t, 1, len(config.Services))
	assert.Equal(t, "RoundRobin", config.Services[0].Strategy)
	assert.Equal(t, "Test service", config.Services[0].Name)
	assert.Equal(t, "/api/v1", config.Services[0].Matcher)
	assert.Equal(t, []string{"localhost:8081", "localhost:8082"}, config.Services[0].Replicas)
}
