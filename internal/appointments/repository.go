package appointments

import (
	"errors"
	"parcial/internal/domain"
	"parcial/pkg/store"
)

type IAppointmentRepository interface {
	GetAllAppointments() (*[]domain.Appointment, error)
	GetAppointmentById(id int) (*domain.Appointment, error)
	GetAppointmentsByDni(dni string) (*[]domain.Appointment, error)
	CreateAppointment(appointment domain.Appointment) (*domain.Appointment, error)
	UpdateAppointment(id int, appointment domain.Appointment) (*domain.Appointment, error)
	DeleteAppointment(id int) error
}

type AppointmentRepository struct {
	Store store.IStoreAppointment
}

func (r *AppointmentRepository) GetAllAppointments() (*[]domain.Appointment, error) {
	appointments, err := r.Store.GetAllAppointments()
	if err != nil {
		return nil, errors.New("No se puede obtener el listado de turnos")
	}
	return appointments, nil
}

func (r *AppointmentRepository) GetAppointmentById(id int) (*domain.Appointment, error) {
	appointment, err := r.Store.GetAppointmentById(id)
	if err != nil {
		return nil, errors.New("No se puede obtener el turno")
	}
	return appointment, nil
}

func (r *AppointmentRepository) GetAppointmentsByDni(dni string) (*[]domain.Appointment, error) {
	appointments, err := r.Store.GetAppointmentsByDni(dni)
	if err != nil {
		return nil, errors.New("No se pueden obtener los turnos del paciente")
	}
	return appointments, nil
}

func (r *AppointmentRepository) CreateAppointment(appointment domain.Appointment) (*domain.Appointment, error) {
	a, err := r.Store.CreateAppointment(appointment)
	if err != nil {
		return nil, errors.New("Error al crear el turno")
	}
	return a, nil
}

func (r *AppointmentRepository) UpdateAppointment(id int, appointment domain.Appointment) (*domain.Appointment, error) {
	_, err := r.Store.GetAppointmentById(id)
	if err != nil {
		return nil, errors.New("El turno no existe")
	}
	a, err := r.Store.UpdateAppointment(id, appointment)
	if err != nil {
		return nil, errors.New("No se pudo actualizar el turno")
	}
	return a, nil
}

func (r *AppointmentRepository) DeleteAppointment(id int) error {
	_, err := r.Store.GetAppointmentById(id)
	if err != nil {
		return err
	}
	err = r.Store.DeleteAppointment(id)
	if err != nil {
		return errors.New("No se pudo eliminar el turno")
	}
	return nil
}
