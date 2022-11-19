package schemas

type SchemaComment struct {
	ID      string `json:"id" validate:"uuid"`
	Content string `json:"content" validate:"required"`
	UserID  string `json:"user_id" validate:"uuid"`
}
