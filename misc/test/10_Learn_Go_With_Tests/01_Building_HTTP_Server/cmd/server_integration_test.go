package main

import (
	"myapp/server"
	"myapp/service"
	"myapp/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := server.NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), utils.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), utils.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), utils.NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, utils.NewGetScoreRequest(player))

		utils.AssertStatus(t, response.Code, http.StatusOK)
		utils.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, utils.NewGetLeagueRequest())

		utils.AssertStatus(t, response.Code, http.StatusOK)

		got := utils.GetLeagueFromResponse(t, response.Body)
		want := []service.Player{
			{Name: "Pepper", Wins: 3},
		}
		utils.AssertLeague(t, got, want)
	})
}
