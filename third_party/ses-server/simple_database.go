package ses_server

import (
	"log"
)

type worldSimplestDatabase struct {
	data map[string]any
}

func newSimpleDatabase() PersistenceStorage {
	return &worldSimplestDatabase{
		data: make(map[string]any),
	}
}

func (w *worldSimplestDatabase) Write(key string, value any) error {
	w.data[key] = value
	return nil
}

func (w *worldSimplestDatabase) Read(key string) any {
	return w.data[key]
}

func (w *worldSimplestDatabase) Init() {
	// TODO: read from file
	w.data = make(map[string]any)
	log.Println("init worldSimplestDatabase")

}

func (w *worldSimplestDatabase) Shutdown() {
	// TODO: write to file
	log.Println("shutting down worldSimplestDatabase")
}
