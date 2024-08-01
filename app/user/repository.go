package user

import (
	"context"
	"database/sql"
	"errors"
	"go-rental/domain"
)

type Repository struct {
}

func getLastId(ctx context.Context, tx *sql.Tx, email string) *string {
	var lastId string

	query := "select id from users where email = ? order by updated_at desc;"

	rows, err := tx.QueryContext(ctx, query, email)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		err = rows.Scan(&lastId)
	}

	return &lastId
}

func (rpo *Repository) Create(ctx context.Context, tx *sql.Tx, user *domain.User) *domain.User {
	query := `insert into users (id,email,password,first_name,last_name,role,otp)
	values (?,?,?,?,?,?,?)`

	_, err := tx.ExecContext(ctx, query, user.Id, user.Email, user.Password, user.FirstName,
		user.LastName, user.Role, user.Otp)
	if err != nil {
		panic(err)
	}

	lastId := getLastId(ctx, tx, user.Email)

	user.Id = lastId

	return user
}

func (rpo *Repository) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*domain.User, error) {
	query := `select id, email, password, first_name, last_name
	from users where email = ?`

	rows, err := tx.QueryContext(ctx, query, email)
	if err != nil {
		panic(err)
	}

	user := new(domain.User)
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Email, &user.Password, &user.FirstName, &user.LastName)

		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}
