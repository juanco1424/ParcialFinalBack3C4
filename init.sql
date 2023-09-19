-- Creamos BD
CREATE DATABASE IF NOT EXISTS dh202302g4;

-- Elegimos BD
USE dh202302g4;

-- Table Odontologo
CREATE TABLE IF NOT EXISTS dentist (
    id INTEGER AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(10) NOT NULL,
    lastName VARCHAR(20) NOT NULL,
    Registration VARCHAR(10) NOT NULL
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
CREATE TABLE IF NOT EXISTS appointment (
    id INTEGER AUTO_INCREMENT PRIMARY KEY,
    date VARCHAR(20) NOT NULL,
    hour VARCHAR(20) NOT NULL,
    description VARCHAR(255) NOT NULL,

    dentist_id INT,
    FOREIGN KEY (dentist_id) REFERENCES dentist(id),

    patient_id INT,
    FOREIGN KEY (patient_id) REFERENCES patient(id)
);
-- Table Dentista
INSERT INTO dentist (name,lastName,registration) VALUES ('Pablo','Montivero','98');

-- Table Paciente
INSERT INTO patient (name,lastName,address,dni,dischargeDate) VALUES ('Laura','Lopez','Street 1214','21484908','01-01-2023');

-- Table Turno
INSERT INTO appointment (date,hour,description,dentist_id,patient_id) VALUES ('01-02-2023','12:10','description',1,1);