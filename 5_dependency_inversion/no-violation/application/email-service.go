package application

import (
	"log"

	"github.com/sebastiaofortes/solid-go/5_dependency_inversion/no-violation/domain"
)

type UserRepositoryInterface interface {
	GetUserByID(id uint) (*domain.User, error)
}

type EmailService struct {
	repository UserRepositoryInterface
	// some email sender
}

func NewEmailService(repository UserRepositoryInterface) *EmailService {
	return &EmailService{
		repository: repository,
	}
}

func (s *EmailService) SendRegistrationEmail(userID uint) error {
	user, err := s.repository.GetUserByID(userID)
	if err != nil {
		return err
	}

	log.Println(user)
	// send email
	return nil
}
