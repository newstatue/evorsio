package common

type PageQuery struct {
	Cursor string `json:"cursor"`
	Size   int    `json:"limit"`
}

func (p PageQuery) Init() {
	if p.Size <= 0 {
		p.Size = 20
	}
}

type PageResult[T any] struct {
	Items      []T    `json:"items"`
	NextCursor string `json:"nextCursor,omitempty"`
	HasMore    bool   `json:"hasMore"`
}
