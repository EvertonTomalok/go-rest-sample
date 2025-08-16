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
func (db *memDB) Upsert(identifier int64, person entities.Person) error {
	// if person is not present on the map but max size was achieve, it will
	// return error
	if len(db.data) == db.maxSize {
		_, found := db.Get(identifier)
		if !found {
			return MaxSizeAchievedErr
		}
	}
	db.data[person.ID] = person
	return nil
}

func (db *memDB) Get(identifier int64) (entities.Person, bool) {
	item, found := db.data[identifier]
	return item, found
}

func (db *memDB) Size() int {
	return len(db.data)
}
