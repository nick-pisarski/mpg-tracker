package repositories

type Repository[T interface{}] interface {
	Create(t T) (T, error)
	Get() ([]T, error)
	GetById(id uint) (T, error)
	Put(t T) (T, error)
	Delete(id uint) (T, error)
}
