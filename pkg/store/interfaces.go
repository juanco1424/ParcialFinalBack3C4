package store

import (
	"parcial/internal/domain"
)

type IStoreDentist interface {
	GetPatientById(id string) (*domain.Dentist, error)
	CreatePatient(patient domain.Dentist) (*domain.Dentist, error)
	UpdatePatient(id string, patient domain.Dentist) (*domain.Dentist, error)
	DeletePatient(id string) error
}
