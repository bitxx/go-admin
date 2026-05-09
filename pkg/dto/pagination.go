package dto

type Pagination struct {
	PageIndex int `form:"pageIndex"`
	PageSize  int `form:"pageSize"`
}

func (m *Pagination) GetPageIndex() int {
	if m.PageIndex <= 0 {
		m.PageIndex = 1
	}
	return m.PageIndex
}

// GetPageSize 最大只能单页100条
func (m *Pagination) GetPageSize() int {
	if m.PageSize <= 0 {
		m.PageSize = 10
	}
	if m.PageSize > 100 {
		m.PageSize = 100
	}
	return m.PageSize
}
