package storage

import (
	domain2 "canchitas-libres-user/internal/pkg/domain/person"
	domain "canchitas-libres-user/internal/pkg/domain/user"
	"context"
	"fmt"
	//"time"
)

const (
	queryInsertPerson = `
        INSERT INTO persons (firstname, lastname, dni, birthdate)
        VALUES ($1, $2, $3, $4)
		RETURNING       id;`
	queryInsertUser = `
        INSERT INTO users (id, email, password, active, role)
        VALUES ($1, $2, $3, $4, $5);`
	querySelectAllUsers = `
    SELECT 
    u.id AS user_id,
    u.email AS email,
    u.password AS password,
    u.active AS active,
    u.role AS role,
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
		u.role
	FROM persons p
	JOIN users u ON p.id = u.id
	WHERE u.id = $1;
`
	queryDeleteUser   = `DELETE FROM users WHERE id = $1;`
	queryDeletePerson = `DELETE FROM person WHERE id = $1;`
)

// type UserAndPerson struct {
// 	ID        int       `db:"id"`
// 	FirstName string    `db:"firstname"`
// 	LastName  string    `db:"lastname"`
// 	DNI       int       `db:"dni"`
// 	BirthDate time.Time `db:"birthdate"`
// 	Id        int       `db:"user_id"`
// 	Email     string    `db:"email"`
// 	Password  string    `db:"password"`
// 	Active    bool      `db:"active"`
// 	Role      string    `db:"role"`
// }

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
		// user := domain.User{
		// 	Id:       up.Id,
		// 	Email:    up.Email,
		// 	Password: up.Password,
		// 	Active:   up.Active,
		// 	Role:     up.Role,
		// 	Person: &domain2.Person{
		// 		ID:        up.ID,
		// 		FirstName: up.FirstName,
		// 		LastName:  up.LastName,
		// 		DNI:       up.DNI,
		// 		BirthDate: up.BirthDate,
		// 	},
		// }
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
	_, err = tx.ExecContext(ctx, queryInsertUser, user.Id, user.Email, user.Password, user.Active, user.Role)
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

	//user.Person.FirstName = u_p.FirstName
	// user := domain.User{
	// 	// Person: &domain2.Person{
	// 	// 	ID:        u_p.ID,
	// 	// 	FirstName: u_p.FirstName,
	// 	// 	LastName:  u_p.LastName,
	// 	// 	DNI:       u_p.DNI,
	// 	// 	BirthDate: u_p.BirthDate,
	// 	// },
	// 	Person:   &u_p.Person,
	// 	Id:       u_p.Id,
	// 	Email:    u_p.Email,
	// 	Password: u_p.Password,
	// 	Active:   u_p.Active,
	// 	Role:     u_p.Role,
	// }

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
