package services

import "mongo-with-golang/models"

type UserService interface {
	CreateDomain(*models.SetTime) error
	GetDomain(*string) (*models.SetTime, error)
	GetAll() ([]*models.SetTime, error)
}
