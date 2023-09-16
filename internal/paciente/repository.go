package paciente

import (
	"errors"
	"parcial/internal/domain"
	"parcial/pkg/store"
)

type IRepository interface {
	CreatePatient(pacient domain.Patient) (*domain.Patient, error)
	UpdatePatient(id int, patient domain.Patient) (*domain.Patient, error)
	GetPatientById(id int) (*domain.Patient, error)
	DeletePatient(id int) error
}

type Repository struct {
	Store store.IStore
}

func (r *Repository) GetPatientById(id int) (*domain.Patient, error) {
	patient, err := r.Store.GetPatientById(id)
	if err != nil {
		return nil, errors.New("el paciente no existe")
	}
	return patient, nil
}

func (r *Repository) CreatePatient(patient domain.Patient) (*domain.Patient, error) {
	existingPatient, _ := r.Store.GetPatientByDni(patient.DNI)
	if existingPatient != nil {
		return nil, errors.New("ya existe un paciente con ese DNI")
	}
	p, err := r.Store.CreatePatient(patient)
	if err != nil {
		return nil, errors.New("error al crear el paciente")
	}
	return p, nil
}

func (r *Repository) DeletePatient(id int) error {
	_, err := r.GetPatientById(id)
	if err != nil {
		return err
	}
	err = r.Store.DeletePatient(id)
	if err != nil {
		return errors.New("error al eliminar el paciente")
	}
	return nil
}

func (r *Repository) UpdatePatient(id int, patient domain.Patient) (*domain.Patient, error) {
	_, err := r.Store.GetPatientById(id)
	if err != nil {
		return nil, errors.New("el paciente no existe")
	}
	p, err := r.Store.UpdatePatient(id, patient)
	if err != nil {
		return nil, errors.New("error al actualizar el paciente")
	}
	return p, nil
}
