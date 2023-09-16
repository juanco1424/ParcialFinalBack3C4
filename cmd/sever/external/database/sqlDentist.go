package database

import (
	"database/sql"
	"fmt"
	"parcial/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func (s *SqlStore) CreateDentist(dentist domain.Dentist) (*domain.Dentist, error) {
	query := "INSERT INTO dentist (id, lastname, name, registration) VALUES (?, ?, ?, ?)"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(dentist.ID, dentist.LastName, dentist.Name, dentist.Registration)
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &dentist, nil
}

func (s *SqlStore) GetDentistById(id int) (*domain.Dentist, error) {
	var dentist domain.Dentist
	query := "SELECT * FROM dentist WHERE ID = ?;"
	row := s.DB.QueryRow(query, id)

	err := row.Scan(&dentist.ID, &dentist.LastName, &dentist.Name, &dentist.Registration)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no se encontró ningún dentista con el ID proporcionado")
		}
		return nil, err
	}

	return &dentist, nil
}

func (s *SqlStore) GetDentistByRegistration(registration string) (*domain.Dentist, error) {
	var dentist domain.Dentist
	query := "SELECT * FROM dentist WHERE registration = ?;"
	row := s.DB.QueryRow(query, registration)
	err := row.Scan(&dentist.ID, &dentist.LastName, &dentist.Name, &dentist.Registration)
	if err != nil {
		return nil, err
	}
	println(&dentist)
	return &dentist, nil
}

func (s *SqlStore) UpdateDentist(id int, dentist domain.Dentist) (*domain.Dentist, error) {
	updateQuery := "UPDATE dentist SET lastname = ?, name = ?, registration = ? WHERE ID = ?"
	_, err := s.DB.Exec(updateQuery, dentist.LastName, dentist.Name, dentist.Registration, id)
	if err != nil {
		return nil, err
	}
	return &dentist, nil
}

func (s *SqlStore) DeleteDentist(id int) error {
	query := "DELETE FROM dentist WHERE ID = ?"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
