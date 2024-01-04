package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgressStore, error) {
	connstr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connstr)

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStore{
		db: db,
	}, nil
}
func (s *PostgressStore) CreateAccount(*Account) error {
	return nil
}
func (s *PostgressStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostgressStore) DeleteAccount(id int) error {
	return nil
}
func (s *PostgressStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
