package user

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID 	  uuid.UUID `gorm:"unique;not null"`
	Email string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewUser(email string) User{
	return User{
		ID: uuid.New(),
		Email: email,
	}
}