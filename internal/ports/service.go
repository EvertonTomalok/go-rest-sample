package ports

import "github.com/evertontomalok/go-rest-sample/internal/app/domain/entities"

//go:generate mockgen -source=./service.go -destination=./service_mock.go -package=ports
type Service interface {
	Upsert(person entities.Person) error
	Get(personId int64) (entities.Person, error)
	Delete(personId int64) error
}
