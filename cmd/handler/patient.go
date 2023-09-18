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
