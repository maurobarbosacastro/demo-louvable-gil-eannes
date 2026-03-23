package pagination

type PaginationResult struct {
	Limit      int         `json:"limit,omitempty"`
	Page       int         `json:"page,omitempty"`
	Sort       string      `json:"sort,omitempty"`
	TotalRows  int64       `json:"totalRows"`
	TotalPages int         `json:"totalPages"`
	Data       interface{} `json:"data"`
}

type PaginationParams struct {
	Limit int    `json:"limit" query:"limit"`
	Page  int    `json:"page" query:"page"`
	Sort  string `json:"sort" query:"sort"`
}

func (p PaginationResult) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}
func (p *PaginationResult) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}
func (p *PaginationResult) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}
func (p *PaginationResult) GetSort() string {
	if p.Sort == "" {
		p.Sort = "uuid desc"
	}
	return p.Sort
}
