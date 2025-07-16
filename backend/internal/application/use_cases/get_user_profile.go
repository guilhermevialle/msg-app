package use_cases

import (
	"app/backend/internal/domain/entities"
	"app/backend/internal/infra/repositories"
	"errors"
)

type GetUserProfile struct {
	userRepo repositories.IUserRepository
}

func NewGetUserProfile(ur repositories.IUserRepository) *GetUserProfile {
	return &GetUserProfile{
		userRepo: ur,
	}
}

func (uc *GetUserProfile) Execute(id string) (*entities.Profile, error) {
	u := uc.userRepo.FindById(id)
	if u == nil {
		return nil, errors.New("user not found")
	}

	return u.Profile, nil
}
