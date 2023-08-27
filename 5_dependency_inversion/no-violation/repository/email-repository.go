package repository

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) SendEmail(email string, message string) error {
	// send email to client...

	return nil
}
