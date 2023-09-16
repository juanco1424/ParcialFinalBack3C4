package store

import (
	"parcial/internal/domain"
)

type IStoreDentist interface {
	GetDentistById(id int) (*domain.Dentist, error)
	GetDentistByRegistration(id string) (*domain.Dentist, error)
	CreateDentist(dentist domain.Dentist) (*domain.Dentist, error)
	UpdateDentist(id int, dentist domain.Dentist) (*domain.Dentist, error)
	DeleteDentist(id int) error
}
