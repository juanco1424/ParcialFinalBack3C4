package odontologo

import (
	"errors"
	"parcial/internal/domain"
	"parcial/pkg/store"
)

type IRepository interface {
	CreateDentist(dentist domain.Dentist) (*domain.Dentist, error)
	GetDentistById(id int) (*domain.Dentist, error)
	UpdateDentist(id int, dentist domain.Dentist) (*domain.Dentist, error)
	DeleteDentist(id int) error
}

type Repository struct {
	Store store.IStoreDentist
}

func (r *Repository) CreateDentist(dentist domain.Dentist) (*domain.Dentist, error) {
	existingDentist, _ := r.Store.GetDentistByRegistration(dentist.Registration)
	if existingDentist != nil {
		return nil, errors.New("ya existe un paciente con ese REGISTRO")
	}
	p, err := r.Store.CreateDentist(dentist)
	if err != nil {
		return nil, errors.New("error al crear el Odontologo")
	}
	return p, nil
}

func (r *Repository) GetDentistById(id int) (*domain.Dentist, error) {
	dentist, err := r.Store.GetDentistById(id)
	if err != nil {
		return nil, errors.New("el odontolgo no existe")
	}
	return dentist, nil
}

func (r *Repository) UpdateDentist(id int, dentist domain.Dentist) (*domain.Dentist, error) {
	_, err := r.Store.GetDentistById(id)
	if err != nil {
		return nil, errors.New("el Odontologo no existe")
	}
	p, err := r.Store.UpdateDentist(id, dentist)
	if err != nil {
		return nil, errors.New("error al actualizar el Odontologo")
	}
	return p, nil
}

func (r *Repository) DeleteDentist(id int) error {
	_, err := r.GetDentistById(id)
	if err != nil {
		return err
	}
	err = r.Store.DeleteDentist(id)
	if err != nil {
		return errors.New("error al eliminar el odontologo")
	}
	return nil
}
