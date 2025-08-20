package storage

type Comment struct {
	ID        int    `json:"id"`
	NewsID    int    `json:"news_id"`
	CommentID int    `json:"comment_id"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	PubTime   int64  `json:"pub_time"`
}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	Comments(n int) ([]Comment, error)
	AddComment(Comment) error
}
