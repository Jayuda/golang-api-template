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

func NathanRoute(router *gin.Engine, DB *gorm.DB, validate *validator.Validate) {

	nathanService := service.NewNathanService(
		repository.NewNathanRepository(),
		DB,
		validate,
	)
	nathanController := controller.NewNathanController(nathanService)

	router.DELETE("/nathans/:id", auth.Auth(nathanController.Delete, []string{}))
	router.GET("/nathans", auth.Auth(nathanController.FindAll, []string{}))
	router.GET("/nathans/:id", auth.Auth(nathanController.FindByID, []string{}))
	router.POST("/nathans", auth.Auth(nathanController.Create, []string{}))
	router.PUT("/nathans/:id", auth.Auth(nathanController.Update, []string{}))
}
