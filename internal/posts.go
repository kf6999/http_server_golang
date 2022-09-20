package database

import (
	"time"
)

type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

func (c Client) GetPosts(userEmail string) ([]Post, error) {
	db, err := c.readDb()
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, post := range db.Posts {
		if post.UserEmail == userEmail {
			posts = append(posts, post)
		}
	}
	return posts, nil
}
