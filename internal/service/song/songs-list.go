package song

import (
	"github.com/DenHax/mobile-songs/internal/domain/models"
	"github.com/DenHax/mobile-songs/internal/repo"
)

type SongsListService struct {
	repo repo.SongsList
}

func NewSongsListService(repo repo.SongsList) *SongsListService {
	return &SongsListService{repo: repo}
}

func (s *SongsListService) SongList(id int) (models.SongsList, error) {
	return models.SongsList{}, nil
}
func (s *SongsListService) Create(list models.SongsList) (int, error) {
	return 0, nil
}
func (s *SongsListService) Update(id int, update models.SongsList) error {
	return nil
}
