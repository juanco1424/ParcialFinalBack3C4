package database

import (
	"database/sql"
	"parcial/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func (s *SqlStore) GetPatientById(id int) (*domain.Patient, error) {
	var patient domain.Patient
	query := "SELECT * FROM patient WHERE ID = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&patient.Id, &patient.Name, &patient.LastName, &patient.Address, &patient.DNI, &patient.DischargeDate)
	if err != nil {
		return nil, err
	}
	return &patient, nil
}
func (s *SqlStore) GetPatientByDni(dni string) (*domain.Patient, error) {
	var patient domain.Patient
	query := "SELECT * FROM patient WHERE DNI = ?;"
	row := s.DB.QueryRow(query, dni)
	err := row.Scan(&patient.Id, &patient.Name, &patient.LastName, &patient.Address, &patient.DNI, &patient.DischargeDate)
	if err != nil {
		return nil, err
	}
	println(&patient)
	return &patient, nil
}

func (s *SqlStore) CreatePatient(patient domain.Patient) (*domain.Patient, error) {
	query := "INSERT INTO patient (Name, LastName, Address, DNI, DischargeDate) VALUES (?, ?, ?, ?, ?)"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(patient.Name, patient.LastName, patient.Address, patient.DNI, patient.DischargeDate)
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &patient, nil
}

func (s *SqlStore) DeletePatient(id int) error {
	query := "DELETE FROM patient WHERE ID = ?"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) UpdatePatient(id int, patient domain.Patient) (*domain.Patient, error) {
	updateQuery := "UPDATE patient SET Name = ?, LastName = ?, Address = ?, DischargeDate = ? WHERE ID = ?"
	_, err := s.DB.Exec(updateQuery, patient.Name, patient.LastName, patient.Address, patient.DischargeDate, id)
	if err != nil {
		return nil, err
	}
	return &patient, nil
}
