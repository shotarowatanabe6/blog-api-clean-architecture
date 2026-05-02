package repository

// IDBRepository : ドメイン層が必要とするストレージ操作のポート（インターフェース）。
// 具体的なDB実装の知識を持たない。
type IDBRepository interface {
	Get(key string) (*string, error)
	Set(key string, value string) error
}
