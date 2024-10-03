package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	SongTable = "Song"
)

type Storage struct {
	DB *sqlx.DB
}

func New(connUrl string) (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := sqlx.Open("postgres", connUrl)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	fmt.Println("Connection!")

	return &Storage{DB: db}, nil
}

func (s *Storage) Close() error {
	fmt.Println("Close storage")
	return s.DB.Close()
}
