package controller

import (
	"encoding/json"
	"net/http"
	"task/internal/app/service"
	"task/pkg/e"
	"task/pkg/response"

	"github.com/go-chi/chi"
)

type StatusController interface {
	GetDetailsByCountry(http.ResponseWriter, *http.Request)
}

type statusControllerImpl struct {
	StatusService service.StatusService
}

func NewStatusController(statusService service.StatusService) StatusController {
	return &statusControllerImpl{
		StatusService: statusService,
	}
}

func (c *statusControllerImpl) GetDetailsByCountry(w http.ResponseWriter, r *http.Request) {
	country := chi.URLParam(r, "country")
	res, err := c.StatusService.GetDetailsByCountry(country)
	if err != nil {
		if err.Error() == "record not found" {
			response.Fail(w, http.StatusNotFound, e.ErrorCodeInvalidCountry, err.Error())
			return
		}
		response.Fail(w, http.StatusInternalServerError, e.ErrorInternalServerError, err.Error())
		return
	}
	jsonResponse, err := json.Marshal(res)
	if err != nil {
		response.Fail(w, http.StatusInternalServerError, e.ErrorInternalServerError, "Failed to marshal response data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
