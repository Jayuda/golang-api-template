package service

import (
	"golang-api-template/auth"
	"golang-api-template/model/web"
)

type ProductCompanyService interface {
	Create(auth *auth.AccessDetails, request *web.ProductCompanyCreateRequest) web.ProductCompanyResponse
	Delete(auth *auth.AccessDetails, id *int)
	Update(auth *auth.AccessDetails, id *int, request *web.ProductCompanyUpdateRequest) web.ProductCompanyResponse
	FindAll(auth *auth.AccessDetails, filters *map[string]string) []web.ProductCompanyResponse
	FindByID(auth *auth.AccessDetails, id *int) web.ProductCompanyResponse
}
