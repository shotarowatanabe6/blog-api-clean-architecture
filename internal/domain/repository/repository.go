package repository

type IDBRepository interface {
	Set(key string, value string)
	Get(key string) (any, bool)
}
