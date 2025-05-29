package domain

import (
	"time"
)

type User struct {
	ID        string    `bson:"_id,omitempty"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	FullName  string    `bson:"full_name"`
	Role      string    `bson:"role"` // default: user
	House     string    `bson:"house"`
	Street    string    `bson:"street"`
	Apartment string    `bson:"apartment"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
