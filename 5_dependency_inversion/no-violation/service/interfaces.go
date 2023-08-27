package service

type RepositoryInterface interface{
	SendEmail(email string, message string) error
}