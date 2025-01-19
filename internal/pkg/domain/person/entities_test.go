package domain

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func Test_NewPerson(t *testing.T) {
	type testCase struct {
		test          string
		FirstName     string
		LastName      string
		DNI           int
		BirthDate     time.Time
		expectedError error
	}
	layout := "2006-01-02"
	dateString := "1990-01-05"
	fecha, _ := time.Parse(layout, dateString)
	testCases := []testCase{
		{
			test:          "Prueba correcta. Ingreso datos validos",
			FirstName:     "Pepe",
			LastName:      "Datos",
			BirthDate:     fecha,
			DNI:           112312,
			expectedError: nil,
		},
		{
			test:          "Error. No ingreso el nombre",
			FirstName:     "",
			LastName:      "Datos",
			BirthDate:     time.Now(),
			DNI:           112312,
			expectedError: ErrMissingParameter,
		},
		{
			test:          "Error. No ingreso el DNI",
			FirstName:     "Pepe",
			LastName:      "Datos",
			BirthDate:     time.Now(),
			DNI:           0,
			expectedError: ErrMissingParameter,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.test, func(t *testing.T) {
			t.Parallel()
			p, err := NewPerson(tc.FirstName, tc.LastName, tc.DNI, tc.BirthDate)
			if !errors.Is(err, tc.expectedError) {
				fmt.Println("person: ", p)
				t.Fatalf("Yo esperaba el error: %v, y obtuve el error: %v", tc.expectedError, err)
			}
		})
	}
}
