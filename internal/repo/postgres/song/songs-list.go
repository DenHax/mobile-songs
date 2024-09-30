package song

import (
	"database/sql"

	"github.com/DenHax/mobile-songs/internal/domain/models"
)

type SongsListPsql struct {
	db *sql.DB
}

func NewSongsListPsql(db *sql.DB) *SongsListPsql {
	return &SongsListPsql{db: db}
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
