package paciente

import (
	"parcial/internal/domain"
)

// Servicio de las funcionalidades de paciente
type IServicePatient interface {
	GetPatientById(id int) (*domain.Patient, error)
	CreatePatient(patient domain.Patient) (*domain.Patient, error)
	UpdatePatient(id int, patient domain.Patient) (*domain.Patient, error)
	DeletePatient(id int) error
}
type Service struct {
	Repository IRepository
}

// Retorna un paciente por ID.
func (ps *Service) GetPatientById(id int) (*domain.Patient, error) {
	patient, err := ps.Repository.GetPatientById(id)
	if err != nil {
		return nil, err
	}
	return patient, nil
}

// Creacion de un nuevo paciente
func (ps *Service) CreatePatient(patient domain.Patient) (*domain.Patient, error) {
	p, err := ps.Repository.CreatePatient(patient)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Elimina un paciente por ID
func (ps *Service) DeletePatient(id int) error {
	err := ps.Repository.DeletePatient(id)
	if err != nil {
		return err
	}
	return nil
}

// Modifica un paciente por ID
func (ps *Service) UpdatePatient(id int, p domain.Patient) (*domain.Patient, error) {
	patient, err := ps.Repository.UpdatePatient(id, p)
	if err != nil {
		return nil, err
	}
	return patient, nil
}
