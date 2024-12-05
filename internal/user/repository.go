package user

import (
	"context"
	"github.com/MaluMSiza/sqlc-api-go/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	Queries *database.Queries
}

func NewUserRepository(db *database.Queries) *UserRepository {
	return &UserRepository{
		Queries: db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, name string, birthDate pgtype.Date, cpf string) (database.User, error) {
	return r.Queries.InsertUser(ctx, database.InsertUserParams{
		Name:      name,
		BirthDate: birthDate,
		Cpf:       cpf,
	})
}

func (r *UserRepository) GetUser(ctx context.Context, id int32) (database.User, error) {
	return r.Queries.GetUser(ctx, id)
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]database.User, error) {
	return r.Queries.GetAllUsers(ctx)
}

func (r *UserRepository) GetUserAge(ctx context.Context, id int32) (float64, error) {
	return r.Queries.GetUserAge(ctx, id)
}

func (r *UserRepository) UpdateUser(ctx context.Context, id int32, name string, birthDate pgtype.Date, cpf string) (database.User, error) {
	return r.Queries.UpdateUser(ctx, database.UpdateUserParams{
		ID:        id,
		Name:      name,
		BirthDate: birthDate,
		Cpf:       cpf,
	})
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int32) (database.User, error) {
	return r.Queries.DeleteUser(ctx, id)
}
