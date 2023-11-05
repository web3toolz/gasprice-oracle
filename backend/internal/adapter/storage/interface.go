package storage

type IStorage interface {
	Get(network string, key string) (int64, error)
	Set(network string, key string, value int64) error
}
