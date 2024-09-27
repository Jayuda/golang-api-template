package domain

import (
	"golang-api-template/model/web"

	"gorm.io/gorm"
)

type Banks []Bank
type Bank struct {
	// Required Fields
	gorm.Model
	CreatedByID uint  `gorm:""`
	UpdatedByID uint  `gorm:""`
	DeletedByID *uint `gorm:""`

	// Fields
	Name string `gorm:"size:100;not null"`
}

func (bank *Bank) ToBankResponse() web.BankResponse {
	return web.BankResponse{
		// Required Fields
		ID:          bank.ID,
		CreatedAt:   bank.CreatedAt,
		CreatedByID: bank.CreatedByID,
		UpdatedAt:   bank.UpdatedAt,
		UpdatedByID: bank.UpdatedByID,

		// Fields
		Name: bank.Name,
	}
}

func (banks Banks) ToBankResponses() []web.BankResponse {
	bankResponses := []web.BankResponse{}
	for _, bank := range banks {
		bankResponses = append(bankResponses, bank.ToBankResponse())
	}
	return bankResponses
}
