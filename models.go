package main

import (
	"rss/internal/database"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json: "id"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
	Name      string    `json: "name"`
}

// connection that takes a user from databse and returns a user struct
func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
	}
}
