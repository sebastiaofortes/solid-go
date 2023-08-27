package service

import (
	"github.com/sebastiaofortes/solid-go/5_dependency_inversion/violation/repository"
)

type EmailService struct {
	repository repository.EmailRepository
}

func (s *EmailService) SendEmailToUser(email string, message string) error {
	err := s.repository.SendEmail(email, message)
	if err != nil {
		return err
	}
	return nil
}

func NewEmailService(r *repository.EmailRepository) *EmailService{
	return &EmailService{
		repository: *r,
	}
}