package repository

import (
	"github.com/CBernieJones/scheduling/src/db"
	"github.com/CBernieJones/scheduling/src/model"
)

func RetriveUsers() ([]model.User, error) {
	users := []model.User{}

	err := db.DB.Select(&users, "SELECT userid, email, name, createdat FROM users")

	return users, err
}
