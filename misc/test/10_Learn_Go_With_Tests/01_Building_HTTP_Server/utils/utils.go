package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"myapp/service"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func NewRequestPath(name string) string {
	return fmt.Sprintf("/players/%s", name)
}

func NewGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, NewRequestPath(name), nil)
	return request
}

func NewGetLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func NewPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, NewRequestPath(name), nil)
	return request
}

func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response `body` is wrong, got %q, but want %q", got, want)
	}
}

func AssertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("response `status` is wrong, got %d, but want %d", got, want)
	}
}

func GetLeagueFromResponse(t testing.TB, body io.Reader) []service.Player {
	t.Helper()
	var got []service.Player
	err := json.NewDecoder(body).Decode(&got)
	if err != nil {
		t.Fatalf("Unable to parse response from the server %q into players: %v", body, err)
	}
	return got
}

func AssertLeague(t testing.TB, got, want []service.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertContentType(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response did not have content-type of %s, got %s", want, got)
	}
}

func AssertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Wrong score received. got %d want %d", got, want)
	}
}

func CreateTempFile(t testing.TB, data string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("Could not create temp file: %v", err)
	}

	tmpfile.Write([]byte(data))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
