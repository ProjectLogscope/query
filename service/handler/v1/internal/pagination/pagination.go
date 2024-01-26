package pagination

import (
	"strconv"
)

type pagination struct {
	Page  int
	Count int
}

func Parse(m map[string]string) pagination {
	p := pagination{
		Page:  DefaultPageValue,
		Count: DefaultCountValue,
	}

	for k, v := range m {
		switch k {
		case "paginationPage":
			if page, err := strconv.Atoi(v); err == nil {
				p.Page = min(max(page, MinPageValue), MaxPageValue)
			}
		case "paginationCount":
			if count, err := strconv.Atoi(v); err == nil {
				p.Count = min(max(count, MinCountValue), MaxCountValue)
			}
		}
	}
	return p
}
