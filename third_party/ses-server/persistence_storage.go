package ses_server

type PersistenceStorage interface {
	Read(key string) any
	Write(key string, value any) error
	Init()
	Shutdown()
}
