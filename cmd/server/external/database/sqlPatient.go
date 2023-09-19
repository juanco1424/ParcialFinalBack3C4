package database

import (
	"database/sql"
	"parcial/internal/domain"
)

type SqlStorePatient struct {
	DB *sql.DB
}

func (s *SqlStorePatient) GetPatientById(id int) (*domain.Patient, error) {
	var patient domain.Patient
	query := "SELECT * FROM patient WHERE ID = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&patient.Id, &patient.Name, &patient.LastName, &patient.Address, &patient.DNI, &patient.DischargeDate)
	if err != nil {
		return nil, err
	}
	return &patient, nil
}
func (s *SqlStorePatient) GetPatientByDni(dni string) (*domain.Patient, error) {
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

func (s *SqlStorePatient) CreatePatient(patient domain.Patient) (*domain.Patient, error) {

	query := "INSERT INTO patient (Name, LastName, Address, DNI, DischargeDate) VALUES (?, ?, ?, ?, ?)"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	var id int64
	result, err := stmt.Exec(patient.Name, patient.LastName, patient.Address, patient.DNI, patient.DischargeDate)
	if err != nil {
		return nil, err
	}

	// Obtener el id autogenerado
	id, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Setear el ID autogenerado in la struct de patient para
	//mostrarlo en el resp
	patient.Id = int(id)

	return &patient, nil
}

func (s *SqlStorePatient) DeletePatient(id int) error {
	query := "DELETE FROM patient WHERE ID = ?"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStorePatient) UpdatePatient(id int, patient domain.Patient) (*domain.Patient, error) {
	updateQuery := "UPDATE patient SET Name = ?, LastName = ?, Address = ?, DischargeDate = ? WHERE ID = ?"
	_, err := s.DB.Exec(updateQuery, patient.Name, patient.LastName, patient.Address, patient.DischargeDate, id)
	if err != nil {
		return nil, err
	}
	patient.Id = id
	return &patient, nil
}
