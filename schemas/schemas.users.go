package schemas

type SchemasUser struct {
	ID        string `json:"id" validate:"uuid"`
	FirstName string `json:"firstname" validate:"required,lowercase"`
	LastName  string `json:"lastname" validate:"required,lowercase"`
	Bio       string `json:"bio" validate:"required,lowercase"`
	Image     string `json:"image"`
	Email     string `json:"email" validate:"required,lowercase"`
	Password  string `json:"password" validate:"required,lowercase"`
}
