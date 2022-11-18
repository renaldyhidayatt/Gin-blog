package schemas

type SchemaTag struct {
	ID          string `json:"id" validate:"uuid"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
