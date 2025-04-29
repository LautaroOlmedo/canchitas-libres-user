package storage

import (
	domain2 "canchitas-libres-user/internal/pkg/domain/person"
	domain "canchitas-libres-user/internal/pkg/domain/user"
	"context"
	"fmt"
)

const (
	queryInsertPerson = `
        INSERT INTO persons (firstname, lastname, dni, birthdate)
        VALUES ($1, $2, $3, $4)
		RETURNING       id;`
	queryInsertUser = `
        INSERT INTO users (id, email, password, active, role, phone)
        VALUES ($1, $2, $3, $4, $5, $6);`
	querySelectAllUsers = `
    SELECT 
		u.id AS user_id,
		u.email AS email,
		u.password AS password,
		u.active AS active,
		u.role AS role,
		u.phone AS phone,
		p.id AS id,
		p.firstname AS firstname,
		p.lastname AS lastname,
		p.dni AS dni,
		p.birthdate AS birthdate
	FROM 
		users u
	JOIN 
		persons p
	ON 
		u.id = p.id;`
	querySelectUserByID = `
	SELECT 
		p.id AS id, 
		p.firstname, 
		p.lastname, 
		p.dni, 
		p.birthdate, 
		u.id AS user_id, 
		u.email, 
		u.password, 
		u.active, 
		u.role,
		u.phone
	FROM persons p
	JOIN users u ON p.id = u.id
	WHERE u.id = $1;
`
	queryDeleteUser      = `DELETE FROM users WHERE id = $1;`
	queryDeletePerson    = `DELETE FROM persons WHERE id = $1;`
	queryUpdateFirstname = `UPDATE persons SET firstname = $1 WHERE id = $2`
	queryUpdateLastname  = `UPDATE persons SET lastname = $1 WHERE id = $2`
	queryUpdateDni       = `UPDATE persons SET dni = $1 WHERE id = $2`
	queryUpdateBirthdate = `UPDATE persons SET birthdate = $1 WHERE id = $2`
	queryUpdateEmail     = `UPDATE users SET email = $1 WHERE id = $2`
	queryUpdatePassword  = `UPDATE users SET password = $1 WHERE id = $2`
	queryUpdateActive    = `UPDATE users SET active = $1 WHERE id = $2`
	queryUpdateRole      = `UPDATE users SET role = $1 WHERE id = $2`
	queryUpdatePhone     = `UPDATE users SET phone = $1 WHERE id = $2`
)

type UserAndPerson struct {
	//Si lo declaro de esta manera no funciona:
	// Person domain2.Person
	// User   domain.User
	// Pero de esta manera si:
	domain2.Person
	domain.User
}

func (p *Postgres) GetAll() ([]domain.User, error) {
	var u_p []UserAndPerson
	var users []domain.User

	err := p.Select(&u_p, querySelectAllUsers)
	if err != nil {
		return nil, err
	}

	for _, up := range u_p {
		user := up.User
		user.Person = &up.Person
		users = append(users, user)
	}
	return users, nil
}

func (p *Postgres) Add(ctx context.Context, user domain.User) error {
	fmt.Println("in infrastructure layer we have a field whit name: ", user.Person.FirstName)

	tx, err := p.Begin()
	if err != nil {
		return err
	}

	var personID int
	err = tx.QueryRowContext(ctx, queryInsertPerson, user.Person.FirstName, user.Person.LastName, user.Person.DNI, user.Person.BirthDate).Scan(&personID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert person: %w", err)
	}
	user.Id = personID
	// Ensure user.Id is the same as the inserted person ID
	_, err = tx.ExecContext(ctx, queryInsertUser, user.Id, user.Email, user.Password, user.Active, user.Role, user.Phone)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert user: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetByID(id int) (domain.User, error) {
	var user domain.User
	var u_p UserAndPerson

	err := p.Get(&u_p, querySelectUserByID, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user by ID %d: %w", id, err)
	}
	user = u_p.User
	user.Person = &u_p.Person

	return user, nil
}

func (p *Postgres) Delete(ctx context.Context, id int) error {
	tx, err := p.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	_, err = tx.ExecContext(ctx, queryDeleteUser, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete user: %w", err)
	} //Importante eliminar primero el user porque este depende del person.

	_, err = tx.ExecContext(ctx, queryDeletePerson, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete person: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func (p *Postgres) Update(ctx context.Context, id int, userU domain.User) error {
	tx, err := p.Begin()
	if err != nil {
		return err
	}

	if userU.Person.FirstName != "" {
		_, err = tx.ExecContext(ctx, queryUpdateFirstname, userU.Person.FirstName, id)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update user firstname: %w", err)
		}
	}
	if userU.Person.LastName != "" {
		_, err = tx.ExecContext(ctx, queryUpdateLastname, userU.Person.LastName, id)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update user lastname: %w", err)
		}
	}
	if userU.Person.DNI != 0 {
		_, err = tx.ExecContext(ctx, queryUpdateDni, userU.Person.DNI, id)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update user dni: %w", err)
		}
	}
	if userU.Person.BirthDate.Format("2006-01-02") != "0001-01-01" {
		_, err = tx.ExecContext(ctx, queryUpdateBirthdate, userU.Person.BirthDate, id)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update user birthdate: %w", err)
		}
	}
	if userU.Email != "" {
		_, err = tx.ExecContext(ctx, queryUpdateEmail, userU.Email, id)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update user email: %w", err)
		}
	}
	if userU.Password != "" {
		_, err = tx.ExecContext(ctx, queryUpdatePassword, userU.Password, id)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update user password: %w", err)
		}
	}
	if userU.Role != "" {
		_, err = tx.ExecContext(ctx, queryUpdateRole, userU.Role, id)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update user role: %w", err)
		}
	}
	if userU.Phone != "" {
		_, err = tx.ExecContext(ctx, queryUpdatePhone, userU.Phone, id)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update user phone")
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
