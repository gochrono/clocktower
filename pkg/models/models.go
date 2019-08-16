package models

import (
	"github.com/jmoiron/sqlx"
)

// Repository is the main model interface, exposing all functions to manipulate
// the different data structures.
type Repository interface {
	GetUserByUsername(username string) (*User, error)
	GetUserByID(id int) (*User, error)
	GetUsers() (Users, error)
	DeleteUserByID(id int) error
}

// DatabaseRepository is an implementation of the `Repository` interface with a
// database.
type DatabaseRepository struct {
	db *sqlx.DB
}

// NewDatabaseRepository returns an instance of `DatabaseRepository`.
func NewDatabaseRepository(db *sqlx.DB) Repository {
	return DatabaseRepository{
		db,
	}
}
