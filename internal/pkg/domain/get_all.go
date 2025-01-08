package domain

func (s *Service) GetAll() ([]Field, error) {
	return s.StorageRepository.GetAll()
}
