package handler

import (
	"errors"
	"net/http"
	"parcial/internal/domain"
	"parcial/internal/odontologo"
	"parcial/pkg/web"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DentistHandler struct {
	IService odontologo.IServiceDentist
}

// @Summary Obtiene un dentista por su ID
// @Description Obtiene la información de un dentista según su ID
// @Tags dentistas
// @Produce json
// @Param id path int true "ID del dentista"
// @Success 200
// @Router /dentists/{id} [get]
func (ph *DentistHandler) GetDentistById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.FailResp(ctx, http.StatusBadRequest, errors.New("id invalido, ingrese un formato valido"))
		return
	}
	dentist, err := ph.IService.GetDentistById(id)
	if err != nil {
		web.FailResp(ctx, http.StatusNotFound, err)
		return
	}
	web.SuccessResp(ctx, http.StatusOK, dentist)
}

// @Summary Crea un nuevo dentista
// @Description Crea un nuevo dentista en el sistema
// @Tags dentistas
// @Accept json
// @Produce json
// @Param dentist body domain.Dentist true "Datos del dentista"
// @Success 201
// @Router /dentists [post]
func (ph *DentistHandler) CreateDentist(ctx *gin.Context) {
	var dentist domain.Dentist
	if err := ctx.ShouldBindJSON(&dentist); err != nil {
		web.FailResp(ctx, http.StatusBadRequest, err)
		return
	}

	p, err := ph.IService.CreateDentist(dentist)
	if err != nil {
		web.FailResp(ctx, http.StatusBadRequest, err)
		return
	}
	web.SuccessResp(ctx, http.StatusCreated, p)
}

// @Summary Actualiza un dentista existente
// @Description Actualiza los datos de un dentista existente en el sistema
// @Tags dentistas
// @Accept json
// @Produce json
// @Param id path int true "ID del dentista"
// @Param dentist body domain.Dentist true "Nuevos datos del dentista"
// @Success 200
// @Router /dentists/{id} [put]
func (ph *DentistHandler) UpdateDentist(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.FailResp(ctx, http.StatusBadRequest, errors.New("id invalido, ingrese un formato valido"))
		return
	}
	var updatedDentist domain.Dentist
	if err := ctx.ShouldBindJSON(&updatedDentist); err != nil {
		web.FailResp(ctx, http.StatusNotFound, err)
		return
	}

	updatedDentistResult, err := ph.IService.UpdateDentist(id, updatedDentist)
	if err != nil {
		web.FailResp(ctx, http.StatusNotFound, err)
		return
	}
	web.SuccessResp(ctx, http.StatusOK, updatedDentistResult)
}

// @Summary Elimina un dentista existente
// @Description Elimina un dentista existente en el sistema
// @Tags dentistas
// @Param id path int true "ID del dentista"
// @Success 204 "Dentista eliminado exitosamente"
// @Router /dentists/{id} [delete]
func (ph *DentistHandler) DeleteDentist(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.FailResp(ctx, http.StatusBadRequest, errors.New("id invalido, ingrese un formato valido"))
		return
	}
	if err := ph.IService.DeleteDentist(id); err != nil {
		web.FailResp(ctx, http.StatusNotFound, err)
		return
	}

	web.SuccessResp(ctx, http.StatusNoContent, "")
}
