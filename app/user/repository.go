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
	query := `insert into users (id,email,password,first_name,last_name,role,otp,otp_expired_time)
	values (?,?,?,?,?,?,?,?)`

	_, err := tx.ExecContext(ctx, query, user.Id, user.Email, user.Password, user.FirstName,
		user.LastName, user.Role, user.Otp, user.OtpExpiredTime)
	if err != nil {
		panic(err)
	}

	lastId := getLastId(ctx, tx, user.Email)

	user.Id = lastId

	return user
}

func (rpo *Repository) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*domain.User, error) {
	query := `select id, email, password, first_name, last_name, otp, otp_expired_time
	from users where email = ?`

	rows, err := tx.QueryContext(ctx, query, email)
	if err != nil {
		panic(err)
	}

	user := new(domain.User)
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Otp,
			&user.OtpExpiredTime)

		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}
