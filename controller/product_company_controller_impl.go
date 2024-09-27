package controller

import (
	"golang-api-template/auth"
	"golang-api-template/helper"
	"golang-api-template/model/web"
	"golang-api-template/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductCompanyControllerImpl struct {
	ProductCompanyService service.ProductCompanyService
}

func NewProductCompanyController(ProductCompanyService service.ProductCompanyService) ProductCompanyController {
	return &ProductCompanyControllerImpl{
		ProductCompanyService: ProductCompanyService,
	}
}

func (controller *ProductCompanyControllerImpl) FindAll(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "name.like")
	ProductCompanyResponses := controller.ProductCompanyService.FindAll(auth, &filters)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(ProductCompanyResponses),
		Data:    ProductCompanyResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *ProductCompanyControllerImpl) FindByID(c *gin.Context, auth *auth.AccessDetails) {
	paramID := c.Param("id")
	ProductCompanyID, err := strconv.Atoi(paramID)
	helper.PanicIfError(err)

	ProductCompanyResponse := controller.ProductCompanyService.FindByID(auth, &ProductCompanyID)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(ProductCompanyResponse),
		Data:    ProductCompanyResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *ProductCompanyControllerImpl) Create(c *gin.Context, auth *auth.AccessDetails) {
	request := web.ProductCompanyCreateRequest{}
	helper.ReadFromRequestBody(c, &request)

	ProductCompanyResponse := controller.ProductCompanyService.Create(auth, &request)
	webResponse := web.WebResponse{
		Success: true,
		Message: "ProductCompany created successfully",
		Data:    ProductCompanyResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *ProductCompanyControllerImpl) Update(c *gin.Context, auth *auth.AccessDetails) {
	paramID := c.Param("id")
	ProductCompanyID, err := strconv.Atoi(paramID)
	helper.PanicIfError(err)

	request := web.ProductCompanyUpdateRequest{}
	helper.ReadFromRequestBody(c, &request)

	ProductCompanyResponse := controller.ProductCompanyService.Update(auth, &ProductCompanyID, &request)
	webResponse := web.WebResponse{
		Success: true,
		Message: "ProductCompany updated successfully",
		Data:    ProductCompanyResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *ProductCompanyControllerImpl) Delete(c *gin.Context, auth *auth.AccessDetails) {
	paramID := c.Param("id")
	ProductCompanyID, err := strconv.Atoi(paramID)
	helper.PanicIfError(err)

	controller.ProductCompanyService.Delete(auth, &ProductCompanyID)
	webResponse := web.WebResponse{
		Success: true,
		Message: "ProductCompany deleted successfully",
	}

	c.JSON(http.StatusOK, webResponse)
}
