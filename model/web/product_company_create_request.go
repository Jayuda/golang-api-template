package web

type ProductCompanyCreateRequest struct {
	// Fields
	NameCompany string `json:"name_company" validate:"required,min=1,max=100"`
}
