package store

import (
	"parcial/internal/domain"
)


type IStore interface {
	GetPatientById(id int) (*domain.Patient, error)
	GetPatientByDni(dni string) (*domain.Patient, error)
	CreatePatient(patient domain.Patient) (*domain.Patient, error)
	UpdatePatient(id int, patient domain.Patient) (*domain.Patient, error)
	DeletePatient(id int) error

type IStoreDentist interface {
	GetDentistById(id int) (*domain.Dentist, error)
	GetDentistByRegistration(id string) (*domain.Dentist, error)
	CreateDentist(dentist domain.Dentist) (*domain.Dentist, error)
	UpdateDentist(id int, dentist domain.Dentist) (*domain.Dentist, error)
	DeleteDentist(id int) error

}
