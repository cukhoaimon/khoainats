package database

type AbstractDatabase interface {
	ReadAll() map[string]any
	Read(key string) any
	Write(key string, value any) error
	Init() error
	Shutdown() error
}

func NewSimpleDatabase() AbstractDatabase {
	return newSimpleDatabase("./tmp")
}
