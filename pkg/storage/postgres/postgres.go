package postgres

import (
	"comments-service/pkg/storage"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db *pgxpool.Pool
}

func New(dbURL string) (*Store, error) {
	db, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func (s *Store) Comments(n int) ([]storage.Comment, error) {
	if n == 0 {
		n = 30
	}
	rows, err := s.db.Query(context.Background(), `
	SELECT c.id, c.news_id, c.comment_id, c.content, c.author, c.pub_time 
	FROM comments c
	ORDER BY id DESC 
	LIMIT $1
	`,
		n,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comments []storage.Comment

	for rows.Next() {
		var comm storage.Comment
		err = rows.Scan(
			&comm.ID,
			&comm.NewsID,
			&comm.CommentID,
			&comm.Content,
			&comm.Author,
			&comm.PubTime,
		)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comm)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, rows.Err()
}

func (s *Store) AddComment(comm storage.Comment) error {
	_, err := s.db.Exec(context.Background(), ` 
	INSERT INTO comments (news_id, comment_id, content, author, pub_time)
	VALUES ($1, $2, $3, $4, $5)
	`,
		comm.NewsID,
		comm.CommentID,
		comm.Content,
		comm.Author,
		comm.PubTime,
	)

	return err
}
