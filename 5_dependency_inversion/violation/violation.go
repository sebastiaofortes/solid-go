package dependencyinversion

import (
	"log"

	"gorm.io/gorm"
)

// camada de domínio

type User struct {
	ID uint
	// some fields
}

// camada de infraestrutura

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserByID(id uint) (*User, error) {
	user := User{}
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// camada de aplicação

type EmailService struct {
	repository *UserRepository
	// some email sender
}

func NewEmailService(repository *UserRepository) *EmailService {
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