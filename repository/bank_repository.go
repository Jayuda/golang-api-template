package repository

import (
	"voltunes-chick-api-master-product/model/domain"

	"gorm.io/gorm"
)

type BankRepository interface {
	Create(db *gorm.DB, bank *domain.Bank) *domain.Bank
	Delete(db *gorm.DB, id *int, deletedByID *uint)
	FindAll(db *gorm.DB, filters *map[string]string) domain.Banks
	FindByID(db *gorm.DB, id *int) domain.Bank
	Update(db *gorm.DB, bank *domain.Bank) *domain.Bank
}
