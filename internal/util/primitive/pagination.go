package primitive

type PaginationInput struct {
	Page     int
	PageSize int
}

type PaginationOutput struct {
	Page      int
	PageSize  int
	PageCount int
	TotalData int
}

func (p PaginationInput) Offset() int {
	offset := int(0)
	if p.Page > 0 {
		offset = (p.Page - 1) * p.PageSize
	}

	return offset
}

func getPageCount(pageSize, totalData int) int {
	pageCount := int(1)
	if pageSize > 0 {
		if pageSize >= totalData {
			return pageCount
		}

		if totalData%pageSize == 0 {
			pageCount = totalData / pageSize
		} else {
			pageCount = totalData/pageSize + 1
		}
	}

	return pageCount
}

func CreatePaginationOutput(input PaginationInput, totalData int) PaginationOutput {
	pageCount := getPageCount(input.PageSize, totalData)
	return PaginationOutput{
		Page:      input.Page,
		PageSize:  input.PageSize,
		TotalData: totalData,
		PageCount: pageCount,
	}
}
