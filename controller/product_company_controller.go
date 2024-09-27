package controller

import (
	"golang-api-template/auth"

	"github.com/gin-gonic/gin"
)

type ProductCompanyController interface {
	Create(context *gin.Context, auth *auth.AccessDetails)
	FindAll(context *gin.Context, auth *auth.AccessDetails)
	FindByID(context *gin.Context, auth *auth.AccessDetails)
	Delete(context *gin.Context, auth *auth.AccessDetails)
	Update(context *gin.Context, auth *auth.AccessDetails)
}
