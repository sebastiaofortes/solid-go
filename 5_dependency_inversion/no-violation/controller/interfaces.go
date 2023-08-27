package controller

type ServiceInterface interface{
	SendEmailToUser(email string, message string) error
}