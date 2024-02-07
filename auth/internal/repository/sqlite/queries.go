package sqlite

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/gogapopp/microservices/auth/models"
)

func (r *Repository) Create(ctx context.Context, user models.User) error {
	builderInsert := sq.Insert("users").
		PlaceholderFormat(sq.Question).
		Columns("username", "email", "roles", "created_at", "updated_at").
		Values(user.Userame, user.Email, user.Role, user.CreateAt, user.UpdateAt)
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
