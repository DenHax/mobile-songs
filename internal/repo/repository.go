package repo

import (
	"github.com/DenHax/mobile-songs/internal/domain/models"
	"github.com/DenHax/mobile-songs/internal/repo/postgres/song"
	"github.com/DenHax/mobile-songs/internal/storage"
)

type Song interface {
	Song(id int) (models.Song, error) // GET
	Create(song models.Song) (int, error)
	Delete(id int)
	Update(id int, update models.UpdateSong) error
}

type Repository struct {
	Song
}

func NewRepository(s *storage.Storage) *Repository {
	return &Repository{
		Song: song.NewSongPsql(s),
	}
}
