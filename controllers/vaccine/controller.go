package vaccine

import (
	"ca-reservaksin/businesses/vaccine"
	"ca-reservaksin/controllers"
	"ca-reservaksin/controllers/vaccine/request"
	"ca-reservaksin/controllers/vaccine/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type VaccineController struct {
	VaccineService vaccine.Service
}

func NewVaccineController(vaccineCtrl vaccine.Service) *VaccineController {
	return &VaccineController{
		VaccineService: vaccineCtrl,
	}
}

func (ctrl *VaccineController) Create(c echo.Context) error {
	req := request.Vaccine{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.VaccineService.Create(req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (ctrl *VaccineController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	req := request.Vaccine{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.VaccineService.Update(id, req.ToDomain())
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
		}

		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (ctrl *VaccineController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	res, err := ctrl.VaccineService.Delete(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, res)
}

func (ctrl *VaccineController) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	req := request.Vaccine{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.VaccineService.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (ctrl *VaccineController) FetchAll(c echo.Context) error {
	res, err := ctrl.VaccineService.FetchAll()
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomainArray(res))
}
