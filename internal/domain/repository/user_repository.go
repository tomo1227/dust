package repository

import (
	"context"

	"github.com/tomo1227/dust/internal/domain/models"
)

type UserRepository interface {
	Create(ctx context.Context, input models.User) error
	CreateAll(ctx context.Context, input models.User) error
}
