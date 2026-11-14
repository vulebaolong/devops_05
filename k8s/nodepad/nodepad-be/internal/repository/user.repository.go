package repository

import (
	"context"
	"fmt"
	"strings"

	"nodepad-be/ent"
	"nodepad-be/ent/user"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, email, password string) (*ent.User, error)
	FindByEmail(ctx context.Context, email string) (*ent.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*ent.User, error)
}

type userRepository struct {
	db *ent.Client
}

func NewUserRepository(db *ent.Client) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, email, password string) (*ent.User, error) {
	item, err := GetClient(ctx, r.db).User.Create().
		SetEmail(strings.ToLower(strings.TrimSpace(email))).
		SetPassword(password).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return item, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*ent.User, error) {
	item, err := GetClient(ctx, r.db).User.Query().
		Where(user.EmailEQ(strings.ToLower(strings.TrimSpace(email)))).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("find user by email: %w", err)
	}
	return item, nil
}

func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	item, err := GetClient(ctx, r.db).User.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find user by id: %w", err)
	}
	return item, nil
}
