package infra

import (
	"testing"

	"github.com/evertontomalok/go-rest-sample/internal/app/domain/entities"
	"github.com/stretchr/testify/assert"
)

func parsePerson(name string, id int64, age int) entities.Person {
	personData := entities.Person{
		Name: name,
		ID:   id,
		Age:  age,
	}

	return personData
}

func Test_MemDB(t *testing.T) {
	key := int64(1)
	personData := parsePerson("Everton Tomalok", key, 25)

	memDB := NewMemDB()
	t.Run("insert person", func(t *testing.T) {
		err := memDB.Insert(personData)
		assert.Nil(t, err)

		value, found := memDB.Get(key)
		assert.True(t, found)
		assert.Equal(t, "Everton Tomalok", value.Name)
	})

	t.Run("update person", func(t *testing.T) {
		personData.Name = "Everton Tomalok Updated" // overwritten name

		err := memDB.Update(personData)
		assert.Nil(t, err)

		value, found := memDB.Get(key)
		assert.True(t, found)
		assert.Equal(t, "Everton Tomalok Updated", value.Name)
	})
}

func Test_MemDB_Insert_Limits(t *testing.T) {
	memDB := NewMemDB(WithMaxSize(1)) // set map to have max size 1

	type testCase struct {
		testName string
		name     string
		id       int64
		age      int
		mustFail bool
	}
	testCases := []testCase{
		{
			testName: "first insert must work",
			name:     "Everton Tomalok",
			id:       1,
			age:      25,
			mustFail: false,
		},
		{
			testName: "second insert must fail",
			name:     "Everton Tomalok",
			id:       2,
			age:      26,
			mustFail: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.testName, func(t *testing.T) {
			personData := parsePerson(test.name, test.id, test.age)
			err := memDB.Insert(personData)
			if test.mustFail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, test.name, personData.Name)
				assert.Equal(t, test.id, personData.ID)
				assert.Equal(t, test.age, personData.Age)
			}
		})
	}
}
