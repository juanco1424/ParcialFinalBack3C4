-- Drop the Patients table if it exists
DROP TABLE IF EXISTS patients;

-- Create the Patients table
CREATE TABLE patients (
    ID VARCHAR(255) PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(255) NOT NULL,
    LastName VARCHAR(255) NOT NULL,
    Address VARCHAR(255) NOT NULL,
    DNI VARCHAR(20) NOT NULL,
    DischargeDate DATE
);


-- Insert sample entries
INSERT INTO patients (Name, LastName, Address, DNI, DischargeDate) VALUES
    ('John', 'Doe', '123 Main St', '1234567890', '2023-01-15'),
    ('Jane', 'Smith', '456 Elm St', '9876543210', '2023-02-20'),
    ('Alice', 'Johnson', '789 Oak St', '5678901234', '2023-02-15');
