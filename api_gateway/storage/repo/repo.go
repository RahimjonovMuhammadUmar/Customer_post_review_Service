package repo

type InMemoryStorageI interface {
	Set(key, value string) error
	SetWithTTl(key, value string, duration int64) error
	Get(key string) (interface{}, error)
}
