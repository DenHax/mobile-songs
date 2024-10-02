package song

import (
	"github.com/DenHax/mobile-songs/internal/domain/models"
	"github.com/DenHax/mobile-songs/internal/storage"
)

type SongsListPsql struct {
	storage *storage.Storage
}

func NewSongsListPsql(s *storage.Storage) *SongsListPsql {
	return &SongsListPsql{storage: s}
}

func (r SongsListPsql) SongList(id int) (models.SongsList, error) {
	return models.SongsList{}, nil
}
func (r SongsListPsql) Create(list models.SongsList) (int, error) {
	return 0, nil
}

func (r SongsListPsql) Update(id int, update models.SongsList) error {
	return nil
}
