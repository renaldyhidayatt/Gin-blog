package schemas

type SchemaArticle struct {
	ID          string             `json:"id" validate:"uuid"`
	Title       string             `json:"title" validate:"required"`
	Description string             `json:"description" validate:"required"`
	Body        string             `json:"body" validate:"required"`
	UserID      string             `json:"user_id" validate:"required"`
	Tag         []SchemaTag        `json:"tag"`
	Categories  []SchemaCategories `json:"category"`
}
