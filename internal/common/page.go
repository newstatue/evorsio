package common

type PageQuery struct {
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

type PageResult[T any] struct {
	Items      []T    `json:"items"`
	NextCursor string `json:"nextCursor,omitempty"`
	HasMore    bool   `json:"hasMore"`
	Total      int    `json:"total"`
}
