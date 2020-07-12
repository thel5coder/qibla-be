package requests

type TermConditionRequest struct {
	TermName    string `json:"term_name" validate:"required"`
	TermType    string `json:"term_type" validate:"required"`
	Description string `json:"description" validate:"required"`
}
