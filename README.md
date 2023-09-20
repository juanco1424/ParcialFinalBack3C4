# ParcialFinalBack3C4

## Sistema de reserva de turnos
Implementacion de una API que permite administrar la reserva de turnos para una clinica odontologica. El proyecto está desarrollado en Go, utiliza el framework Gin y docker.

## Instalación

### Prerrequisitos

- Asegúrate de tener instalado Go en tu sistema. Puedes descargarlo desde [sitio oficial de Go](https://golang.org/dl/).
- Asegurate de tener instalado Docker. Puedes descargarlo desde [sitio oficial de Docker](https://docs.docker.com/get-docker/).

### Pasos de instalación

1. Clona este repositorio en tu sistema local:

   ```bash
   git clone https://github.com/juanco1424/ParcialFinalBack3C4.git
   cd ParcialFinalBack3C4

2.  Ejecuta el siguiente comando para construir y levantar el contenedor de Docker:
    ```bash
    docker-compose up -d

3. Inicia el servidor de desarrollo:
   ```bash
    cd cmd && go run main.go

### Uso
La aplicación está configurada para ejecutarse en el puerto 8080. Base url: http://localhost:8080/

#### Endpoints
Patient Endpoints:

    GET  /patient/{id}: Obtiene un paciente por su ID.
    POST /patient: Crea un nuevo paciente.
    PUT  /patient/{id}: Actualiza la información de un paciente existente por su ID.
    DELETE /patient/{id}: Elimina un paciente por su ID.

Dentist Endpoints:

    GET /dentist/{id}: Obtiene un dentista por su ID.
    POST /dentist: Crea un nuevo dentista.
    PUT /dentist/{id}: Actualiza la información de un dentista existente por su ID.
    DELETE /dentist/{id}: Elimina un dentista por su ID.

Appointment Endpoints:

    GET /appointments: Obtiene todos los turnos.
    GET /appointments/{id}: Obtiene un turno por su ID.
    GET /appointments/patient/{dni}: Obtiene todos los turnos de un paciente por su DNI.
    POST /appointments: Crea un nuevo turno.
    PUT /appointments/{id}: Actualiza la información de un turno existente por su ID.
    DELETE /appointments/{id}: Elimina un turno por su ID.
