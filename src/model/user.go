package model

import "time"

type User struct {
	UserID    int       `db:"userid"` // o "user_id" según tu naming
	Email     string    `db:"email"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}
