package users_test

import (
	"context"

	"github.com/fanfit/user-service/models/users/repository"
)

type mockService struct {
	GetByIDResponse repository.User
	CreateResponse  repository.GetClientByIDRow

	GetByIDErr error
	CreateErr  error
	DeleteErr  error
}

func (m mockService) GetByEmail(ctx context.Context, id string) (repository.User, error) {
	return m.GetByIDResponse, m.GetByIDErr
}

func (m mockService) Create(ctx context.Context, user repository.User) (repository.GetClientByIDRow, error) {
	return m.CreateResponse, m.CreateErr
}

func (m mockService) Delete(ctx context.Context, id string) error {
	return m.DeleteErr
}
