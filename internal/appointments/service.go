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

type appointmentService struct {
	r IAppointmentRepository
}

func (s *appointmentService) GetAllAppointments() (*[]domain.Appointment, error) {
	appointments, err := s.r.GetAllAppointments()
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (s *appointmentService) GetAppointmentByID(id int) (*domain.Appointment, error) {
	appointment, err := s.r.GetAppointmentById(id)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}

func (s *appointmentService) CreateAppointment(a domain.Appointment) (*domain.Appointment, error) {
	appointment, err := s.r.CreateAppointment(a)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}

func (s *appointmentService) UpdateAppointment(id int, a domain.Appointment) (*domain.Appointment, error) {
	appointment, err := s.r.UpdateAppointment(id, a)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}

func (s *appointmentService) DeleteAppointment(id int) error {
	err := s.r.DeleteAppointment(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *appointmentService) GetAppointmentsByDni(dni string) (*[]domain.Appointment, error) {
	// Llamar al repositorio para obtener las citas por DNI
	appointments, err := s.r.GetAppointmentsByDni(dni)
	if err != nil {
		return nil, err
	}
	return appointments, nil
}
