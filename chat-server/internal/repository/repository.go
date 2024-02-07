package repository

import (
	"context"

	"github.com/gogapopp/microservices/chat-server/internal/models"
)

type Repository interface {
	Create(ctx context.Context, message models.Message) error
}
