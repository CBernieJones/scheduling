package model

import "time"

type User struct {
	UserID    int       `db:"userid"`
	Email     string    `db:"email"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"createdat"`
}
