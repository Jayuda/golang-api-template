package repository

import (
	"golang-api-template/helper"
	"golang-api-template/model/domain"

	"gorm.io/gorm"
)

type BankRepositoryImpl struct {
}

func NewBankRepository() BankRepository {
	return &BankRepositoryImpl{}
}

func (repository *BankRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.Banks {
	banks := domain.Banks{}
	tx := db.Model(&domain.Bank{}) // SELECT * FROM banks

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	// SELECT * FROM banks WHERE deleted_at IS NULL AND name = 'Bank Caty'

	err = tx.Find(&banks).Error
	helper.PanicIfError(err)

	return banks
}

func (repository *BankRepositoryImpl) Create(db *gorm.DB, bank *domain.Bank) *domain.Bank {
	err := db.Create(&bank).Error
	helper.PanicIfError(err)
	return bank
}

func (repository *BankRepositoryImpl) Delete(db *gorm.DB, id *int, deletedByID *uint) {
	err := db.First(&domain.Bank{}, id).Error
	helper.PanicIfError(err)

	err = db.Updates(
		&domain.Bank{
			Model:       gorm.Model{ID: uint(*id)},
			DeletedByID: deletedByID,
		},
	).Delete(&domain.Bank{}, id).Error
	helper.PanicIfError(err)
}

func (repository *BankRepositoryImpl) Update(db *gorm.DB, bank *domain.Bank) *domain.Bank {
	err := db.Updates(&bank).Error
	helper.PanicIfError(err)

	err = db.First(&bank, bank.ID).Error
	helper.PanicIfError(err)

	return bank
}

func (repository *BankRepositoryImpl) FindByID(db *gorm.DB, id *int) domain.Bank {
	var bank domain.Bank
	err := db.First(&bank, id).Error
	helper.PanicIfError(err)
	return bank
}
