package repo

import (
	"database/sql"

	"github.com/DenHax/mobile-songs/internal/domain/models"
	"github.com/DenHax/mobile-songs/internal/repo/postgres/song"
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

type Repository struct {
	Song
	SongsList
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Song:      song.NewSongPsql(db),
		SongsList: song.NewSongsListPsql(db),
	}
}
