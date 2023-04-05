package repository

type BaseRepository[T any] interface {
	Search(by map[string]interface{}) ([]T, error)
	List() ([]T, error)
	Get(id string) (*T, error)
	Save(payload *T) error
	Delete(id string) error
}

type BaseRepositoryEmailPhone[T any] interface {
	GetByEmail(email string) (*T, error)
	GetByPhone(phone string) (*T, error)
}

type BaseRepositoryCount[T any] interface {
	CountData(fieldName string, id string) error
}
