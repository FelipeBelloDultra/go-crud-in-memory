package database

import (
	"errors"

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

	return id.String()
}

func (db Application) ListUsers() database {
	return db.data
}

func (db Application) GetUserByID(id string) (*user, error) {
	user, ok := db.data[id]

	if !ok {
		return nil, errors.New("the user with the specified ID does not exist")
	}

	return &user, nil
}

func (db *Application) UpdateUser(id string, firstName, lastName, biography string) (user, error) {
	_, err := db.GetUserByID(id)
	if err != nil {
		return user{}, err
	}

	db.data[id] = user{
		FirstName: firstName,
		LastName:  lastName,
		Biography: biography,
	}

	return db.data[id], nil
}

func (db *Application) DeleteUser(id string) error {
	_, ok := db.data[id]

	if !ok {
		return errors.New("the user with the specified ID does not exist")
	}

	delete(db.data, id)

	return nil
}
