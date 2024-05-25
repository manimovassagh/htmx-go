package models

type Post struct {
	UserID int
	ID     int
	Title  string
	Body   string
}

type Posts struct {
	Posts []Post
}

var PublicPosts = Posts{
	Posts: []Post{
		{
			UserID: 1,
			ID:     1,
			Title:  "some title from mani",
			Body:   "Body 1",
		},
		{
			UserID: 2,
			ID:     2,
			Title:  "some more",
			Body:   "Body 2",
		},
		{
			UserID: 3,
			ID:     3,
			Title:  "some more",
			Body:   "Body 3",
		},
	},
}
