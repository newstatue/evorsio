package common

type PageQuery struct {
	Cursor string
	Size   int
}

func (p PageQuery) Init() {
	if p.Size <= 0 {
		p.Size = 20
	}
}

type PageResult[T any] struct {
	Items      []T
	NextCursor string
	HasMore    bool
}
