package storage

type Comment struct {
	ID        int
	NewsID    int
	CommentID int
	Content   string
	Author    string
	PubTime   int64
}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	Comments(n int) ([]Comment, error)
	AddComment(Comment) error
}
