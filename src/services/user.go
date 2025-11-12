package services

import (
	"github.com/CBernieJones/scheduling/src/model"
	repo "github.com/CBernieJones/scheduling/src/repository"
)

type UserService struct{}

func (UserService) List() ([]model.User, error) {
	return repo.RetriveUsers()
}
