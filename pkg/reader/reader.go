package reader

type Reader interface {
	Open() error
	Read() ([]string, error)
	Close() error
}
