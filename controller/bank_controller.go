package controller

import (
	"voltunes-chick-api-master-product/auth"

	"github.com/gin-gonic/gin"
)

type BankController interface {
	Create(context *gin.Context, auth *auth.AccessDetails)
	FindAll(context *gin.Context, auth *auth.AccessDetails)
	FindByID(context *gin.Context, auth *auth.AccessDetails)
	Delete(context *gin.Context, auth *auth.AccessDetails)
	Update(context *gin.Context, auth *auth.AccessDetails)
}
