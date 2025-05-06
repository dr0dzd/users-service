package user

import "github.com/google/uuid"

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service{
	return &Service{repo: r}
}

func (s *Service) CreateUser(u User) (User, error) {
	return s.repo.CreateUser(u)
}

func (s *Service) GetUser(userID uuid.UUID) (User, error){
	return s.repo.GetUser(userID)
}

func (s *Service) GetAllUsers() ([]User, error){
	return s.repo.GetAllUsers()
}

func (s *Service) UpdateUser(u User) (User, error){
	return s.repo.UpdateUser(u)
}

func (s *Service) DeleteUser(userID uuid.UUID) error {
	return s.repo.DeleteUser(userID)
}