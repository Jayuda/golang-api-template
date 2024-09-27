package repository

import (
	"golang-api-template/model/domain"

	"gorm.io/gorm"
)

type ProductCompanyRepository interface {
	Create(db *gorm.DB, bank *domain.ProductCompany) *domain.ProductCompany
	Delete(db *gorm.DB, id *int, deletedByID *uint)
	FindAll(db *gorm.DB, filters *map[string]string) domain.ProductCompanys
	FindByID(db *gorm.DB, id *int) domain.ProductCompany
	Update(db *gorm.DB, bank *domain.ProductCompany) *domain.ProductCompany
}
