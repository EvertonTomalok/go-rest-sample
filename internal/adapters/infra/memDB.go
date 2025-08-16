package infra

import (
	"github.com/evertontomalok/go-rest-sample/internal/app/domain/entities"
)

func NewMemDB(options ...Option) *memDB {
	config := &config{maxLen: 1048576} // default max len is 2 ** 20
	for _, o := range options {
		o(config)
	}

	db := &memDB{data: make(map[int64]entities.Person), maxSize: config.maxLen}
	return db
}

// In memory map struct database with a maxSize attribute to control
// max size of the database.
type memDB struct {
	data    map[int64]entities.Person
	maxSize int
}

// Upsert the database person collection.
func (db *memDB) Insert(person entities.Person) error {
	if db.Size() >= db.maxSize {
		return MaxSizeAchievedErr
	}
	_, found := db.Get(person.ID)
	if found {
		return RecordExistsErr
	}
	db.data[person.ID] = person
	return nil
}

func (db *memDB) Update(person entities.Person) error {
	_, found := db.Get(person.ID)
	if !found {
		return RecordNotFoundErr
	}
	db.data[person.ID] = person
	return nil
}

func (db *memDB) Get(identifier int64) (entities.Person, bool) {
	item, found := db.data[identifier]
	return item, found
}

func (db *memDB) Delete(key int64) error {
	return nil
}

func (db *memDB) Size() int {
	return len(db.data)
}
