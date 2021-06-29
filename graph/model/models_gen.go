// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Post struct {
	ID             *string `json:"id"`
	Title          *string `json:"title"`
	Author         *User   `json:"author"`
	AuthorUsername *string `json:"authorUsername"`
}

type User struct {
	Username string  `json:"username"`
	Posts    []*Post `json:"posts"`
}
