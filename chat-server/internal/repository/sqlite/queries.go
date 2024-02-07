package sqlite

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/gogapopp/microservices/chat-server/internal/models"
)

func (r *Repository) Create(ctx context.Context, message models.Message) error {
	builderInsert := sq.Insert("messages").
		PlaceholderFormat(sq.Question).
		Columns("message_from", "message_text", "timestamp").
		Values(message.From, message.Text, message.Timestamp)
	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Get() {

}

func (r *Repository) Update() {

}

func (r *Repository) Delete() {

}
