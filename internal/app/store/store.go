package store

import (
	"database/sql"
	_ "github.com/lib/pq" //anonym import for to avoid methods imports from package
)

type Store struct {
	config *Config
	db *sql.DB
}
// New create new store
func New(config *Config) *Store {
	return &Config{
		config: config,
	}
}
// Open use for init of store, try to connect to store and return error
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}
// Close for disconnect from store and some operations after server finished his work
func (s *Store) Close() {
	//...
}