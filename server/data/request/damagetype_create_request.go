package request

type DamageTypeCreateRequest struct {
	Name string `validate:"required" json:"name"`
}
