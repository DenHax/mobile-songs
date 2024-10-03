package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(connUrl string) (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := sql.Open("postgres", connUrl)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	fmt.Println("Connection!")

	return &Storage{db: db}, nil
}

func (s *Storage) Close() error {
	fmt.Println("Close storage")
	return s.db.Close()
}
