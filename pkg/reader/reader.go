package reader

// Reader defines an interface for reading from an input
type Reader interface {
	Open() error
	Read() ([]string, error)
	Close() error
}
