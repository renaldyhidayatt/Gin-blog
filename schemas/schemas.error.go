package schemas

type SchemaDatabaseError struct {
	Type string
	Code int
}

type SchemaErrorResponse struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

type SchemaUnathorizatedError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
