package domain

func (s *Service) GetAll() error {
	return s.StorageRepository.GetAll()
}
