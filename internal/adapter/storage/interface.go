package storage

type IStorage interface {
	Get(key string) (int64, error)
	Set(key string, value int64) error
}
