package song

import (
	"github.com/DenHax/mobile-songs/internal/domain/models"
	"github.com/DenHax/mobile-songs/internal/storage"
)

type SongPsql struct {
	storage *storage.Storage
}

func NewSongPsql(s *storage.Storage) *SongPsql {
	return &SongPsql{storage: s}
}

func (r SongPsql) Song(id int) (models.Song, error) {
	return models.Song{}, nil
}

func (r SongPsql) Create(song models.Song) (int, error) {
	return 0, nil
}

func (r SongPsql) Delete(id int) {
}

func (r SongPsql) Update(id int, update models.UpdateSong) error {
	return nil
}
