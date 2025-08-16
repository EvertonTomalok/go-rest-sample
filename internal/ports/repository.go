package ports

import "github.com/evertontomalok/go-rest-sample/internal/app/domain/entities"

//go:generate mockgen -source=./repository.go -destination=./repository_mock.go -package=ports
type Repository interface {
	Update(value entities.Person) error
	Insert(value entities.Person) error
	Get(key int64) (entities.Person, bool)
	Delete(key int64) error
	Size() int
}
