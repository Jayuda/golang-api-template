package service

import (
	"voltunes-chick-api-master-product/auth"
	"voltunes-chick-api-master-product/helper"
	"voltunes-chick-api-master-product/model/domain"
	"voltunes-chick-api-master-product/model/web"
	"voltunes-chick-api-master-product/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type BankServiceImpl struct {
	BankRepository repository.BankRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewBankService(
	city repository.BankRepository,
	db *gorm.DB,
	validate *validator.Validate,
) BankService {
	return &BankServiceImpl{
		BankRepository: city,
		DB:             db,
		Validate:       validate,
	}
}

func (service *BankServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string) []web.BankResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	banks := service.BankRepository.FindAll(tx, filters)
	return banks.ToBankResponses()
}

func (service *BankServiceImpl) Create(auth *auth.AccessDetails, request *web.BankCreateRequest) web.BankResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	err = tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	helper.PanicIfError(err)
	city := &domain.Bank{
		// Required Fields
		CreatedByID: auth.UserID,
		UpdatedByID: auth.UserID,

		// Fields
		Name: request.Name,
	}
	city = service.BankRepository.Create(tx, city)

	return city.ToBankResponse()
}

func (service *BankServiceImpl) Delete(auth *auth.AccessDetails, id *int) {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	service.BankRepository.Delete(tx, id, &auth.UserID)
}

func (service *BankServiceImpl) Update(auth *auth.AccessDetails, id *int, request *web.BankUpdateRequest) web.BankResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	helper.PanicIfError(err)
	subject := &domain.Bank{
		// Required Fields
		Model:       gorm.Model{ID: uint(*id)},
		UpdatedByID: auth.UserID,

		//  Fields
		Name: request.Name,
	}
	subject = service.BankRepository.Update(tx, subject)
	return subject.ToBankResponse()
}

func (service *BankServiceImpl) FindByID(auth *auth.AccessDetails, id *int) web.BankResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	bank := service.BankRepository.FindByID(tx, id)
	return bank.ToBankResponse()
}
