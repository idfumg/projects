package store

import (
	"myapp/service"
	"myapp/utils"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		data := `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`

		database, cleanDatabase := utils.CreateTempFile(t, data)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.GetLeague()

		want := []service.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}

		utils.AssertLeague(t, got, want)

		got = store.GetLeague()
		utils.AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		data := `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`

		database, cleanDatabase := utils.CreateTempFile(t, data)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Chris")
		want := 33

		utils.AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		data := `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`

		database, cleanDatabase := utils.CreateTempFile(t, data)
		defer cleanDatabase()
		
		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34

		utils.AssertScoreEquals(t, got, want)
	})
}
