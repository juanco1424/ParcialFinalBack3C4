package handler

import (
	"errors"
	"net/http"
	"parcial/internal/domain"
	"parcial/internal/paciente"
	"parcial/pkg/web"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	PService paciente.IServicePatient
}

// @Summary Obtiene un paciente por su ID
// @Description Obtiene la información de un paciente según su ID
// @Tags pacientes
// @Produce json
// @Param id path int true "ID del paciente"
// @Success 200
// @Router /patients/{id} [get]
func (ph *PatientHandler) GetPatientById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.FailResp(ctx, http.StatusBadRequest, errors.New("id invalido, ingrese un formato valido"))
		return
	}
	patient, err := ph.PService.GetPatientById(id)
	if err != nil {
		web.FailResp(ctx, http.StatusNotFound, err)
		return
	}
	web.SuccessResp(ctx, http.StatusOK, patient)
}

// @Summary Crea un nuevo paciente
// @Description Crea un nuevo paciente en el sistema
// @Tags pacientes
// @Accept json
// @Produce json
// @Param patient body domain.Patient true "Datos del paciente"
// @Success 201
// @Router /patients [post]
func (ph *PatientHandler) CreatePatient(ctx *gin.Context) {
	var patient domain.Patient
	if err := ctx.ShouldBindJSON(&patient); err != nil {
		web.FailResp(ctx, http.StatusBadRequest, err)
		return
	}

	p, err := ph.PService.CreatePatient(patient)
	if err != nil {
		web.FailResp(ctx, http.StatusBadRequest, err)
		return
	}
	web.SuccessResp(ctx, http.StatusCreated, p)
}

// @Summary Elimina un paciente existente
// @Description Elimina un paciente existente en el sistema
// @Tags pacientes
// @Param id path int true "ID del paciente"
// @Success 204 "Paciente eliminado exitosamente"
// @Router /patients/{id} [delete]
func (ph *PatientHandler) DeletePatient(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.FailResp(ctx, http.StatusBadRequest, errors.New("id invalido, ingrese un formato valido"))
		return
	}
	if err := ph.PService.DeletePatient(id); err != nil {
		web.FailResp(ctx, http.StatusNotFound, err)
		return
	}

	web.SuccessResp(ctx, http.StatusNoContent, "")
}

// @Summary Actualiza un paciente existente
// @Description Actualiza los datos de un paciente existente en el sistema
// @Tags pacientes
// @Accept json
// @Produce json
// @Param id path int true "ID del paciente"
// @Param patient body domain.Patient true "Nuevos datos del paciente"
// @Success 200
// @Router /patients/{id} [put]
func (ph *PatientHandler) UpdatePatient(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.FailResp(ctx, http.StatusBadRequest, errors.New("id invalido, ingrese un formato valido"))
		return
	}
	var updatedPatient domain.Patient
	if err := ctx.ShouldBindJSON(&updatedPatient); err != nil {
		web.FailResp(ctx, http.StatusNotFound, err)
		return
	}

	updatedPatientResult, err := ph.PService.UpdatePatient(id, updatedPatient)
	if err != nil {
		web.FailResp(ctx, http.StatusNotFound, err)
		return
	}
	web.SuccessResp(ctx, http.StatusOK, updatedPatientResult)
}
