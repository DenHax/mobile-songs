package song

import (
	"github.com/DenHax/mobile-songs/internal/domain/models"
	"github.com/DenHax/mobile-songs/internal/repo"
)

type SongService struct {
	repo     repo.Song
	listRepo repo.SongsList
}

func NewSongService(repo repo.Song, listRepo repo.SongsList) *SongService {
	return &SongService{repo: repo, listRepo: listRepo}
}

func (s *SongService) Song(id int) (models.Song, error) {
	return models.Song{}, nil
}
func (s *SongService) Create(song models.Song) (int, error) {
	return 0, nil
}
func (s *SongService) Delete(id int) {

}
func (s *SongService) Update(id int, update models.UpdateSong) error {
	return nil
}
