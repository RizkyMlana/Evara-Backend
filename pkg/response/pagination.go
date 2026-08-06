package response

type Pagination struct {
	NextCursor string `json:"next_cursor,omitempty"`
	HasMore    bool   `json:"has_more"`
}

func NewPagination(
	nextCursor string,
	hasMore bool,
) *Pagination {
	return &Pagination{
		NextCursor: nextCursor,
		HasMore: hasMore,
	}
}