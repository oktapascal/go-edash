package user

import (
	"context"
	"database/sql"
	"errors"
	"go-rental/domain"
)

type Repository struct {
}

func (rpo *Repository) Create(ctx context.Context, tx *sql.Tx, user *domain.User) *domain.User {
	query := `insert into users (id_number, email, password, phone_number, address, first_name,
    last_name, role, provider, provider_id, photo_id_card)
    values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := tx.ExecContext(ctx, query, user.IdNumber, user.Email, user.Password, user.PhoneNumber, user.Address,
		user.FirstName, user.LastName, "USER", user.Provider, user.ProviderId, user.PhotoIdCard)
	if err != nil {
		panic(err)
	}

	return user
}

func (rpo *Repository) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*domain.User, error) {
	query := `select id_number, email, password, phone_number, address, first_name, last_name, photo_id_card
	from users where email = ?`

	rows, err := tx.QueryContext(ctx, query, email)
	if err != nil {
		panic(err)
	}

	user := new(domain.User)
	if rows.Next() {
		err = rows.Scan(&user.IdNumber, &user.Email, &user.Password, &user.PhoneNumber, &user.Address,
			&user.FirstName, &user.LastName, &user.PhotoIdCard)

		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}
