package infra

import (
	"github.com/sebastiaofortes/solid-go/5_dependency_inversion/no-violation/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserByID(id uint) (*domain.User, error) {
	user := domain.User{}
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
