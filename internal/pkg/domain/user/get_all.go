package domain

// type UserResponse struct {
// 	PersonID  int       `db:"person_id"`
// 	FirstName string    `db:"first_name"`
// 	LastName  string    `db:"last_name"`
// 	DNI       int       `db:"dni"`
// 	BirthDate time.Time `db:"birth_date"`
// 	Id        int       `db:"user_id"`
// 	Email     string    `db:"email"`
// 	Password  string    `db:"password"`
// 	Active    bool      `db:"active"`
// 	Role      string    `db:"role"`
// } //Esto donde va?

func (s *Service) GetAll() ([]User, error) {
	return s.StorageRepository.GetAll()
}
