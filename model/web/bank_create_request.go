package web

type BankCreateRequest struct {
	// Fields
	Name string `json:"name" validate:"required,min=1,max=100"`
}
