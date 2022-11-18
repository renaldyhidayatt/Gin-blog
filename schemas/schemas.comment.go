package schemas

type SchemaComment struct {
	ID      string `json:"id" validate:"uuid"`
	Content string `json:"content" validate:"required"`
}
