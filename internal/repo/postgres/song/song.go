package song

import (
	"fmt"

	"github.com/DenHax/mobile-songs/internal/domain/models"
	"github.com/DenHax/mobile-songs/internal/storage"
)

type SongPsql struct {
	storage *storage.Storage
}

func NewSongPsql(s *storage.Storage) *SongPsql {
	return &SongPsql{storage: s}
}

func (r *SongPsql) Song(id int) (models.Song, error) {
	var song models.Song
	query := fmt.Sprintf(`SELECT id, song, group, lyrics FROM %s`,
		storage.SongTable)
	if err := r.storage.DB.Get(&song, query); err != nil {
		return song, err
	}
	return song, nil
}

func (r *SongPsql) GetAll() ([]models.Song, error) {
	var songs []models.Song
	query := fmt.Sprintf(`SELECT id, song, group, lyrics FROM %s`,
		storage.SongTable)
	if err := r.storage.DB.Select(&songs, query); err != nil {
		return nil, err
	}
	return songs, nil

}

func (r *SongPsql) Create(song models.Song) (int, error) {
	tx, err := r.storage.DB.Begin()
	if err != nil {
		return 0, err
	}
	var songId int
	createSongQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", storage.SongTable)

	row := tx.QueryRow(createSongQuery, song.Name, song.Group)
	err = row.Scan(&songId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return songId, tx.Commit()
}

func (r *SongPsql) Delete(id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`,
		storage.SongTable)
	_, err := r.storage.DB.Exec(query, id)
	return err
}

func (r *SongPsql) Update(id int, update models.UpdateSong) error {
	return nil
}
