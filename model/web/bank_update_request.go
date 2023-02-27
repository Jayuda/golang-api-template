package web

type BankUpdateRequest struct {
	// Fields
	Name string `json:"name" validate:"required,min=1,max=100"`
}
