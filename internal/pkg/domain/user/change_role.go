package domain

import "strings"

func (s *Service) ChangeRole(id string, newRole string) error {

	newRole = strings.TrimSpace(newRole)
	newRole = strings.ToLower(newRole)

	if newRole != "admin" && newRole != "user" {
		return ErrRoleInvalid
	}

	user, err := s.GetByID(id)
	if err != nil {
		return err
	}

	user.Role = newRole

	s.Update(user.Id, user)

	return nil
}
