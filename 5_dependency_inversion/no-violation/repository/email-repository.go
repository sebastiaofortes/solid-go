package repository

type EmailRepository struct {
}

func NewEmailRepository() *EmailRepository {
	return &EmailRepository{}
}

func (r *EmailRepository) SendEmail(email string, message string) error {
	// send email to client...

	return nil
}
