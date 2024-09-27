package repository

import (
	"golang-api-template/helper"
	"golang-api-template/model/domain"

	"gorm.io/gorm"
)

type ProductCompanyRepositoryImpl struct {
}

func NewProductCompanyRepository() ProductCompanyRepository {
	return &ProductCompanyRepositoryImpl{}
}

func (repository *ProductCompanyRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.ProductCompanys {
	productCompanys := domain.ProductCompanys{}
	tx := db.Model(&domain.ProductCompany{})

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	err = tx.Find(&productCompanys).Error
	helper.PanicIfError(err)

	return productCompanys
}

func (repository *ProductCompanyRepositoryImpl) Create(db *gorm.DB, productCompany *domain.ProductCompany) *domain.ProductCompany {
	err := db.Create(&productCompany).Error
	helper.PanicIfError(err)
	return productCompany
}

func (repository *ProductCompanyRepositoryImpl) Delete(db *gorm.DB, id *int, deletedByID *uint) {
	err := db.First(&domain.ProductCompany{}, id).Error
	helper.PanicIfError(err)

	err = db.Updates(
		&domain.ProductCompany{
			Model:       gorm.Model{ID: uint(*id)},
			DeletedByID: deletedByID,
		},
	).Delete(&domain.ProductCompany{}, id).Error
	helper.PanicIfError(err)
}

func (repository *ProductCompanyRepositoryImpl) Update(db *gorm.DB, product_company *domain.ProductCompany) *domain.ProductCompany {
	err := db.Updates(&product_company).Error
	helper.PanicIfError(err)

	err = db.First(&product_company, product_company.ID).Error
	helper.PanicIfError(err)

	return product_company
}

func (repository *ProductCompanyRepositoryImpl) FindByID(db *gorm.DB, id *int) domain.ProductCompany {
	var product_company domain.ProductCompany
	err := db.First(&product_company, id).Error
	helper.PanicIfError(err)
	return product_company
}
