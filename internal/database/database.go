package database

import "github.com/google/uuid"

type ID uuid.UUID

type User struct {
	FirstName string
	LastName  string
	Biography string
}

func MakeDatabase() Database {
	return make(map[ID]User)
}

type Database map[ID]User
