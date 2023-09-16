package odontologo

import "parcial/internal/domain"

type IService interface {
	CreateDentist(dentist domain.Dentist) (*domain.Dentist, error)
	GetDentistById(id int) (*domain.Dentist, error)
	UpdateDentist(id int, patient domain.Dentist) (*domain.Dentist, error)
	DeleteDentist(id int) error
}

type Service struct {
	Repository IRepository
}

func (ps *Service) CreateDentist(dentist domain.Dentist) (*domain.Dentist, error) {
	p, err := ps.Repository.CreateDentist(dentist)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (ps *Service) GetDentistById(id int) (*domain.Dentist, error) {
	patient, err := ps.Repository.GetDentistById(id)
	if err != nil {
		return nil, err
	}
	return patient, nil
}

func (ps *Service) UpdatePatient(id int, d domain.Dentist) (*domain.Dentist, error) {
	patient, err := ps.Repository.UpdateDentist(id, d)
	if err != nil {
		return nil, err
	}
	return patient, nil
}

func (ps *Service) DeleteDentist(id int) error {
	err := ps.Repository.DeleteDentist(id)
	if err != nil {
		return err
	}
	return nil
}
