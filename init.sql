-- Creamos BD
CREATE DATABASE IF NOT EXISTS grupo4;

-- Elegimos BD
USE grupo4;

-- Table Odontologo
CREATE TABLE IF NOT EXISTS dentist (
    id INTEGER AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(10) NOT NULL,
    lastName VARCHAR(20) NOT NULL,
    registration VARCHAR(10) NOT NULL
    );

-- Table Paciente
CREATE TABLE IF NOT EXISTS patient (
    id INTEGER AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(10) NOT NULL,
    lastName VARCHAR(20) NOT NULL,
    address VARCHAR(30) NOT NULL,
    dni VARCHAR(8) NOT NULL,
    dischargeDate VARCHAR (20)
    );

-- Table Turnos
CREATE TABLE IF NOT EXISTS appt (
    id INTEGER AUTO_INCREMENT PRIMARY KEY,
    dateTime VARCHAR(20) NOT NULL,
    description VARCHAR(255),

    foreign_key_to_dentist INT,
    FOREIGN KEY (foreign_key_to_dentist) REFERENCES dentist(id),

    foreign_key_to_patient INT,
    FOREIGN KEY (foreign_key_to_patient) REFERENCES patient(id)
);
-- Table Dentista
INSERT INTO dentist (name,lastName,registration) VALUES ('Pablo','Montivero','98');

-- Table Paciente
INSERT INTO patient (name,lastName,address,dni,dischargeDate) VALUES ('Laura','Lopez','Street 1214','21484908','01-01-2023');

-- Table Turno
INSERT INTO appt (dateTime,foreign_key_to_dentist,foreign_key_to_patient) VALUES ('01-02-2023 12:10',1,1);