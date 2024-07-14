package request

type SavingThrowCreateRequest struct {
	Name string `validate:"required" json:"name"`
}
