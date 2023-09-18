package database

import (
	"database/sql"
	"fmt"
	"parcial/internal/domain"
)

type SqlStoreAppointment struct {
	DB *sql.DB
}

func (s *SqlStoreAppointment) GetAllAppointments() (*[]domain.Appointment, error) {
	query := "SELECT * FROM appointment"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	appointments := make([]domain.Appointment, 0)

	for rows.Next() {
		var appointment domain.Appointment
		err := rows.Scan(&appointment.ID, &appointment.Patient.Id, &appointment.Dentist.ID, &appointment.Date, &appointment.Hour, &appointment.Description)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &appointments, nil 
}


func (s *SqlStoreAppointment) GetAppointmentById(id int) (*domain.Appointment, error) {
	query := "SELECT * FROM appointment WHERE ID = ?"
	row := s.DB.QueryRow(query, id)

	var appointment domain.Appointment
	err := row.Scan(&appointment.ID, &appointment.Patient.Id, &appointment.Dentist.ID, &appointment.Date, &appointment.Hour, &appointment.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no se encontró ningún turno con el ID proporcionado")
		}
		return nil, err
	}

	return &appointment, nil
}

func (s *SqlStoreAppointment) GetAppointmentsByDni(dni string) (*[]domain.Appointment, error) {
	query := "SELECT a.* FROM appointment a JOIN patient p ON a.patient_id = p.id WHERE p.dni = ?"
	rows, err := s.DB.Query(query, dni)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	appointments := make([]domain.Appointment, 0) 

	for rows.Next() {
		var appointment domain.Appointment
		err := rows.Scan(&appointment.ID, &appointment.Patient.Id, &appointment.Dentist.ID, &appointment.Date, &appointment.Hour, &appointment.Description)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &appointments, nil 
}


func (s *SqlStoreAppointment) CreateAppointment(appointment domain.Appointment) (*domain.Appointment, error) {
	query := "INSERT INTO appointment (patient_id, dentist_id, date, hour, description) VALUES (?, ?, ?, ?, ?)"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(appointment.Patient.Id, appointment.Dentist.ID, appointment.Date, appointment.Hour, appointment.Description)
	if err != nil {
		return nil, err
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	appointment.ID = int(insertedID)
	return &appointment, nil
}

func (s *SqlStoreAppointment) UpdateAppointment(id int, appointment domain.Appointment) (*domain.Appointment, error) {
	query := "UPDATE appointment SET patient_id = ?, dentist_id = ?, date = ?, hour = ?, description = ? WHERE ID = ?"
	result, err := s.DB.Exec(query, appointment.Patient.Id, appointment.Dentist.ID, appointment.Date, appointment.Hour, appointment.Description, id)
	if err != nil {
		return nil, fmt.Errorf("error al actualizar el turno con ID %d: %v", id, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no se encontró ningún turno con el ID %d para actualizar", id)
	}

	return &appointment, nil
}

func (s *SqlStoreAppointment) DeleteAppointment(id int) error {
	query := "DELETE FROM appointment WHERE ID = ?"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
