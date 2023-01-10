package schemas

type SchemaAuth struct {
	FirstName string `json:"firstname" validate:"required,lowercase"`
	LastName  string `json:"lastname" validate:"required,lowercase"`
	Bio       string `json:"bio" validate:"required,lowercase"`
	Email     string `json:"email" validate:"required,lowercase"`
	Password  string `json:"password" validate:"required,lowercase"`
}
