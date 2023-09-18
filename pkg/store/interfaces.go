package store

import (
	"parcial/internal/domain"
)

type IStorePatient interface {
	GetPatientById(id int) (*domain.Patient, error)
	GetPatientByDni(dni string) (*domain.Patient, error)
	CreatePatient(patient domain.Patient) (*domain.Patient, error)
	UpdatePatient(id int, patient domain.Patient) (*domain.Patient, error)
	DeletePatient(id int) error
}

type IStoreDentist interface {
	GetDentistById(id int) (*domain.Dentist, error)
	GetDentistByRegistration(id string) (*domain.Dentist, error)
	CreateDentist(dentist domain.Dentist) (*domain.Dentist, error)
	UpdateDentist(id int, dentist domain.Dentist) (*domain.Dentist, error)
	DeleteDentist(id int) error
}

type IStoreAppointment interface {
	GetAllAppointments() (*[]domain.Appointment, error)
	GetAppointmentById(id int) (*domain.Appointment, error)
	GetAppointmentsByDni(dni string) (*[]domain.Appointment, error)
	CreateAppointment(appointment domain.Appointment) (*domain.Appointment, error)
	UpdateAppointment(id int, appointment domain.Appointment) (*domain.Appointment, error)
	DeleteAppointment(id int) error
}
