package service

type EmailService struct {
	repository RepositoryInterface
}

func (s *EmailService) SendEmailToUser(email string, message string) error {
	err := s.repository.SendEmail(email, message)
	if err != nil {
		return err
	}
	return nil
}

// Abstract Factory Example
func NewEmailService(r RepositoryInterface) *EmailService {
	return &EmailService{
		repository: r,
	}
}
