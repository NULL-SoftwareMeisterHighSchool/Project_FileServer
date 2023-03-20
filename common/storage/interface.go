package storage

import "errors"

type Storage interface {
	// PUBLIC
	Read(id string) ([]byte, error)
	// TODO: declare some types for io.Copy or something
	Create() error
	Update() error
	Delete() error

	// PRIVATE
	exists(id string) bool
}

var (
	ErrFileNotFound = errors.New("file not found")
	ErrFileExists   = errors.New("file with that id already exists")
)
