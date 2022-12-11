package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"myapp/service"
	utils "myapp/utils"
)

type StubServerStore struct {
	scores   map[string]int
	winCalls []string
	league   service.League
}

func (s *StubServerStore) GetPlayerScore(name string) (int, bool) {
	score, ok := s.scores[name]
	return score, ok
}

func (s *StubServerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubServerStore) GetLeague() service.League {
	return s.league
}

func TestGETPlayers(t *testing.T) {
	server := NewPlayerServer(
		&StubServerStore{
			map[string]int{
				"Pepper": 20,
				"Floyd":  10,
			},
			[]string{},
			service.League{},
		},
	)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := utils.NewGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utils.AssertStatus(t, response.Code, http.StatusOK)
		utils.AssertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := utils.NewGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utils.AssertStatus(t, response.Code, http.StatusOK)
		utils.AssertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing player in a store", func(t *testing.T) {
		request := utils.NewGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utils.AssertStatus(t, response.Code, http.StatusNotFound)
	})

	t.Run("returns 400 on missing player in a request path", func(t *testing.T) {
		request := utils.NewGetScoreRequest("")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utils.AssertStatus(t, response.Code, http.StatusBadRequest)
	})
}

func TestStoreWins(t *testing.T) {
	store := &StubServerStore{
		map[string]int{},
		[]string{},
		nil,
	}

	server := NewPlayerServer(store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Pepper"
		request := utils.NewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utils.AssertStatus(t, response.Code, http.StatusAccepted)
	})

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"
		request := utils.NewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utils.AssertStatus(t, response.Code, http.StatusAccepted)

		const CALLS_NUM = 2
		if len(store.winCalls) != CALLS_NUM {
			t.Errorf("god %d calls to RecordWin, but want %d", len(store.winCalls), CALLS_NUM)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner: got %q, but want %q", store.winCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {
	store := StubServerStore{
		league: []service.Player{
			{Name: "Cleo", Wins: 32},
			{Name: "Chris", Wins: 20},
			{Name: "Tiest", Wins: 14},
		},
	}
	server := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request := utils.NewGetLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utils.AssertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		request := utils.NewGetLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := utils.GetLeagueFromResponse(t, response.Body)
		utils.AssertStatus(t, response.Code, http.StatusOK)
		utils.AssertLeague(t, got, store.league)

		contentType := response.Result().Header.Get("content-type")
		utils.AssertContentType(t, contentType, jsonContentType)
	})
}
