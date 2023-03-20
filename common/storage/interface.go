package storage

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
