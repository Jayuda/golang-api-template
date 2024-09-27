package domain

import (
	"golang-api-template/model/web"

	"gorm.io/gorm"
)

type ProductCompanys []ProductCompany
type ProductCompany struct {
	// Required Fields
	gorm.Model
	CreatedByID uint  `gorm:""`
	UpdatedByID uint  `gorm:""`
	DeletedByID *uint `gorm:""`

	// Fields
	NameCompany string `gorm:"size:100;not null"`
}

func (product_company *ProductCompany) ToProductCompanyResponse() web.ProductCompanyResponse {
	return web.ProductCompanyResponse{
		// Required Fields
		ID:          product_company.ID,
		CreatedAt:   product_company.CreatedAt,
		CreatedByID: product_company.CreatedByID,
		UpdatedAt:   product_company.UpdatedAt,
		UpdatedByID: product_company.UpdatedByID,

		// Fields
		NameCompany: product_company.NameCompany,
	}
}

func (product_companys ProductCompanys) ToProductCompanyResponses() []web.ProductCompanyResponse {
	product_companyResponses := []web.ProductCompanyResponse{}
	for _, product_company := range product_companys {
		product_companyResponses = append(product_companyResponses, product_company.ToProductCompanyResponse())
	}
	return product_companyResponses
}
