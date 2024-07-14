package request

type ConditionTypeCreateRequest struct {
	Name string `validate:"required" json:"name"`
}
