package service

import (
	"github.com/DenHax/mobile-songs/internal/domain/models"
	"github.com/DenHax/mobile-songs/internal/repo"
	"github.com/DenHax/mobile-songs/internal/service/song"
)

type Song interface {
	Song(id int) (models.Song, error) // GET
	Create(song models.Song) (int, error)
	Delete(id int)
	Update(id int, update models.UpdateSong) error
}

type SongsList interface {
	SongList(id int) (models.SongsList, error)
	Create(list models.SongsList) (int, error)
	Update(id int, update models.SongsList) error
}

type Service struct {
	Song
	SongsList
}

func NewService(repos *repo.Repository) *Service {
	return &Service{
		Song:      song.NewSongService(repos.Song, repos.SongsList),
		SongsList: song.NewSongsListService(repos.SongsList),
	}
}
