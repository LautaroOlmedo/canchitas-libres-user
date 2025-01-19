package domain

import (
	domain "canchitas-libres-user/internal/pkg/domain/person"
	"errors"
	"fmt"
	"testing"
)

func Test_User(t *testing.T) {
	type testCase struct {
		nombreTest string
		person     domain.Person //Esto lo tengo que arreglar
		// firstName     string
		// lastName      string
		// dni           int
		// birthDate     time.Time
		id            string
		email         string
		password      string
		active        bool
		role          string
		expectedError error
	}

	usersTest := []testCase{
		{
			nombreTest:    "Prueba correcta.",
			id:            "456",
			email:         "jtomaspeiretti@gmail.com",
			password:      "gonzaloquito",
			active:        true,
			role:          "admin",
			expectedError: nil,
		},
		{
			nombreTest:    "Falta email.",
			id:            "456",
			email:         "",
			password:      "gonzaloquito",
			active:        true,
			role:          "admin",
			expectedError: ErrMissingParameter,
		},
		{
			nombreTest:    "error en el ID.",
			id:            "",
			email:         "jtomaspeiretti@gmail.com",
			password:      "gonzaloquito",
			active:        true,
			role:          "admin",
			expectedError: ErrMissingParameter,
		},
		{
			nombreTest:    "Falta contraseña.",
			id:            "456",
			email:         "jtomaspeiretti@gmail.com",
			password:      "",
			active:        true,
			role:          "admin",
			expectedError: ErrMissingParameter,
		},
		{
			nombreTest:    "Rol vacio.",
			id:            "456",
			email:         "jtomaspeiretti@gmail.com",
			password:      "gonzaloquito",
			active:        true,
			role:          "",
			expectedError: ErrMissingParameter,
		},
		{
			nombreTest:    "Rol inexistente.",
			id:            "456",
			email:         "jtomaspeiretti@gmail.com",
			password:      "gonzaloquito",
			active:        true,
			role:          "adminnn",
			expectedError: ErrRoleInvalid,
		},
		{
			nombreTest:    "Contraseña tiene menos de 5 caracteres.",
			id:            "456",
			email:         "jtomaspeiretti@gmail.com",
			password:      "gonz",
			active:        true,
			role:          "admin",
			expectedError: ErrPasswordMinCharacters,
		},
	}

	for i := range usersTest {
		test := usersTest[i]
		t.Run(test.nombreTest, func(t *testing.T) {
			t.Parallel()
			testUser, err := NewUser(test.person.FirstName, test.person.LastName, test.person.DNI, test.person.BirthDate, test.email, test.password, test.role)
			if !errors.Is(err, test.expectedError) {
				fmt.Println("User testeado: ", testUser)
				t.Fatalf("Yo esperaba el error: %v, y obtuve el error: %v", test.expectedError, err)
			}
		})
	}
}
