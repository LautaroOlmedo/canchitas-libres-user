package domain

// import (
// 	"canchitas-libres-user/internal/configuration"
// 	domain "canchitas-libres-user/internal/pkg/domain/person"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/mock"
// )

// //var sm Service

// // func TestMain(m *testing.M) {
// // 	storageMock := &MockStorageRepository{}
// // 	config := &configuration.Configuration{}
// // 	sm = *NewService(config, storageMock)

// // 	code := m.Run()
// // 	os.Exit(code)
// // }

// func Test_Add(t *testing.T) {
// 	type testCase struct {
// 		nombreTest    string
// 		user          User
// 		expectedError error
// 	}

// 	dateValid, _ := time.Parse("2006-01-02", "1990-01-05")
// 	dateInvalid, _ := time.Parse("2006-01-02", "2010-01-05")
// 	dateInvalidButSameYear, _ := time.Parse("2006-01-02", "2007-07-05") //Esta fecha tiene que ser el año que cumple 18 pero sin cumplirlos todavía.

// 	tCases := []testCase{
// 		{
// 			nombreTest: "Prueba correcta.",
// 			user: User{
// 				Person: &domain.Person{
// 					ID:        "5",
// 					FirstName: "Nikola",
// 					LastName:  "Jokic",
// 					DNI:       40220220,
// 					BirthDate: dateValid,
// 				},
// 				Id:       "5",
// 				Email:    "elgordojokic@example.com",
// 				Password: "gonzaloquito",
// 				Active:   true,
// 				Role:     "admin",
// 				Phone:    "2615555555",
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			nombreTest: "Error. Contraseña muy corta",
// 			user: User{
// 				Person: &domain.Person{
// 					ID:        "5",
// 					FirstName: "Nikola",
// 					LastName:  "Jokic",
// 					DNI:       40220220,
// 					BirthDate: dateValid,
// 				},
// 				Id:       "5",
// 				Email:    "elgordojokic@example.com",
// 				Password: "gonz",
// 				Active:   true,
// 				Role:     "admin",
// 				Phone:    "2615555555",
// 			},
// 			expectedError: ErrPasswordMinCharacters,
// 		},
// 		{
// 			nombreTest: "Error. Rol inexistente. ",
// 			user: User{
// 				Person: &domain.Person{
// 					ID:        "5",
// 					FirstName: "Nikola",
// 					LastName:  "Jokic",
// 					DNI:       40220220,
// 					BirthDate: dateValid,
// 				},
// 				Id:       "5",
// 				Email:    "elgordojokic@example.com",
// 				Password: "gonzaloquito",
// 				Active:   true,
// 				Role:     "ad",
// 				Phone:    "2615555555",
// 			},
// 			expectedError: ErrRoleInvalid,
// 		},
// 		{
// 			nombreTest: "Error. Menor de edad.",
// 			user: User{
// 				Person: &domain.Person{
// 					ID:        "5",
// 					FirstName: "Nikola",
// 					LastName:  "Jokic",
// 					DNI:       40220220,
// 					BirthDate: dateInvalid,
// 				},
// 				Id:       "5",
// 				Email:    "elgordojokic@example.com",
// 				Password: "gonzaloquito",
// 				Active:   true,
// 				Role:     "admin",
// 				Phone:    "2615555555",
// 			},
// 			expectedError: ErrMinAge,
// 		},
// 		{
// 			nombreTest: "Error. Menor de edad por unos meses en el mismo año.",
// 			user: User{
// 				Person: &domain.Person{
// 					ID:        "5",
// 					FirstName: "Nikola",
// 					LastName:  "Jokic",
// 					DNI:       40220220,
// 					BirthDate: dateInvalidButSameYear,
// 				},
// 				Id:       "5",
// 				Email:    "elgordojokic@example.com",
// 				Password: "gonzaloquito",
// 				Active:   true,
// 				Role:     "admin",
// 				Phone:    "2615555555",
// 			},
// 			expectedError: ErrMinAge,
// 		},
// 	}

// 	//strorageMock := &MockStorageRepository{}
// 	config := &configuration.Configuration{}

// 	//var sm = NewService(config, strorageMock)

// 	for i := range tCases {
// 		tc := tCases[i]

// 		t.Run(tc.nombreTest, func(t *testing.T) {

// 			t.Parallel()

// 			mockStorage := new(MockStorageRepository) // mock nuevo por cada subtest, si no en la ejecucion paralela al todos consumir el mismo mock, falla.
// 			sm := NewService(config, mockStorage)     // implementamos cada mock al service.

// 			if tc.expectedError == nil {
// 				mockStorage.On("Add", mock.Anything, tc.user).Return(nil).Once()
// 			} //Aca le decimos que solo ejecute el add del mock en caso de superar las otras validaciones y que lo haga con un ctx cualquiera, ese user y que retorne nil.

// 			err := sm.Add(tc.user)
// 			if err != tc.expectedError {
// 				t.Fatalf("Yo esperaba el error: %v, y obtuve el error: %v", tc.expectedError, err)
// 			}
// 			mockStorage.AssertExpectations(t)
// 		})
// 	}
// }
