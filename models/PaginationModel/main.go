package PaginationModel

type Pagination struct {
	Limit int8   `json:"limit"`
	Page  int8   `json:"page"`
	Order string `json:"order"`
}
