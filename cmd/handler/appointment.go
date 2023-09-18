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

func (h *AppointmentHandler) GetAllAppointments(c *gin.Context) {
	appointments, err := h.AppointmentService.GetAllAppointments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, appointments)
}

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

func (h *AppointmentHandler) GetAppointmentsByDni(c *gin.Context) {
	dni := c.Param("dni") // Obtener el parámetro "dni" de la URL
	appointments, err := h.AppointmentService.GetAppointmentsByDni(dni)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

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
