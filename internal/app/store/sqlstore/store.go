package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq" // Postgres driver
)

// Store ...
type Store struct {
	db *sql.DB
	userRepository *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		sqlstore: s,
	}

	return s.userRepository
}