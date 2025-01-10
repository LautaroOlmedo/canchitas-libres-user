package domain

func (s *Service) GetByID(id int) (Field, error) {
	return s.StorageRepository.GetByID(id)
}
