package service

import (
	"fmt"
	"voltunes-chick-api-master-product/auth"
	"voltunes-chick-api-master-product/helper"
	"voltunes-chick-api-master-product/model/domain"
	"voltunes-chick-api-master-product/model/web"
	"voltunes-chick-api-master-product/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductCompanyServiceImpl struct {
	ProductCompanyRepository repository.ProductCompanyRepository
	DB                       *gorm.DB
	Validate                 *validator.Validate
}

func NewProductCompanyService(
	city repository.ProductCompanyRepository,
	db *gorm.DB,
	validate *validator.Validate,
) ProductCompanyService {
	return &ProductCompanyServiceImpl{
		ProductCompanyRepository: city,
		DB:                       db,
		Validate:                 validate,
	}
}

func (service *ProductCompanyServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string) []web.ProductCompanyResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	ProductCompanys := service.ProductCompanyRepository.FindAll(tx, filters)
	return ProductCompanys.ToProductCompanyResponses()
}

func (service *ProductCompanyServiceImpl) Create(auth *auth.AccessDetails, request *web.ProductCompanyCreateRequest) web.ProductCompanyResponse {
	fmt.Println("=======================DDD==========================")
	fmt.Println(request)
	fmt.Println("========================FFF=========================")
	err := service.Validate.Struct(request)
	//helper.PanicIfError(err)

	tx := service.DB.Begin()
	err = tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	helper.PanicIfError(err)
	city := &domain.ProductCompany{
		// Required Fields
		CreatedByID: auth.UserID,
		UpdatedByID: auth.UserID,

		// Fields
		NameCompany: request.NameCompany,
	}
	city = service.ProductCompanyRepository.Create(tx, city)

	return city.ToProductCompanyResponse()
}

func (service *ProductCompanyServiceImpl) Delete(auth *auth.AccessDetails, id *int) {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	service.ProductCompanyRepository.Delete(tx, id, &auth.UserID)
}

func (service *ProductCompanyServiceImpl) Update(auth *auth.AccessDetails, id *int, request *web.ProductCompanyUpdateRequest) web.ProductCompanyResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	helper.PanicIfError(err)
	subject := &domain.ProductCompany{
		// Required Fields
		Model:       gorm.Model{ID: uint(*id)},
		UpdatedByID: auth.UserID,

		//  Fields
		NameCompany: request.NameCompany,
	}
	subject = service.ProductCompanyRepository.Update(tx, subject)
	return subject.ToProductCompanyResponse()
}

func (service *ProductCompanyServiceImpl) FindByID(auth *auth.AccessDetails, id *int) web.ProductCompanyResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	ProductCompany := service.ProductCompanyRepository.FindByID(tx, id)
	return ProductCompany.ToProductCompanyResponse()
}
