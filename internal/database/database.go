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
		data: make(map[string]user),
	}
}

type Application struct {
	data map[string]user
}

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
