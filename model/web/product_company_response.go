package web

import "time"

type ProductCompanyResponse struct {
	// Required Fields
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedByID uint      `json:"created_by_id"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedByID uint      `json:"updated_by_id"`

	// Fields
	NameCompany string `json:"name_company"`
}
