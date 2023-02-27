package controller

import (
	"net/http"
	"strconv"
	"voltunes-chick-api-master-product/auth"
	"voltunes-chick-api-master-product/helper"
	"voltunes-chick-api-master-product/model/web"
	"voltunes-chick-api-master-product/service"

	"github.com/gin-gonic/gin"
)

type BankControllerImpl struct {
	BankService service.BankService
}

func NewBankController(bankService service.BankService) BankController {
	return &BankControllerImpl{
		BankService: bankService,
	}
}

func (controller *BankControllerImpl) FindAll(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "name.like")
	bankResponses := controller.BankService.FindAll(auth, &filters)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(bankResponses),
		Data:    bankResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *BankControllerImpl) FindByID(c *gin.Context, auth *auth.AccessDetails) {
	paramID := c.Param("id")
	bankID, err := strconv.Atoi(paramID)
	helper.PanicIfError(err)

	bankResponse := controller.BankService.FindByID(auth, &bankID)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(bankResponse),
		Data:    bankResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *BankControllerImpl) Create(c *gin.Context, auth *auth.AccessDetails) {
	request := web.BankCreateRequest{}
	helper.ReadFromRequestBody(c, &request)

	bankResponse := controller.BankService.Create(auth, &request)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Bank created successfully",
		Data:    bankResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *BankControllerImpl) Update(c *gin.Context, auth *auth.AccessDetails) {
	paramID := c.Param("id")
	bankID, err := strconv.Atoi(paramID)
	helper.PanicIfError(err)

	request := web.BankUpdateRequest{}
	helper.ReadFromRequestBody(c, &request)

	bankResponse := controller.BankService.Update(auth, &bankID, &request)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Bank updated successfully",
		Data:    bankResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *BankControllerImpl) Delete(c *gin.Context, auth *auth.AccessDetails) {
	paramID := c.Param("id")
	bankID, err := strconv.Atoi(paramID)
	helper.PanicIfError(err)

	controller.BankService.Delete(auth, &bankID)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Bank deleted successfully",
	}

	c.JSON(http.StatusOK, webResponse)
}
