package schemas

type SchemaResponses struct {
	StatusCode int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type Pagination struct {
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	Count      int64   `json:"count"`
	TotalPages float64 `json:"total_pages"`
}
