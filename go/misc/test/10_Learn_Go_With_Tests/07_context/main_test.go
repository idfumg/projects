package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

type StubStore struct {
	response string
}

func (s *StubStore) Fetch(ctx context.Context) (string, error) {
	return s.response, nil
}

func (s *StubStore) Cancel() {

}

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	select {
	case <-time.After(100 * time.Millisecond):
		return s.response, nil
	case <-ctx.Done():
		s.Cancel()
		return "", ctx.Err()
	}
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func assertWasCancelled(t testing.TB, s *SpyStore) {
	t.Helper()
	if !s.cancelled {
		t.Errorf("store was not told to cancel")
	}
}

func assertWasNotCancelled(t testing.TB, s *SpyStore) {
	t.Helper()
	if s.cancelled {
		t.Errorf("store was told to cancel")
	}
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err == nil {
			fmt.Fprint(w, data)
		}
	}
}

type SpyWriterResponse struct {
	written bool
}

func (s *SpyWriterResponse) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyWriterResponse) Write(b []byte) (int, error) {
	s.written = true
	return 0, nil
}

func (s *SpyWriterResponse) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, cancelled: false}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %q, want %q", response.Body.String(), data)
		}

		assertWasNotCancelled(t, store)
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, cancelled: false}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(1*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		assertWasCancelled(t, store)
	})

	t.Run("nothing will be written to the response on an error", func(t *testing.T) {
		data := "heelo, world"
		store := &SpyStore{response: data, cancelled: false}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(1*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyWriterResponse{written: false}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Errorf("a response should not have been written")
		}
	})
}
