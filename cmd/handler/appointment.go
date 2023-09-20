package handler

import (
	"net/http"
	"parcial/internal/appointments"
	"parcial/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	AppointmentService appointments.IAppointmentService
}

func NewAppointmentHandler(appointmentService appointments.IAppointmentService) *AppointmentHandler {
	return &AppointmentHandler{
		AppointmentService: appointmentService,
	}
}

// @Summary Obtiene todas las citas
// @Description Obtiene todas las citas registradas en el sistema
// @Tags citas
// @Produce json
// @Success 200 
// @Router /appointments [get]
func (h *AppointmentHandler) GetAllAppointments(c *gin.Context) {
	appointments, err := h.AppointmentService.GetAllAppointments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, appointments)
}

// @Summary Obtiene una cita por su ID
// @Description Obtiene la cita correspondiente al ID proporcionado
// @Tags citas
// @Produce json
// @Param id path int true "ID de la cita"
// @Success 200 
// @Router /appointments/{id} [get]
func (h *AppointmentHandler) GetAppointmentById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	appointment, err := h.AppointmentService.GetAppointmentById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

// @Summary Obtiene todas las citas de un paciente por su DNI
// @Description Obtiene todas las citas registradas en el sistema para un paciente específico
// @Tags citas
// @Produce json
// @Param dni path string true "DNI del paciente"
// @Success 200 
// @Router /appointments/dni/{dni} [get]
func (h *AppointmentHandler) GetAppointmentsByDni(c *gin.Context) {
	dni := c.Param("dni") // Obtener el parámetro "dni" de la URL
	appointments, err := h.AppointmentService.GetAppointmentsByDni(dni)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

// @Summary Crea una nueva cita
// @Description Crea una nueva cita en el sistema
// @Tags citas
// @Accept json
// @Produce json
// @Param appointment body domain.Appointment true "Datos de la cita"
// @Success 201 
// @Router /appointments [post]
func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
	var appointment domain.Appointment
	if err := c.BindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de cita no válidos"})
		return
	}

	createdAppointment, err := h.AppointmentService.CreateAppointment(appointment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdAppointment)
}

// @Summary Actualiza una cita existente
// @Description Actualiza los datos de una cita existente en el sistema
// @Tags citas
// @Accept json
// @Produce json
// @Param id path int true "ID de la cita"
// @Param appointment body domain.Appointment true "Nuevos datos de la cita"
// @Success 200 
// @Router /appointments/{id} [put]
func (h *AppointmentHandler) UpdateAppointment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var appointment domain.Appointment
	if err := c.BindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de cita no válidos"})
		return
	}

	updatedAppointment, err := h.AppointmentService.UpdateAppointment(id, appointment)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedAppointment)
}

// @Summary Elimina una cita existente
// @Description Elimina una cita existente en el sistema
// @Tags citas
// @Param id path int true "ID de la cita"
// @Success 204 "Cita eliminada exitosamente"
// @Router /appointments/{id} [delete]
func (h *AppointmentHandler) DeleteAppointment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.AppointmentService.DeleteAppointment(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

