package controller

import (
	"github.com/sebastiaofortes/solid-go/5_dependency_inversion/violation/service"
)

type EmailController struct {
	service *service.EmailService
	// some email sender
}

func NewEmailController(serv *service.EmailService) *EmailController {
	return &EmailController{
		service: serv,
	}
}

func (s *EmailController) SendEmailToUser(email string, message string) error {
	// validations
	err := s.service.SendEmailToUser(email, message)
	if err != nil {
		return err
	}

	return nil
}
