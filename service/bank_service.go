package service

import (
	"voltunes-chick-api-master-product/auth"
	"voltunes-chick-api-master-product/model/web"
)

type BankService interface {
	Create(auth *auth.AccessDetails, request *web.BankCreateRequest) web.BankResponse
	Delete(auth *auth.AccessDetails, id *int)
	Update(auth *auth.AccessDetails, id *int, request *web.BankUpdateRequest) web.BankResponse
	FindAll(auth *auth.AccessDetails, filters *map[string]string) []web.BankResponse
	FindByID(auth *auth.AccessDetails, id *int) web.BankResponse
}
