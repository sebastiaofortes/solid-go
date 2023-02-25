package dependencyinversion

import (
	"testing"
	// some dependencies
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestUserRepository_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	dialector := mysql.New(mysql.Config{
		DSN:        "dummy",
		DriverName: "mysql",
		Conn:       db,
	})

	finalDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)
	
	repository := NewUserRepository(finalDB)

	mock.ExpectQuery("mocked sql querie")

	r, err := repository.GetUserByID(11)

	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestEmailService_SendRegistrationEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	dialector := mysql.New(mysql.Config{
		DSN:        "dummy",
		DriverName: "mysql",
		Conn:       db,
	})

	finalDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	repository := NewUserRepository(finalDB)
	service := NewEmailService(repository)

	mock.ExpectQuery("mocked sql querie")

	r := service.SendEmailToUser(13)

	assert.NoError(t, r)
}
