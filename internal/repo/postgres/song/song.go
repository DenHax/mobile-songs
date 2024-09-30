package song

import (
	"database/sql"

	"github.com/DenHax/mobile-songs/internal/domain/models"
)

type SongPsql struct {
	db *sql.DB
}

func NewSongPsql(db *sql.DB) *SongPsql {
	return &SongPsql{db: db}
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
