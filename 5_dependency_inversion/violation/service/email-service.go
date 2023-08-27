package service

import (
	"github.com/sebastiaofortes/solid-go/5_dependency_inversion/violation/repository"
)

type EmailService struct {
	repository repository.UserRepository
}

func (s *EmailService) SendEmailToUser(email string, message string) error {
	err := s.repository.SendEmail(email, message)
	if err != nil {
		return err
	}
	return nil
}

func NewEmailService(r *repository.UserRepository) *EmailService{
	return &EmailService{
		repository: *r,
	}
}