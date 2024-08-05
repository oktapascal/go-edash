package company

import (
	"context"
	"database/sql"
	"errors"
	"go-rental/domain"
)

type Repository struct {
}

func getLastId(ctx context.Context, tx *sql.Tx) string {
	var lastId string

	query := "select id from companies order by updated_at desc;"

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		err = rows.Scan(&lastId)
	}

	return lastId
}

func (rpo *Repository) Create(ctx context.Context, tx *sql.Tx, company *domain.Company) *domain.Company {
	query := "insert into companies (id,name,description,employee_count) values (?,?,?,?)"

	_, err := tx.ExecContext(ctx, query, company.Id, company.Name, company.Description, company.Category)
	if err != nil {
		panic(err)
	}

	lastId := getLastId(ctx, tx)

	company.Id = lastId

	return company
}

func (rpo *Repository) Update(ctx context.Context, tx *sql.Tx, company *domain.Company) *domain.Company {
	query := "update companies set name=?,description=?,employee_count=? where id = ?"

	_, err := tx.ExecContext(ctx, query, company.Name, company.Description, company.Category, company.Id)
	if err != nil {
		panic(err)
	}

	return company
}

func (rpo *Repository) FindById(ctx context.Context, tx *sql.Tx, id string) (*domain.Company, error) {
	query := "select name,description,employee_count from companies where id =?"

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	company := new(domain.Company)

	if rows.Next() {
		err = rows.Scan(&company.Name, &company.Description, &company.Category)

		return company, nil
	} else {
		return company, errors.New("company not found")
	}
}
