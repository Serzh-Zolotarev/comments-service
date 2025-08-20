package postgres

import (
	"comments-service/pkg/storage"
	cryptoRand "crypto/rand"
	"math/rand"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	_, err := New("postgres://postgres:postgres@localhost:5432/comments")
	if err != nil {
		t.Fatal(err)
	}
}

func TestStore_Comments(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	comment1 := storage.Comment{
		Author:  "Test Author",
		Content: cryptoRand.Text(),
	}
	comment2 := storage.Comment{
		Author:  "Test Author 2",
		Content: cryptoRand.Text(),
	}
	db, err := New("postgres://postgres:postgres@localhost:5432/comments")
	if err != nil {
		t.Fatal(err)
	}
	err = db.AddComment(comment1)
	if err != nil {
		t.Fatal(err)
	}
	err = db.AddComment(comment2)
	if err != nil {
		t.Fatal(err)
	}
	allComments, err := db.Comments(2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", allComments)
}

func TestStore_AddComment(t *testing.T) {
	comment := storage.Comment{
		Author:  "Author 3",
		Content: "Some Content",
	}
	db, err := New("postgres://postgres:postgres@localhost:5432/comments")
	if err != nil {
		t.Fatal(err)
	}
	err = db.AddComment(comment)
	if err != nil {
		t.Fatal(err)
	}
}
