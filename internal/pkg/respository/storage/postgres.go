package storage

import (
	"canchitas-libres-field/internal/pkg/domain/person"
	user "canchitas-libres-field/internal/pkg/domain/user"
	"context"
)

const querySelectAllUsers = `
SELECT 
    u.id AS id,
    p.firstname AS firstname,
    p.lastname AS lastname,
    p.dni AS dni,
    p.birthdate AS birthdate,
    u.email AS email,
    u.password AS password,
    u.active AS active,
    u.role AS role
FROM 
    "user" u
INNER JOIN 
    person p ON u.id = p.id;
`

func (p *Postgres) GetAllUsers(ctx context.Context) ([]user.User, error) {
	rows, err := p.QueryContext(ctx, querySelectAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var u user.User
		var p person.Person

		// Escanear los datos desde la fila
		err = rows.Scan(
			&u.ID,
			&p.FirstName,
			&p.LastName,
			&p.DNI,
			&p.BirthDate,
			&u.Email,
			&u.Password,
			&u.Active,
			&u.Role,
		)
		if err != nil {
			return nil, err
		}

		u.Person = &p // Asociar la persona al usuario
		users = append(users, u)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return users, nil
}
