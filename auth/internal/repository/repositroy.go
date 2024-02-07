package repository

import (
	"context"

	"github.com/gogapopp/microservices/auth/models"
)

type Repository interface {
	Create(ctx context.Context, user models.User) error
}
