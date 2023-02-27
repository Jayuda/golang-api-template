package service

import (
	"voltunes-chick-api-master-product/auth"
	"voltunes-chick-api-master-product/model/web"
)

type ProductCompanyService interface {
	Create(auth *auth.AccessDetails, request *web.ProductCompanyCreateRequest) web.ProductCompanyResponse
	Delete(auth *auth.AccessDetails, id *int)
	Update(auth *auth.AccessDetails, id *int, request *web.ProductCompanyUpdateRequest) web.ProductCompanyResponse
	FindAll(auth *auth.AccessDetails, filters *map[string]string) []web.ProductCompanyResponse
	FindByID(auth *auth.AccessDetails, id *int) web.ProductCompanyResponse
}
