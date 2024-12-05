package user

import (
	"context"
	"github.com/MaluMSiza/sqlc-api-go/database"
	"github.com/jackc/pgx/v5/pgtype"
	"errors"
	"regexp"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func CPFValidator(cpf string) error {
	re := regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}-\d{2}$`)

	if !re.MatchString(cpf) {
		return errors.New("invalid CPF format, expected XXX.XXX.XXX-XX")
	}

	return nil
}

func (s *UserService) CreateUser(ctx context.Context, name string, birthDate pgtype.Date, cpf string) (database.User, error) {
	if err := CPFValidator(cpf); err != nil {
		return database.User{}, err
	}

	return s.repo.CreateUser(ctx, name, birthDate, cpf)
}

func (s *UserService) GetUser(ctx context.Context, id int32) (database.User, error) {
	return s.repo.GetUser(ctx, id)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]database.User, error) {
	return s.repo.GetAllUsers(ctx)
}

func (s *UserService) GetUserAge(ctx context.Context, id int32) (float64, error) {
	return s.repo.GetUserAge(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, name string, birthDate pgtype.Date, cpf string) (database.User, error) {
	if err := CPFValidator(cpf); err != nil {
		return database.User{}, err
	}
	return s.repo.UpdateUser(ctx, id, name, birthDate, cpf)
}

func (s *UserService) DeleteUser(ctx context.Context, id int32) (database.User, error) {
	return s.repo.DeleteUser(ctx, id)
}
