package storage

import (
	"canchitas-libres-field/internal/pkg/domain"
	"context"
	"fmt"
)

const (
	queryInsertUser = `
        INSERT INTO field (name, number)
        VALUES ($1, $2);`
	querySelectAllFields = `SELECT id, name, number FROM field`
)

func (p *Postgres) GetAll() ([]domain.Field, error) {
	var fields []domain.Field

	err := p.Select(&fields, querySelectAllFields)
	if err != nil {
		return nil, err
	}
	return fields, nil

}

func (p *Postgres) Add(ctx context.Context, field domain.Field) error {
	fmt.Println("in infrastructure layer we have a field whit name: ", field.Name)
	tx, err := p.Begin()
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, queryInsertUser, field.Name, field.Number)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetByID(id int) (domain.Field, error) {
	return domain.Field{}, nil
}

func (p *Postgres) Delete(ctx context.Context, id string) error {
	return nil
}

/*func (p *Postgres) Add(ctx context.Context, field Field) error {
	tx, err := p.Begin()
	if err != nil {
		return application.InternalServerError
	}
	_, err = tx.ExecContext(ctx, queryInsertUser, name, email, password)
	if err != nil {
		tx.Rollback()
		return application.InternalServerError
	}

	err = tx.Commit()
	if err != nil {
		return application.InternalServerError
	}
	return nil
}*/
