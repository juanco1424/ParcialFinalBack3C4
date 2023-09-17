package database

import (
	"database/sql"
	"fmt"
	"parcial/internal/domain"
)

type SqlStorePatient struct {
	DB *sql.DB
}

func (s *SqlStorePatient) CreatePatient(patient domain.Patient) (*domain.Patient, error) {
	query := "INSERT INTO patient (id, name, last_name, address, dni, discharge_date) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(patient.Id, patient.Name, patient.LastName, patient.Address, patient.DNI, patient.DischargeDate)
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &patient, nil
}

func (s *SqlStorePatient) GetPatientById(id int) (*domain.Patient, error) {
	var patient domain.Patient
	query := "SELECT * FROM patient WHERE id = ?;"
	row := s.DB.QueryRow(query, id)

	err := row.Scan(&patient.Id, &patient.Name, &patient.LastName, &patient.Address, &patient.DNI, &patient.DischargeDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no se encontró ningún paciente con el ID proporcionado")
		}
		return nil, err
	}

	return &patient, nil
}

func (s *SqlStorePatient) UpdatePatient(id int, patient domain.Patient) (*domain.Patient, error) {
	updateQuery := "UPDATE patient SET name = ?, last_name = ?, address = ?, dni = ?, discharge_date = ? WHERE id = ?"
	result, err := s.DB.Exec(updateQuery, patient.Name, patient.LastName, patient.Address, patient.DNI, patient.DischargeDate, id)
	if err != nil {
		return nil, fmt.Errorf("error al actualizar el paciente con ID %d: %v", id, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no se encontró ningún paciente con el ID %d para actualizar", id)
	}

	return &patient, nil
}

func (s *SqlStorePatient) DeletePatient(id int) error {
	query := "DELETE FROM patient WHERE id = ?"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStorePatient) GetPatientByDni(dni string) (*domain.Patient, error) {
	var patient domain.Patient
	query := "SELECT * FROM patient WHERE dni = ?;"
	row := s.DB.QueryRow(query, dni)

	err := row.Scan(&patient.Id, &patient.Name, &patient.LastName, &patient.Address, &patient.DNI, &patient.DischargeDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no se encontró ningún paciente con el DNI proporcionado")
		}

		return nil, fmt.Errorf("error al buscar paciente por DNI: %v", err)
	}

	return &patient, nil
}
