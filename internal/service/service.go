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

type Service struct {
	Song
}

func NewService(repos *repo.Repository) *Service {
	return &Service{
		Song: song.NewSongService(repos.Song),
	}
}
