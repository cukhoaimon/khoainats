package ses_server

type PersistenceStorage interface {
	ReadAll() map[string]any
	Read(key string) any
	Write(key string, value any) error
	Init() error
	Shutdown() error
}
