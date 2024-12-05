package user

import (
	"context"
	"github.com/MaluMSiza/sqlc-api-go/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, name string, birthDate pgtype.Date, cpf string) (database.User, error) {
	return s.repo.CreateUser(ctx, name, birthDate, cpf)
}

func (s *UserService) GetUser(ctx context.Context, id int32) (database.User, error) {
	return s.repo.GetUser(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, name string, birthDate pgtype.Date, cpf string) (database.User, error) {
	return s.repo.UpdateUser(ctx, id, name, birthDate, cpf)
}

func (s *UserService) DeleteUser(ctx context.Context, id int32) (database.User, error) {
	return s.repo.DeleteUser(ctx, id)
}
