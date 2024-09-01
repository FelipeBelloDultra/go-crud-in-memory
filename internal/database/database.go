package database

import (
	"fmt"

	"github.com/google/uuid"
)

type user struct {
	FirstName string
	LastName  string
	Biography string
}

func MakeDatabase() Application {
	return Application{
		data: make(database),
	}
}

type Application struct {
	data database
}

type database map[string]user

func (db *Application) CreateUser(firstName, lastName, biography string) string {
	id, _ := uuid.NewUUID()

	user := user{
		FirstName: firstName,
		LastName:  lastName,
		Biography: biography,
	}

	db.data[id.String()] = user

	fmt.Println(db.data)

	return id.String()
}

func (db Application) ListUsers() database {
	return db.data
}

func (db Application) GetUserByID(id string) (*user, error) {
	user, ok := db.data[id]

	if !ok {
		return nil, fmt.Errorf("the user with the specified ID does not exist")
	}

	return &user, nil
}
