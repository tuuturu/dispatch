package core

type Handler interface {
	Handle(payload []byte) error
}

type Watcher interface {
	Open() error
	Close() error

	RegisterHandler(handler Handler)
}
