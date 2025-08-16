package ports

import "github.com/evertontomalok/go-rest-sample/internal/app/domain/entities"

//go:generate mockgen -source=./repository.go -destination=./repository_mock.go -package=ports
type Repository interface {
	Upsert(key string, value entities.Person) error
	Get(key string) (entities.Person, bool)
	Delete(key string) error
	Size() int
}
