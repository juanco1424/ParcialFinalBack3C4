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
	IService odontologo.IService
}

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
