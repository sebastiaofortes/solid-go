package application

import (
	"log"

	"github.com/sebastiaofortes/solid-go/5_dependency_inversion/violation/infra"
)

type EmailService struct {
	repository *infra.UserRepository
	// some email sender
}

func NewEmailService(repository *infra.UserRepository) *EmailService {
	return &EmailService{
		repository: repository,
	}
}

func (s *EmailService) SendEmailToUser(userID uint) error {
	user, err := s.repository.GetUserByID(userID)
	if err != nil {
		return err
	}

	log.Println(user)
	// send email
	return nil
}
