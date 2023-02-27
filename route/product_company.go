package route

import (
	"voltunes-chick-api-master-product/auth"
	"voltunes-chick-api-master-product/controller"
	"voltunes-chick-api-master-product/repository"
	"voltunes-chick-api-master-product/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ProductCompanyRoute(router *gin.Engine, DB *gorm.DB, validate *validator.Validate) {

	product_companyService := service.NewProductCompanyService(
		repository.NewProductCompanyRepository(),
		DB,
		validate,
	)
	product_companyController := controller.NewProductCompanyController(product_companyService)

	router.DELETE("/product_company/:id", auth.Auth(product_companyController.Delete, []string{}))
	router.GET("/product_company", auth.Auth(product_companyController.FindAll, []string{}))
	router.GET("/product_company/:id", auth.Auth(product_companyController.FindByID, []string{}))
	router.POST("/product_company", auth.Auth(product_companyController.Create, []string{}))
	router.PUT("/product_company/:id", auth.Auth(product_companyController.Update, []string{}))
}
