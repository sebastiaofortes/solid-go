package controller

type EmailController struct {
	service ServiceInterface
	// some email sender
}

// Abstract Factory Example
func NewEmailController(serv ServiceInterface) *EmailController {
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
