package store

import (
	"encoding/json"
	"io"
	"log"
	"myapp/service"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func NewFileSystemPlayerStore(r io.ReadWriteSeeker) *FileSystemPlayerStore {
	return &FileSystemPlayerStore{
		database: r,
	}
}

func (f *FileSystemPlayerStore) GetLeague() service.League {
	f.database.Seek(0, io.SeekStart)
	var ans []service.Player
	err := json.NewDecoder(f.database).Decode(&ans)
	if err != nil {
		log.Fatalf("Unable to parse json data: %v", err)
	}
	return ans
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	}

	f.database.Seek(0, io.SeekStart)
	json.NewEncoder(f.database).Encode(league)
}
