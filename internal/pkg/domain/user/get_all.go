package domain

func (s *Service) GetAll() ([]User, error) {
	return s.StorageRepository.GetAll()
}
