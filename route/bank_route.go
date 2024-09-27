package route

import (
	"golang-api-template/auth"
	"golang-api-template/controller"
	"golang-api-template/repository"
	"golang-api-template/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func BankRoute(router *gin.Engine, DB *gorm.DB, validate *validator.Validate) {

	bankService := service.NewBankService(
		repository.NewBankRepository(),
		DB,
		validate,
	)
	bankController := controller.NewBankController(bankService)

	router.DELETE("/banks/:id", auth.Auth(bankController.Delete, []string{}))
	router.GET("/banks", auth.Auth(bankController.FindAll, []string{}))
	router.GET("/banks/:id", auth.Auth(bankController.FindByID, []string{}))
	router.POST("/banks", auth.Auth(bankController.Create, []string{}))
	router.PUT("/banks/:id", auth.Auth(bankController.Update, []string{}))
}
