package appointments

import (
	"parcial/internal/domain"
)

type IAppointmentService interface {
	GetAllAppointments() (*[]domain.Appointment, error)
	GetAppointmentById(id int) (*domain.Appointment, error)
	GetAppointmentsByDni(dni string) (*[]domain.Appointment, error)
	CreateAppointment(a domain.Appointment) (*domain.Appointment, error)
	UpdateAppointment(id int, a domain.Appointment) (*domain.Appointment, error)
	DeleteAppointment(id int) error
}

type AppointmentService struct {
	Repository IAppointmentRepository
}

func (s *AppointmentService) GetAllAppointments() (*[]domain.Appointment, error) {
	appointments, err := s.Repository.GetAllAppointments()
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (s *AppointmentService) GetAppointmentById(id int) (*domain.Appointment, error) {
	appointment, err := s.Repository.GetAppointmentById(id)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}

func (s *AppointmentService) CreateAppointment(a domain.Appointment) (*domain.Appointment, error) {
	appointment, err := s.Repository.CreateAppointment(a)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}

func (s *AppointmentService) UpdateAppointment(id int, a domain.Appointment) (*domain.Appointment, error) {
	appointment, err := s.Repository.UpdateAppointment(id, a)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}

func (s *AppointmentService) DeleteAppointment(id int) error {
	err := s.Repository.DeleteAppointment(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *AppointmentService) GetAppointmentsByDni(dni string) (*[]domain.Appointment, error) {
	// Llamar al repositorio para obtener las citas por DNI
	appointments, err := s.Repository.GetAppointmentsByDni(dni)
	if err != nil {
		return nil, err
	}
	return appointments, nil
}
